package protobuf

import (
	"github.com/golang/protobuf/proto"
)

type CMsgClientCMList struct {
	CmAddresses      []uint32 `protobuf:"varint,1,rep,name=cm_addresses" json:"cm_addresses,omitempty"`
	CmPorts          []uint32 `protobuf:"varint,2,rep,name=cm_ports" json:"cm_ports,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *CMsgClientCMList) Reset()         { *m = CMsgClientCMList{} }
func (m *CMsgClientCMList) String() string { return proto.CompactTextString(m) }
func (*CMsgClientCMList) ProtoMessage()    {}

func (m *CMsgClientCMList) GetCmAddresses() []uint32 {
	if m != nil {
		return m.CmAddresses
	}
	return nil
}

func (m *CMsgClientCMList) GetCmPorts() []uint32 {
	if m != nil {
		return m.CmPorts
	}
	return nil
}
