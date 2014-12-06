package steamlang

import (
	"encoding/binary"
	"github.com/smithfox/gosteam/rwu"
	// "github.com/smithfox/gosteam/steamid"
	"io"
)

const MsgChannelEncryptRequest_PROTOCOL_VERSION uint32 = 1

type MsgChannelEncryptRequest struct {
	ProtocolVersion uint32
	Universe        EUniverse
}

func NewMsgChannelEncryptRequest() *MsgChannelEncryptRequest {
	return &MsgChannelEncryptRequest{
		ProtocolVersion: MsgChannelEncryptRequest_PROTOCOL_VERSION,
		Universe:        EUniverse_Invalid,
	}
}

func (d *MsgChannelEncryptRequest) GetEMsg() EMsg {
	return EMsg_ChannelEncryptRequest
}

func (d *MsgChannelEncryptRequest) Serialize(w io.Writer) error {
	var err error
	err = binary.Write(w, binary.LittleEndian, d.ProtocolVersion)
	if err != nil {
		return err
	}
	err = binary.Write(w, binary.LittleEndian, d.Universe)
	return err
}

func (d *MsgChannelEncryptRequest) Deserialize(r io.Reader) error {
	var err error
	d.ProtocolVersion, err = rwu.ReadUint32(r)
	if err != nil {
		return err
	}
	t0, err := rwu.ReadInt32(r)
	d.Universe = EUniverse(t0)
	return err
}

type MsgChannelEncryptResponse struct {
	ProtocolVersion uint32
	KeySize         uint32
}

func NewMsgChannelEncryptResponse() *MsgChannelEncryptResponse {
	return &MsgChannelEncryptResponse{
		ProtocolVersion: MsgChannelEncryptRequest_PROTOCOL_VERSION,
		KeySize:         128,
	}
}

func (d *MsgChannelEncryptResponse) GetEMsg() EMsg {
	return EMsg_ChannelEncryptResponse
}

func (d *MsgChannelEncryptResponse) Serialize(w io.Writer) error {
	var err error
	err = binary.Write(w, binary.LittleEndian, d.ProtocolVersion)
	if err != nil {
		return err
	}
	err = binary.Write(w, binary.LittleEndian, d.KeySize)
	return err
}

func (d *MsgChannelEncryptResponse) Deserialize(r io.Reader) error {
	var err error
	d.ProtocolVersion, err = rwu.ReadUint32(r)
	if err != nil {
		return err
	}
	d.KeySize, err = rwu.ReadUint32(r)
	return err
}
