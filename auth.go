package gosteam

import (
	"github.com/golang/protobuf/proto"
	"crypto/sha1"
	"fmt"
	"github.com/smithfox/gosteam/bot"
	. "github.com/smithfox/gosteam/internal"
	. "github.com/smithfox/gosteam/internal/protobuf"
	. "github.com/smithfox/gosteam/internal/steamlang"
	. "github.com/smithfox/gosteam/steamid"
	"time"
)

type Auth struct {
	client *Client
	b      *bot.BotRunTime
}

// Log on with the given details. You must always specify username and
// password. For the first login, don't set an authcode or a hash and you'll receive an error
// and Steam will send you an authcode. Then you have to login again, this time with the authcode.
// Shortly after logging in, you'll receive a MachineAuthUpdateEvent with a hash which allows
// you to login without using an authcode in the future.
//
// If you don't use Steam Guard, username and password are enough.
func (a *Auth) LogOn() error {
	if len(a.b.Username) == 0 || len(a.b.Password) == 0 {
		return fmt.Errorf("Username and password must be set!")
	}

	logon := new(CMsgClientLogon)
	logon.AccountName = &a.b.Username
	logon.Password = &a.b.Password
	if a.b.AuthCode != "" {
		logon.AuthCode = proto.String(a.b.AuthCode)
	}
	logon.ClientLanguage = proto.String("english")
	logon.ProtocolVersion = proto.Uint32(MsgClientLogon_CurrentProtocol)
	logon.ShaSentryfile = a.b.SentryFileHash
	tmp_steamid := NewIdAdv(0, 1, int32(EUniverse_Public), int32(EAccountType_Individual))
	a.b.SetSteamId(tmp_steamid)

	a.client.Write(NewClientMsgProtobuf(EMsg_ClientLogon, logon))
	return nil
}

func (a *Auth) HandlePacket(packet *Packet) {
	switch packet.EMsg {
	case EMsg_ClientLogOnResponse:
		a.handleLogOnResponse(packet)
	case EMsg_ClientNewLoginKey:
		a.handleLoginKey(packet)
	case EMsg_ClientSessionToken:
		a.b.Debugf("EMsg_ClientSessionToken\n")
	case EMsg_ClientLoggedOff:
		a.handleLoggedOff(packet)
	case EMsg_ClientUpdateMachineAuth:
		a.handleUpdateMachineAuth(packet)
	case EMsg_ClientAccountInfo:
		a.handleAccountInfo(packet)
	case EMsg_ClientWalletInfoUpdate:
	case EMsg_ClientRequestWebAPIAuthenticateUserNonceResponse:
		//case EMsg_ClientMarketingMessageUpdate:
	}
}

func (a *Auth) handleLogOnResponse(packet *Packet) {
	if !packet.IsProto {
		a.client.DisconnectWithErrorf("Got non-proto logon response!")
		return
	}

	body := new(CMsgClientLogonResponse)
	msg := packet.ReadProtoMsg(body)

	result := EResult(body.GetEresult())
	if result == EResult_OK {
		a.b.SetSessionId(msg.Header.Proto.GetClientSessionid())
		a.b.SetSteamId(SteamId(msg.Header.Proto.GetSteamid()))
		a.b.SetWebLoginKey(*body.WebapiAuthenticateUserNonce)
		go a.client.heartbeatLoop(time.Duration(body.GetOutOfGameHeartbeatSeconds()))

		a.b.Emit(&LoggedOnEvent{
			Result:                    EResult(body.GetEresult()),
			ExtendedResult:            EResult(body.GetEresultExtended()),
			OutOfGameSecsPerHeartbeat: body.GetOutOfGameHeartbeatSeconds(),
			InGameSecsPerHeartbeat:    body.GetInGameHeartbeatSeconds(),
			PublicIp:                  body.GetPublicIp(),
			ServerTime:                body.GetRtime32ServerTime(),
			AccountFlags:              EAccountFlags(body.GetAccountFlags()),
			ClientSteamId:             SteamId(body.GetClientSuppliedSteamid()),
			EmailDomain:               body.GetEmailDomain(),
			CellId:                    body.GetCellId(),
			CellIdPingThreshold:       body.GetCellIdPingThreshold(),
			Steam2Ticket:              body.GetSteam2Ticket(),
			UsePics:                   body.GetUsePics(),
			WebApiUserNonce:           body.GetWebapiAuthenticateUserNonce(),
			IpCountryCode:             body.GetIpCountryCode(),
			VanityUrl:                 body.GetVanityUrl(),
			NumLoginFailuresToMigrate: body.GetCountLoginfailuresToMigrate(),
			NumDisconnectsToMigrate:   body.GetCountDisconnectsToMigrate(),
		})
	} else if result == EResult_Fail || result == EResult_ServiceUnavailable || result == EResult_TryAnotherCM {
		// some error on Steam's side, we'll get an EOF later
	} else {
		a.client.DisconnectWithErrorf("Login error: %v", result)
	}
}

func (a *Auth) handleLoginKey(packet *Packet) {
	body := new(CMsgClientNewLoginKey)
	packet.ReadProtoMsg(body)
	a.client.Write(NewClientMsgProtobuf(EMsg_ClientNewLoginKeyAccepted, &CMsgClientNewLoginKeyAccepted{
		UniqueId: proto.Uint32(body.GetUniqueId()),
	}))
	a.b.Emit(&LoginKeyEvent{
		UniqueId: body.GetUniqueId(),
		LoginKey: body.GetLoginKey(),
	})
}

func (a *Auth) handleLoggedOff(packet *Packet) {
	result := EResult_Invalid
	if packet.IsProto {
		body := new(CMsgClientLoggedOff)
		packet.ReadProtoMsg(body)
		result = EResult(body.GetEresult())
	} else {
		body := new(MsgClientLoggedOff)
		packet.ReadClientMsg(body)
		result = body.Result
	}
	a.b.Emit(&LoggedOffEvent{Result: result})
}

func (a *Auth) handleUpdateMachineAuth(packet *Packet) {
	body := new(CMsgClientUpdateMachineAuth)
	packet.ReadProtoMsg(body)
	hash := sha1.New()
	hash.Write(packet.Data)
	sha := hash.Sum(nil)

	msg := NewClientMsgProtobuf(EMsg_ClientUpdateMachineAuthResponse, &CMsgClientUpdateMachineAuthResponse{
		ShaFile: sha,
	})
	msg.SetTargetJobId(packet.SourceJobId)
	a.client.Write(msg)

	a.b.Emit(&MachineAuthUpdateEvent{sha})
}

func (a *Auth) handleAccountInfo(packet *Packet) {
	body := new(CMsgClientAccountInfo)
	packet.ReadProtoMsg(body)
	a.b.Emit(&AccountInfoEvent{
		PersonaName:          body.GetPersonaName(),
		Country:              body.GetIpCountry(),
		PasswordSalt:         body.GetSaltPassword(),
		PasswordSHADisgest:   body.GetShaDigest_Password(),
		CountAuthedComputers: body.GetCountAuthedComputers(),
		LockedWithIpt:        body.GetLockedWithIpt(),
		AccountFlags:         EAccountFlags(body.GetAccountFlags()),
		FacebookId:           body.GetFacebookId(),
		FacebookName:         body.GetFacebookName(),
	})
}
