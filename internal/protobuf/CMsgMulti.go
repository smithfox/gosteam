package protobuf

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type CMsgMulti struct {
	SizeUnzipped     *uint32 `protobuf:"varint,1,opt,name=size_unzipped" json:"size_unzipped,omitempty"`
	MessageBody      []byte  `protobuf:"bytes,2,opt,name=message_body" json:"message_body,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CMsgMulti) Reset()         { *m = CMsgMulti{} }
func (m *CMsgMulti) String() string { return proto.CompactTextString(m) }
func (*CMsgMulti) ProtoMessage()    {}

func (m *CMsgMulti) GetSizeUnzipped() uint32 {
	if m != nil && m.SizeUnzipped != nil {
		return *m.SizeUnzipped
	}
	return 0
}

func (m *CMsgMulti) GetMessageBody() []byte {
	if m != nil {
		return m.MessageBody
	}
	return nil
}
