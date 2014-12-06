package protobuf

import (
	"github.com/golang/protobuf/proto"
)

type CMsgClientLogOff struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *CMsgClientLogOff) Reset()         { *m = CMsgClientLogOff{} }
func (m *CMsgClientLogOff) String() string { return proto.CompactTextString(m) }
func (*CMsgClientLogOff) ProtoMessage()    {}

type CMsgClientLoggedOff struct {
	Eresult          *int32 `protobuf:"varint,1,opt,name=eresult,def=2" json:"eresult,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *CMsgClientLoggedOff) Reset()         { *m = CMsgClientLoggedOff{} }
func (m *CMsgClientLoggedOff) String() string { return proto.CompactTextString(m) }
func (*CMsgClientLoggedOff) ProtoMessage()    {}

const Default_CMsgClientLoggedOff_Eresult int32 = 2

func (m *CMsgClientLoggedOff) GetEresult() int32 {
	if m != nil && m.Eresult != nil {
		return *m.Eresult
	}
	return Default_CMsgClientLoggedOff_Eresult
}
