package protobuf

import (
	"github.com/golang/protobuf/proto"
)

type CMsgClientNewLoginKey struct {
	UniqueId         *uint32 `protobuf:"varint,1,opt,name=unique_id" json:"unique_id,omitempty"`
	LoginKey         *string `protobuf:"bytes,2,opt,name=login_key" json:"login_key,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CMsgClientNewLoginKey) Reset()         { *m = CMsgClientNewLoginKey{} }
func (m *CMsgClientNewLoginKey) String() string { return proto.CompactTextString(m) }
func (*CMsgClientNewLoginKey) ProtoMessage()    {}

func (m *CMsgClientNewLoginKey) GetUniqueId() uint32 {
	if m != nil && m.UniqueId != nil {
		return *m.UniqueId
	}
	return 0
}

func (m *CMsgClientNewLoginKey) GetLoginKey() string {
	if m != nil && m.LoginKey != nil {
		return *m.LoginKey
	}
	return ""
}

type CMsgClientNewLoginKeyAccepted struct {
	UniqueId         *uint32 `protobuf:"varint,1,opt,name=unique_id" json:"unique_id,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CMsgClientNewLoginKeyAccepted) Reset()         { *m = CMsgClientNewLoginKeyAccepted{} }
func (m *CMsgClientNewLoginKeyAccepted) String() string { return proto.CompactTextString(m) }
func (*CMsgClientNewLoginKeyAccepted) ProtoMessage()    {}

func (m *CMsgClientNewLoginKeyAccepted) GetUniqueId() uint32 {
	if m != nil && m.UniqueId != nil {
		return *m.UniqueId
	}
	return 0
}
