package gosteam

import (
	"bytes"
	"compress/gzip"
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/smithfox/gosteam/bot"
	"github.com/smithfox/gosteam/cryptoutil"
	. "github.com/smithfox/gosteam/internal"
	. "github.com/smithfox/gosteam/internal/protobuf"
	. "github.com/smithfox/gosteam/internal/steamlang"
	"github.com/smithfox/gosteam/netutil"
	"hash/crc32"
	"io/ioutil"
	"net"
	"sync"
	"time"
)

// When this event is emitted by the Client, the connection is automatically closed.
// This may be caused by a network error, for example.
type FatalErrorEvent error

type ConnectedEvent struct{}

type DisconnectedEvent struct{}

// A list of connection manager addresses to connect to in the future.
// You should always save them and then select one of these
// instead of the builtin ones for the next connection.
type ClientCMListEvent struct {
	Addresses []*netutil.PortAddr
}

// Represents a client to the Steam network.
// Always poll events from the channel returned by Events() or receiving messages will stop.
// All access, unless otherwise noted, should be threadsafe.
//
// When a FatalErrorEvent is emitted, the connection is automatically closed. The same client can be used to reconnect.
// Other errors don't have any effect.
type Client struct {
	b              *bot.BotRunTime
	Auth           *Auth
	Web            *Web
	handlers       []PacketHandler
	handlersMutex  sync.RWMutex
	tempSessionKey []byte
	mutex          sync.RWMutex // guarding conn and writeChan
	conn           connection
	writeChan      chan IMsg
	writeBuf       *bytes.Buffer
	heartbeat      *time.Ticker
}

type PacketHandler interface {
	HandlePacket(*Packet)
}

func NewClient(b *bot.BotRunTime) *Client {
	client := &Client{
		b:         b,
		writeChan: make(chan IMsg, 5),
		writeBuf:  new(bytes.Buffer),
	}
	client.Auth = &Auth{client: client, b: b}
	client.RegisterPacketHandler(client.Auth)
	client.Web = &Web{client: client, b: b}
	client.RegisterPacketHandler(client.Web)

	return client
}

// Registers a PacketHandler that receives all incoming packets.
func (c *Client) RegisterPacketHandler(handler PacketHandler) {
	c.handlersMutex.Lock()
	defer c.handlersMutex.Unlock()
	c.handlers = append(c.handlers, handler)
}

func (c *Client) Connected() bool {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.conn != nil
}

// Connects to a random server of the included list of connection managers and returns the address.
// If this client is already connected, it is disconnected first.
//
// You will receive a ServerListEvent after logging in which contains a new list of servers of which you
// should choose one yourself and connect with ConnectTo since the included list may not always be up to date.
func (c *Client) Connect() error {
	server := GetRandomCM()
	if server != nil {
		c.b.Debugf("choice cm server:%s\n", server.String())
	}
	return c.ConnectTo(server)
}

// Connects to a specific server.
// If this client is already connected, it is disconnected first.
func (c *Client) ConnectTo(addr *netutil.PortAddr) error {
	c.Disconnect()

	conn, err := dialTCP(addr.String())
	if err != nil {
		return err
	}
	c.conn = conn

	go c.readLoop()
	go c.writeLoop()
	return nil
}

func (c *Client) DisconnectWithErrorf(format string, a ...interface{}) {
	c.b.Emit(FatalErrorEvent(fmt.Errorf(format, a...)))
	c.Disconnect()
}

func (c *Client) Disconnect() {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.conn == nil {
		return
	}

	c.conn.Close()
	c.conn = nil
	if c.heartbeat != nil {
		c.heartbeat.Stop()
	}
	close(c.writeChan)
	c.b.Emit(&DisconnectedEvent{})

}

// Adds a message to the send queue. Modifications to the given message after
// writing are not allowed (possible race conditions).
//
// Writes to this client when not connected are ignored.
func (c *Client) Write(msg IMsg) {
	if cm, ok := msg.(IClientMsg); ok {
		cm.SetSessionId(c.b.SessionId())
		cm.SetSteamId(c.b.SteamId())
	}
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	if c.conn == nil {
		return
	}
	c.writeChan <- msg
}

func (c *Client) readLoop() {
	for {
		// This *should* be atomic on most platforms, but the Go spec doesn't guarantee it
		c.mutex.RLock()
		conn := c.conn
		c.mutex.RUnlock()
		if conn == nil {
			return
		}
		packet, err := conn.Read()

		if err != nil {
			c.DisconnectWithErrorf("Error reading from the connection: %v", err)
			return
		}
		c.handlePacket(packet)
	}
}

func (c *Client) writeLoop() {
	defer c.Disconnect()
	for {
		c.mutex.RLock()
		conn := c.conn
		c.mutex.RUnlock()
		if conn == nil {
			return
		}

		msg, ok := <-c.writeChan
		if !ok {
			return
		}

		err := msg.Serialize(c.writeBuf)
		if err != nil {
			c.writeBuf.Reset()
			c.b.EmitErrorf("Error serializing message %v: %v", msg, err)
			return
		}

		err = conn.Write(c.writeBuf.Bytes())

		c.writeBuf.Reset()

		if err != nil {
			c.b.EmitErrorf("Error writing message %v: %v", msg, err)
			return
		}
	}
}

func (c *Client) heartbeatLoop(seconds time.Duration) {
	if c.heartbeat != nil {
		c.heartbeat.Stop()
	}
	c.heartbeat = time.NewTicker(seconds * time.Second)
	for {
		_, ok := <-c.heartbeat.C
		if !ok {
			break
		}
		c.Write(NewClientMsgProtobuf(EMsg_ClientHeartBeat, new(CMsgClientHeartBeat)))
	}
	c.heartbeat = nil
}

func (c *Client) handlePacket(packet *Packet) {
	switch packet.EMsg {
	case EMsg_ChannelEncryptRequest:
		c.handleChannelEncryptRequest(packet)
	case EMsg_ChannelEncryptResult:
		c.handleChannelEncryptResult(packet)
	case EMsg_Multi:
		c.handleMulti(packet)
	case EMsg_ClientCMList:
		c.handleClientCMList(packet)
	}

	c.handlersMutex.RLock()
	defer c.handlersMutex.RUnlock()
	for _, handler := range c.handlers {
		handler.HandlePacket(packet)
	}
}

func (c *Client) handleChannelEncryptRequest(packet *Packet) {
	body := NewMsgChannelEncryptRequest()
	packet.ReadMsg(body)

	if body.Universe != EUniverse_Public {
		c.DisconnectWithErrorf("Invalid univserse %v!", body.Universe)
	}

	c.tempSessionKey = make([]byte, 32)
	rand.Read(c.tempSessionKey)
	encryptedKey := cryptoutil.RSAEncrypt(GetPublicKey(EUniverse_Public), c.tempSessionKey)

	payload := new(bytes.Buffer)
	payload.Write(encryptedKey)
	binary.Write(payload, binary.LittleEndian, crc32.ChecksumIEEE(encryptedKey))
	payload.WriteByte(0)
	payload.WriteByte(0)
	payload.WriteByte(0)
	payload.WriteByte(0)

	c.Write(NewMsg(NewMsgChannelEncryptResponse(), payload.Bytes()))
}

func (c *Client) handleChannelEncryptResult(packet *Packet) {
	body := NewMsgChannelEncryptResult()
	packet.ReadMsg(body)

	if body.Result != EResult_OK {
		c.DisconnectWithErrorf("Encryption failed: %v", body.Result)
		return
	}
	c.conn.SetEncryptionKey(c.tempSessionKey)
	c.tempSessionKey = nil

	c.b.Emit(&ConnectedEvent{})
}

func (c *Client) handleMulti(packet *Packet) {
	body := new(CMsgMulti)
	packet.ReadProtoMsg(body)

	payload := body.GetMessageBody()

	if body.GetSizeUnzipped() > 0 {
		r, err := gzip.NewReader(bytes.NewReader(payload))
		if err != nil {
			c.b.EmitErrorf("handleMulti: Error while decompressing: %v", err)
			return
		}

		payload, err = ioutil.ReadAll(r)
		if err != nil {
			c.b.EmitErrorf("handleMulti: Error while decompressing: %v", err)
			return
		}
	}

	pr := bytes.NewReader(payload)
	for pr.Len() > 0 {
		var length uint32
		binary.Read(pr, binary.LittleEndian, &length)
		packetData := make([]byte, length)
		pr.Read(packetData)
		p, err := NewPacket(packetData)
		if err != nil {
			c.b.EmitErrorf("Error reading packet in Multi msg %v: %v", packet, err)
			continue
		}
		c.handlePacket(p)
	}
}

func (c *Client) handleClientCMList(packet *Packet) {
	body := new(CMsgClientCMList)
	packet.ReadProtoMsg(body)

	l := make([]*netutil.PortAddr, 0)
	for i, ip := range body.GetCmAddresses() {
		l = append(l, &netutil.PortAddr{
			readIp(ip),
			uint16(body.GetCmPorts()[i]),
		})
	}

	c.b.Emit(&ClientCMListEvent{l})
}

func readIp(ip uint32) net.IP {
	r := make(net.IP, 4)
	r[3] = byte(ip)
	r[2] = byte(ip >> 8)
	r[1] = byte(ip >> 16)
	r[0] = byte(ip >> 24)
	return r
}

// Sets the local user's persona state and broadcasts it over the network
func (c *Client) SetPersonaState(state EPersonaState) {
	c.Write(NewClientMsgProtobuf(EMsg_ClientChangeStatus, &CMsgClientChangeStatus{
		PersonaState: proto.Uint32(uint32(state)),
	}))
}

func (c *Client) Run(f func(ee interface{})) {
	c.Connect()

	for event := range c.b.Events() {
		switch e := event.(type) {
		case *ConnectedEvent:
			err := c.Auth.LogOn()
			if err != nil {
				c.b.Errorf("auth logon err=%v\n", err)
			}
		case *MachineAuthUpdateEvent:
			c.b.Debugf("event: MachineAuthUpdateEvent\n")
			c.b.WriteSentry(bot.SentryHash(e.Hash))
		case *LoggedOnEvent:
			c.b.Debugf("event: LoggedOnEvent\n")
			c.SetPersonaState(EPersonaState_Online)
		case *ClientCMListEvent:
			c.b.Debugf("event: ClientCMListEvent, len=%d\n", len(e.Addresses))
			UpdateCMServers(e.Addresses)
		case *LoginKeyEvent:
			c.b.Debugf("event: LoginKeyEvent\n")
			// case *WebSessionIdEvent:
			// 	c.b.Debugf("event: WebSessionIdEvent\n")
			err := c.Web.LogOn()
			if err != nil {
				c.b.Errorf("web logon err=%v\n", err)
			}
		case *WebLoggedOnEvent:
			c.b.Debugf("event: WebLoggedOnEvent,websessionid=%s,websteamlogin=%s\n", c.b.WebSessionId(), c.b.WebSteamLogin())
			if f != nil {
				f(event)
			}
		case FatalErrorEvent:
			c.b.Debugf("event: FatalErrorEvent, e=%v\n", e)
			if f != nil {
				f(event)
			}
		case error:
			c.b.Debugf("error: client events error, e=%v\n", e)
		default:
			if f != nil {
				f(event)
			}
		}
	}
}
