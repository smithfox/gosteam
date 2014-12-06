package steamlang

import (
	"encoding/binary"
	"github.com/golang/protobuf/proto"
	. "github.com/smithfox/gosteam/internal/protobuf"
	"github.com/smithfox/gosteam/rwu"
	"io"
)

type MsgHdrProtoBuf struct {
	Msg          EMsg
	HeaderLength int32
	Proto        *CMsgProtoBufHeader
}

func NewMsgHdrProtoBuf() *MsgHdrProtoBuf {
	return &MsgHdrProtoBuf{
		Msg:   EMsg_Invalid,
		Proto: new(CMsgProtoBufHeader),
	}
}

func (d *MsgHdrProtoBuf) Serialize(w io.Writer) error {
	var err error
	buf0, err := proto.Marshal(d.Proto)
	if err != nil {
		return err
	}
	d.HeaderLength = int32(len(buf0))
	err = binary.Write(w, binary.LittleEndian, EMsg(uint32(d.Msg)|ProtoMask))
	if err != nil {
		return err
	}
	err = binary.Write(w, binary.LittleEndian, d.HeaderLength)
	if err != nil {
		return err
	}
	_, err = w.Write(buf0)
	return err
}

func (d *MsgHdrProtoBuf) Deserialize(r io.Reader) error {
	var err error
	t0, err := rwu.ReadInt32(r)
	if err != nil {
		return err
	}
	d.Msg = EMsg(uint32(t0) & EMsgMask)
	d.HeaderLength, err = rwu.ReadInt32(r)
	if err != nil {
		return err
	}
	buf1 := make([]byte, d.HeaderLength, d.HeaderLength)
	_, err = io.ReadFull(r, buf1)
	if err != nil {
		return err
	}
	err = proto.Unmarshal(buf1, d.Proto)
	return err
}
