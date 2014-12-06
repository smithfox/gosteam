package steamlang

import (
	"encoding/binary"
	"github.com/smithfox/gosteam/rwu"
	"github.com/smithfox/gosteam/steamid"
	"io"
)

type ExtendedClientMsgHdr struct {
	Msg           EMsg
	HeaderSize    uint8
	HeaderVersion uint16
	TargetJobID   uint64
	SourceJobID   uint64
	HeaderCanary  uint8
	SteamID       steamid.SteamId
	SessionID     int32
}

func NewExtendedClientMsgHdr() *ExtendedClientMsgHdr {
	return &ExtendedClientMsgHdr{
		Msg:           EMsg_Invalid,
		HeaderSize:    36,
		HeaderVersion: 2,
		TargetJobID:   ^uint64(0),
		SourceJobID:   ^uint64(0),
		HeaderCanary:  239,
	}
}

func (d *ExtendedClientMsgHdr) Serialize(w io.Writer) error {
	var err error
	err = binary.Write(w, binary.LittleEndian, d.Msg)
	if err != nil {
		return err
	}
	err = binary.Write(w, binary.LittleEndian, d.HeaderSize)
	if err != nil {
		return err
	}
	err = binary.Write(w, binary.LittleEndian, d.HeaderVersion)
	if err != nil {
		return err
	}
	err = binary.Write(w, binary.LittleEndian, d.TargetJobID)
	if err != nil {
		return err
	}
	err = binary.Write(w, binary.LittleEndian, d.SourceJobID)
	if err != nil {
		return err
	}
	err = binary.Write(w, binary.LittleEndian, d.HeaderCanary)
	if err != nil {
		return err
	}
	err = binary.Write(w, binary.LittleEndian, d.SteamID)
	if err != nil {
		return err
	}
	err = binary.Write(w, binary.LittleEndian, d.SessionID)
	return err
}

func (d *ExtendedClientMsgHdr) Deserialize(r io.Reader) error {
	var err error
	t0, err := rwu.ReadInt32(r)
	if err != nil {
		return err
	}
	d.Msg = EMsg(t0)
	d.HeaderSize, err = rwu.ReadUint8(r)
	if err != nil {
		return err
	}
	d.HeaderVersion, err = rwu.ReadUint16(r)
	if err != nil {
		return err
	}
	d.TargetJobID, err = rwu.ReadUint64(r)
	if err != nil {
		return err
	}
	d.SourceJobID, err = rwu.ReadUint64(r)
	if err != nil {
		return err
	}
	d.HeaderCanary, err = rwu.ReadUint8(r)
	if err != nil {
		return err
	}
	t1, err := rwu.ReadUint64(r)
	if err != nil {
		return err
	}
	d.SteamID = steamid.SteamId(t1)
	d.SessionID, err = rwu.ReadInt32(r)
	return err
}
