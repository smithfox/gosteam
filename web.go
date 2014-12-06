package gosteam

import (
	"github.com/golang/protobuf/proto"
	"crypto/aes"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/smithfox/gosteam/bot"
	"github.com/smithfox/gosteam/cryptoutil"
	. "github.com/smithfox/gosteam/internal"
	. "github.com/smithfox/gosteam/internal/protobuf"
	. "github.com/smithfox/gosteam/internal/steamlang"
	"net/http"
	"net/url"
	"strconv"
	"sync/atomic"
	"time"
)

type WebLoggedOnEvent struct{}

type WebSessionIdEvent struct{}

type Web struct {
	// 64 bit alignment
	relogOnNonce uint32
	_            uint32 //align on 32 bit pc

	client *Client
	b      *bot.BotRunTime
}

func (w *Web) HandlePacket(packet *Packet) {
	switch packet.EMsg {
	case EMsg_ClientNewLoginKey:
		w.handleNewLoginKey(packet)
	case EMsg_ClientRequestWebAPIAuthenticateUserNonceResponse:
		w.handleAuthNonceResponse(packet)
	}
}

// Fetches the `steamLogin` cookie. This may only be called after the first
// WebSessionIdEvent or it will panic.
func (w *Web) LogOn() error {
	if w.b.WebLoginKey() == "" {
		return fmt.Errorf("Web: webLoginKey not initialized!")
	}

	go func() {
		// retry three times. yes, I know about loops.
		err := w.apiLogOn()
		// if err != nil {
		// 	err = w.apiLogOn()
		// 	if err != nil {
		// 		err = w.apiLogOn()
		// 	}
		// }
		if err != nil {
			w.client.DisconnectWithErrorf("Web: Error logging on: %v", err)
			return
		}
	}()
	return nil
}

func (w *Web) apiLogOn() error {
	sessionKey := make([]byte, 32)
	rand.Read(sessionKey)

	cryptedSessionKey := cryptoutil.RSAEncrypt(GetPublicKey(EUniverse_Public), sessionKey)
	ciph, _ := aes.NewCipher(sessionKey)
	cryptedLoginKey := cryptoutil.SymmetricEncrypt(ciph, []byte(w.b.WebLoginKey()))
	data := make(url.Values)
	// data.Add("format", "json")
	data.Add("steamid", strconv.FormatUint(uint64(w.b.SteamId()), 10))
	data.Add("sessionkey", string(cryptedSessionKey))
	data.Add("encrypted_loginkey", string(cryptedLoginKey))
	resp, err := http.PostForm("https://api.steampowered.com/ISteamUserAuth/AuthenticateUser/v1", data)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		time.Sleep(500)
		// our login key has expired, request a new one
		atomic.StoreUint32(&w.relogOnNonce, 1)
		w.client.Write(NewClientMsgProtobuf(EMsg_ClientRequestWebAPIAuthenticateUserNonce, new(CMsgClientRequestWebAPIAuthenticateUserNonce)))
		return nil
		//return errors.New("steam.Web.apiLogOn: request failed with status " + resp.Status)
	}

	result := new(struct {
		Authenticateuser struct {
			Token       string
			Tokensecure string
		}
	})

	w.b.Debugf("apiLogon, response cookies=%v\n", resp.Cookies())
	err = json.NewDecoder(resp.Body).Decode(result)
	if err != nil {
		return err
	}
	w.b.SetWebSteamLogin(result.Authenticateuser.Token)
	w.b.SetWebSteamLoginSecure(result.Authenticateuser.Tokensecure)

	w.b.Emit(new(WebLoggedOnEvent))
	return nil
}

func (w *Web) handleNewLoginKey(packet *Packet) {
	msg := new(CMsgClientNewLoginKey)
	packet.ReadProtoMsg(msg)

	w.client.Write(NewClientMsgProtobuf(EMsg_ClientNewLoginKeyAccepted, &CMsgClientNewLoginKeyAccepted{
		UniqueId: proto.Uint32(msg.GetUniqueId()),
	}))

	//w.webLoginKey = msg.GetLoginKey()
	// number -> string -> bytes -> base64
	websessionid := base64.StdEncoding.EncodeToString([]byte(strconv.FormatUint(uint64(msg.GetUniqueId()), 10)))
	w.b.SetWebSessionId(websessionid)
	w.b.Emit(new(WebSessionIdEvent))
}

func (w *Web) handleAuthNonceResponse(packet *Packet) {
	// this has to be the best name for a message yet.
	msg := new(CMsgClientRequestWebAPIAuthenticateUserNonceResponse)
	packet.ReadProtoMsg(msg)
	w.b.SetWebLoginKey(msg.GetWebapiAuthenticateUserNonce())

	// if the nonce was specifically requested in apiLogOn(),
	// don't emit an event.
	if atomic.CompareAndSwapUint32(&w.relogOnNonce, 1, 0) {
		w.LogOn()
	}
}
