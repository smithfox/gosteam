package steamlang

import (
	"encoding/binary"
	"github.com/smithfox/gosteam/rwu"
	"io"
)

type MsgHdr struct {
	Msg         EMsg
	TargetJobID uint64
	SourceJobID uint64
}

func NewMsgHdr() *MsgHdr {
	return &MsgHdr{
		Msg:         EMsg_Invalid,
		TargetJobID: ^uint64(0),
		SourceJobID: ^uint64(0),
	}
}

func (d *MsgHdr) Serialize(w io.Writer) error {
	var err error
	err = binary.Write(w, binary.LittleEndian, d.Msg)
	if err != nil {
		return err
	}
	err = binary.Write(w, binary.LittleEndian, d.TargetJobID)
	if err != nil {
		return err
	}
	err = binary.Write(w, binary.LittleEndian, d.SourceJobID)
	return err
}

func (d *MsgHdr) Deserialize(r io.Reader) error {
	var err error
	t0, err := rwu.ReadInt32(r)
	if err != nil {
		return err
	}
	d.Msg = EMsg(t0)
	d.TargetJobID, err = rwu.ReadUint64(r)
	if err != nil {
		return err
	}
	d.SourceJobID, err = rwu.ReadUint64(r)
	return err
}
