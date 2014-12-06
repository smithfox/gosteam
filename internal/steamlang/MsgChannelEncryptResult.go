package steamlang

import (
	"encoding/binary"
	"github.com/smithfox/gosteam/rwu"
	"io"
)

type MsgChannelEncryptResult struct {
	Result EResult
}

func NewMsgChannelEncryptResult() *MsgChannelEncryptResult {
	return &MsgChannelEncryptResult{
		Result: EResult_Invalid,
	}
}

func (d *MsgChannelEncryptResult) GetEMsg() EMsg {
	return EMsg_ChannelEncryptResult
}

func (d *MsgChannelEncryptResult) Serialize(w io.Writer) error {
	var err error
	err = binary.Write(w, binary.LittleEndian, d.Result)
	return err
}

func (d *MsgChannelEncryptResult) Deserialize(r io.Reader) error {
	var err error
	t0, err := rwu.ReadInt32(r)
	d.Result = EResult(t0)
	return err
}
