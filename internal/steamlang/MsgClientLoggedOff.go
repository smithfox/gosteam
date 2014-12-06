package steamlang

import (
	"encoding/binary"
	"github.com/smithfox/gosteam/rwu"
	"github.com/smithfox/gosteam/steamid"
	"io"
)

const MsgClientLogon_CurrentProtocol uint32 = 65579

type MsgClientLoggedOff struct {
	Result              EResult
	SecMinReconnectHint int32
	SecMaxReconnectHint int32
}

func NewMsgClientLoggedOff() *MsgClientLoggedOff {
	return &MsgClientLoggedOff{}
}

func (d *MsgClientLoggedOff) GetEMsg() EMsg {
	return EMsg_ClientLoggedOff
}

func (d *MsgClientLoggedOff) Serialize(w io.Writer) error {
	var err error
	err = binary.Write(w, binary.LittleEndian, d.Result)
	if err != nil {
		return err
	}
	err = binary.Write(w, binary.LittleEndian, d.SecMinReconnectHint)
	if err != nil {
		return err
	}
	err = binary.Write(w, binary.LittleEndian, d.SecMaxReconnectHint)
	return err
}

func (d *MsgClientLoggedOff) Deserialize(r io.Reader) error {
	var err error
	t0, err := rwu.ReadInt32(r)
	if err != nil {
		return err
	}
	d.Result = EResult(t0)
	d.SecMinReconnectHint, err = rwu.ReadInt32(r)
	if err != nil {
		return err
	}
	d.SecMaxReconnectHint, err = rwu.ReadInt32(r)
	return err
}

type MsgClientLogOnResponse struct {
	Result                    EResult
	OutOfGameHeartbeatRateSec int32
	InGameHeartbeatRateSec    int32
	ClientSuppliedSteamId     steamid.SteamId
	IpPublic                  uint32
	ServerRealTime            uint32
}

func NewMsgClientLogOnResponse() *MsgClientLogOnResponse {
	return &MsgClientLogOnResponse{}
}

func (d *MsgClientLogOnResponse) GetEMsg() EMsg {
	return EMsg_ClientLogOnResponse
}

func (d *MsgClientLogOnResponse) Serialize(w io.Writer) error {
	var err error
	err = binary.Write(w, binary.LittleEndian, d.Result)
	if err != nil {
		return err
	}
	err = binary.Write(w, binary.LittleEndian, d.OutOfGameHeartbeatRateSec)
	if err != nil {
		return err
	}
	err = binary.Write(w, binary.LittleEndian, d.InGameHeartbeatRateSec)
	if err != nil {
		return err
	}
	err = binary.Write(w, binary.LittleEndian, d.ClientSuppliedSteamId)
	if err != nil {
		return err
	}
	err = binary.Write(w, binary.LittleEndian, d.IpPublic)
	if err != nil {
		return err
	}
	err = binary.Write(w, binary.LittleEndian, d.ServerRealTime)
	return err
}

func (d *MsgClientLogOnResponse) Deserialize(r io.Reader) error {
	var err error
	t0, err := rwu.ReadInt32(r)
	if err != nil {
		return err
	}
	d.Result = EResult(t0)
	d.OutOfGameHeartbeatRateSec, err = rwu.ReadInt32(r)
	if err != nil {
		return err
	}
	d.InGameHeartbeatRateSec, err = rwu.ReadInt32(r)
	if err != nil {
		return err
	}
	t1, err := rwu.ReadUint64(r)
	if err != nil {
		return err
	}
	d.ClientSuppliedSteamId = steamid.SteamId(t1)
	d.IpPublic, err = rwu.ReadUint32(r)
	if err != nil {
		return err
	}
	d.ServerRealTime, err = rwu.ReadUint32(r)
	return err
}
