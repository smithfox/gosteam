package gosteam

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/binary"
	"fmt"
	"github.com/smithfox/gosteam/cryptoutil"
	. "github.com/smithfox/gosteam/internal"
	"io"
	"net"
	"sync"
	"time"
)

type connection interface {
	Read() (*Packet, error)
	Write([]byte) error
	Close() error
	SetEncryptionKey([]byte)
	IsEncrypted() bool
}

const tcpConnectionMagic uint32 = 0x31305456 // "VT01"

type tcpConnection struct {
	conn         net.Conn
	ciph         cipher.Block
	cipherMutex  sync.RWMutex
	readTimeout  time.Duration //unit is second
	writeTimeout time.Duration //unit is second
}

func dialTCPTimeout(addr string, conn_timeout int, read_timeout int, write_timeout int) (*tcpConnection, error) {
	//fmt.Printf("before dialtimeout\n")

	conn, err := net.DialTimeout("tcp", addr, time.Duration(conn_timeout)*time.Second)
	//conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		return nil, err
	}
	//fmt.Printf("after dialtimeout\n")
	//conn := *net.TCPConn(conn1)

	return &tcpConnection{
		conn:         conn,
		readTimeout:  time.Duration(read_timeout),
		writeTimeout: time.Duration(write_timeout),
	}, nil
}

func dialTCP(addr string) (*tcpConnection, error) {
	return dialTCPTimeout(addr, 30, 30, 30)
}

func (c *tcpConnection) Read() (*Packet, error) {
	c.conn.SetReadDeadline(time.Now().Add(c.readTimeout * time.Second))
	c.conn.SetWriteDeadline(time.Now().Add(c.writeTimeout * time.Second))
	// All packets begin with a packet length
	var packetLen uint32
	err := binary.Read(c.conn, binary.LittleEndian, &packetLen)
	//fmt.Printf("tcpConnection.Read1\n")
	if err != nil {
		return nil, err
	}

	// A magic value follows for validation
	var packetMagic uint32
	err = binary.Read(c.conn, binary.LittleEndian, &packetMagic)
	//fmt.Printf("tcpConnection.Read2\n")
	if err != nil {
		return nil, err
	}
	if packetMagic != tcpConnectionMagic {
		return nil, fmt.Errorf("Invalid connection magic! Expected %d, got %d!", tcpConnectionMagic, packetMagic)
	}

	buf := make([]byte, packetLen, packetLen)
	_, err = io.ReadFull(c.conn, buf)
	//fmt.Printf("tcpConnection.Read3\n")
	if err == io.ErrUnexpectedEOF {
		return nil, io.EOF
	}
	if err != nil {
		return nil, err
	}

	// Packets after ChannelEncryptResult are encrypted
	c.cipherMutex.RLock()
	if c.ciph != nil {
		buf = cryptoutil.SymmetricDecrypt(c.ciph, buf)
	}
	c.cipherMutex.RUnlock()

	return NewPacket(buf)
}

// Writes a message. This may only be used by one goroutine at a time.
func (c *tcpConnection) Write(message []byte) error {
	c.conn.SetReadDeadline(time.Now().Add(c.readTimeout * time.Second))
	c.conn.SetWriteDeadline(time.Now().Add(c.writeTimeout * time.Second))
	c.cipherMutex.RLock()
	if c.ciph != nil {
		message = cryptoutil.SymmetricEncrypt(c.ciph, message)
	}
	c.cipherMutex.RUnlock()

	err := binary.Write(c.conn, binary.LittleEndian, uint32(len(message)))
	//fmt.Printf("tcpConnection.Write1\n")
	if err != nil {
		return err
	}
	err = binary.Write(c.conn, binary.LittleEndian, tcpConnectionMagic)
	//fmt.Printf("tcpConnection.Write2\n")
	if err != nil {
		return err
	}

	_, err = c.conn.Write(message)
	//fmt.Printf("tcpConnection.Write3\n")
	return err
}

func (c *tcpConnection) Close() error {
	return c.conn.Close()
}

func (c *tcpConnection) SetEncryptionKey(key []byte) {
	c.cipherMutex.Lock()
	defer c.cipherMutex.Unlock()
	if key == nil {
		c.ciph = nil
		return
	}
	if len(key) != 32 {
		panic("Connection AES key is not 32 bytes long!")
	}

	var err error
	c.ciph, err = aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
}

func (c *tcpConnection) IsEncrypted() bool {
	c.cipherMutex.RLock()
	defer c.cipherMutex.RUnlock()
	return c.ciph != nil
}
