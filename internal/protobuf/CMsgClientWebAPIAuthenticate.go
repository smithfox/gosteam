package protobuf

import (
	"github.com/golang/protobuf/proto"
)

type CMsgClientRequestWebAPIAuthenticateUserNonce struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *CMsgClientRequestWebAPIAuthenticateUserNonce) Reset() {
	*m = CMsgClientRequestWebAPIAuthenticateUserNonce{}
}
func (m *CMsgClientRequestWebAPIAuthenticateUserNonce) String() string {
	return proto.CompactTextString(m)
}
func (*CMsgClientRequestWebAPIAuthenticateUserNonce) ProtoMessage() {}

type CMsgClientRequestWebAPIAuthenticateUserNonceResponse struct {
	Eresult                     *int32  `protobuf:"varint,1,opt,name=eresult,def=2" json:"eresult,omitempty"`
	WebapiAuthenticateUserNonce *string `protobuf:"bytes,11,opt,name=webapi_authenticate_user_nonce" json:"webapi_authenticate_user_nonce,omitempty"`
	XXX_unrecognized            []byte  `json:"-"`
}

func (m *CMsgClientRequestWebAPIAuthenticateUserNonceResponse) Reset() {
	*m = CMsgClientRequestWebAPIAuthenticateUserNonceResponse{}
}
func (m *CMsgClientRequestWebAPIAuthenticateUserNonceResponse) String() string {
	return proto.CompactTextString(m)
}
func (*CMsgClientRequestWebAPIAuthenticateUserNonceResponse) ProtoMessage() {}

const Default_CMsgClientRequestWebAPIAuthenticateUserNonceResponse_Eresult int32 = 2

func (m *CMsgClientRequestWebAPIAuthenticateUserNonceResponse) GetEresult() int32 {
	if m != nil && m.Eresult != nil {
		return *m.Eresult
	}
	return Default_CMsgClientRequestWebAPIAuthenticateUserNonceResponse_Eresult
}

func (m *CMsgClientRequestWebAPIAuthenticateUserNonceResponse) GetWebapiAuthenticateUserNonce() string {
	if m != nil && m.WebapiAuthenticateUserNonce != nil {
		return *m.WebapiAuthenticateUserNonce
	}
	return ""
}
