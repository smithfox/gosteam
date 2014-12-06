package protobuf

import (
	"github.com/golang/protobuf/proto"
)

type CMsgClientUpdateMachineAuth struct {
	Filename         *string `protobuf:"bytes,1,opt,name=filename" json:"filename,omitempty"`
	Offset           *uint32 `protobuf:"varint,2,opt,name=offset" json:"offset,omitempty"`
	Cubtowrite       *uint32 `protobuf:"varint,3,opt,name=cubtowrite" json:"cubtowrite,omitempty"`
	Bytes            []byte  `protobuf:"bytes,4,opt,name=bytes" json:"bytes,omitempty"`
	OtpType          *uint32 `protobuf:"varint,5,opt,name=otp_type" json:"otp_type,omitempty"`
	OtpIdentifier    *string `protobuf:"bytes,6,opt,name=otp_identifier" json:"otp_identifier,omitempty"`
	OtpSharedsecret  []byte  `protobuf:"bytes,7,opt,name=otp_sharedsecret" json:"otp_sharedsecret,omitempty"`
	OtpTimedrift     *uint32 `protobuf:"varint,8,opt,name=otp_timedrift" json:"otp_timedrift,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CMsgClientUpdateMachineAuth) Reset()         { *m = CMsgClientUpdateMachineAuth{} }
func (m *CMsgClientUpdateMachineAuth) String() string { return proto.CompactTextString(m) }
func (*CMsgClientUpdateMachineAuth) ProtoMessage()    {}

func (m *CMsgClientUpdateMachineAuth) GetFilename() string {
	if m != nil && m.Filename != nil {
		return *m.Filename
	}
	return ""
}

func (m *CMsgClientUpdateMachineAuth) GetOffset() uint32 {
	if m != nil && m.Offset != nil {
		return *m.Offset
	}
	return 0
}

func (m *CMsgClientUpdateMachineAuth) GetCubtowrite() uint32 {
	if m != nil && m.Cubtowrite != nil {
		return *m.Cubtowrite
	}
	return 0
}

func (m *CMsgClientUpdateMachineAuth) GetBytes() []byte {
	if m != nil {
		return m.Bytes
	}
	return nil
}

func (m *CMsgClientUpdateMachineAuth) GetOtpType() uint32 {
	if m != nil && m.OtpType != nil {
		return *m.OtpType
	}
	return 0
}

func (m *CMsgClientUpdateMachineAuth) GetOtpIdentifier() string {
	if m != nil && m.OtpIdentifier != nil {
		return *m.OtpIdentifier
	}
	return ""
}

func (m *CMsgClientUpdateMachineAuth) GetOtpSharedsecret() []byte {
	if m != nil {
		return m.OtpSharedsecret
	}
	return nil
}

func (m *CMsgClientUpdateMachineAuth) GetOtpTimedrift() uint32 {
	if m != nil && m.OtpTimedrift != nil {
		return *m.OtpTimedrift
	}
	return 0
}

type CMsgClientUpdateMachineAuthResponse struct {
	Filename         *string `protobuf:"bytes,1,opt,name=filename" json:"filename,omitempty"`
	Eresult          *uint32 `protobuf:"varint,2,opt,name=eresult" json:"eresult,omitempty"`
	Filesize         *uint32 `protobuf:"varint,3,opt,name=filesize" json:"filesize,omitempty"`
	ShaFile          []byte  `protobuf:"bytes,4,opt,name=sha_file" json:"sha_file,omitempty"`
	Getlasterror     *uint32 `protobuf:"varint,5,opt,name=getlasterror" json:"getlasterror,omitempty"`
	Offset           *uint32 `protobuf:"varint,6,opt,name=offset" json:"offset,omitempty"`
	Cubwrote         *uint32 `protobuf:"varint,7,opt,name=cubwrote" json:"cubwrote,omitempty"`
	OtpType          *int32  `protobuf:"varint,8,opt,name=otp_type" json:"otp_type,omitempty"`
	OtpValue         *uint32 `protobuf:"varint,9,opt,name=otp_value" json:"otp_value,omitempty"`
	OtpIdentifier    *string `protobuf:"bytes,10,opt,name=otp_identifier" json:"otp_identifier,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CMsgClientUpdateMachineAuthResponse) Reset()         { *m = CMsgClientUpdateMachineAuthResponse{} }
func (m *CMsgClientUpdateMachineAuthResponse) String() string { return proto.CompactTextString(m) }
func (*CMsgClientUpdateMachineAuthResponse) ProtoMessage()    {}

func (m *CMsgClientUpdateMachineAuthResponse) GetFilename() string {
	if m != nil && m.Filename != nil {
		return *m.Filename
	}
	return ""
}

func (m *CMsgClientUpdateMachineAuthResponse) GetEresult() uint32 {
	if m != nil && m.Eresult != nil {
		return *m.Eresult
	}
	return 0
}

func (m *CMsgClientUpdateMachineAuthResponse) GetFilesize() uint32 {
	if m != nil && m.Filesize != nil {
		return *m.Filesize
	}
	return 0
}

func (m *CMsgClientUpdateMachineAuthResponse) GetShaFile() []byte {
	if m != nil {
		return m.ShaFile
	}
	return nil
}

func (m *CMsgClientUpdateMachineAuthResponse) GetGetlasterror() uint32 {
	if m != nil && m.Getlasterror != nil {
		return *m.Getlasterror
	}
	return 0
}

func (m *CMsgClientUpdateMachineAuthResponse) GetOffset() uint32 {
	if m != nil && m.Offset != nil {
		return *m.Offset
	}
	return 0
}

func (m *CMsgClientUpdateMachineAuthResponse) GetCubwrote() uint32 {
	if m != nil && m.Cubwrote != nil {
		return *m.Cubwrote
	}
	return 0
}

func (m *CMsgClientUpdateMachineAuthResponse) GetOtpType() int32 {
	if m != nil && m.OtpType != nil {
		return *m.OtpType
	}
	return 0
}

func (m *CMsgClientUpdateMachineAuthResponse) GetOtpValue() uint32 {
	if m != nil && m.OtpValue != nil {
		return *m.OtpValue
	}
	return 0
}

func (m *CMsgClientUpdateMachineAuthResponse) GetOtpIdentifier() string {
	if m != nil && m.OtpIdentifier != nil {
		return *m.OtpIdentifier
	}
	return ""
}

type CMsgClientRequestMachineAuth struct {
	Filename              *string `protobuf:"bytes,1,opt,name=filename" json:"filename,omitempty"`
	EresultSentryfile     *uint32 `protobuf:"varint,2,opt,name=eresult_sentryfile" json:"eresult_sentryfile,omitempty"`
	Filesize              *uint32 `protobuf:"varint,3,opt,name=filesize" json:"filesize,omitempty"`
	ShaSentryfile         []byte  `protobuf:"bytes,4,opt,name=sha_sentryfile" json:"sha_sentryfile,omitempty"`
	LockAccountAction     *int32  `protobuf:"varint,6,opt,name=lock_account_action" json:"lock_account_action,omitempty"`
	OtpType               *uint32 `protobuf:"varint,7,opt,name=otp_type" json:"otp_type,omitempty"`
	OtpIdentifier         *string `protobuf:"bytes,8,opt,name=otp_identifier" json:"otp_identifier,omitempty"`
	OtpSharedsecret       []byte  `protobuf:"bytes,9,opt,name=otp_sharedsecret" json:"otp_sharedsecret,omitempty"`
	OtpValue              *uint32 `protobuf:"varint,10,opt,name=otp_value" json:"otp_value,omitempty"`
	MachineName           *string `protobuf:"bytes,11,opt,name=machine_name" json:"machine_name,omitempty"`
	MachineNameUserchosen *string `protobuf:"bytes,12,opt,name=machine_name_userchosen" json:"machine_name_userchosen,omitempty"`
	XXX_unrecognized      []byte  `json:"-"`
}

func (m *CMsgClientRequestMachineAuth) Reset()         { *m = CMsgClientRequestMachineAuth{} }
func (m *CMsgClientRequestMachineAuth) String() string { return proto.CompactTextString(m) }
func (*CMsgClientRequestMachineAuth) ProtoMessage()    {}

func (m *CMsgClientRequestMachineAuth) GetFilename() string {
	if m != nil && m.Filename != nil {
		return *m.Filename
	}
	return ""
}

func (m *CMsgClientRequestMachineAuth) GetEresultSentryfile() uint32 {
	if m != nil && m.EresultSentryfile != nil {
		return *m.EresultSentryfile
	}
	return 0
}

func (m *CMsgClientRequestMachineAuth) GetFilesize() uint32 {
	if m != nil && m.Filesize != nil {
		return *m.Filesize
	}
	return 0
}

func (m *CMsgClientRequestMachineAuth) GetShaSentryfile() []byte {
	if m != nil {
		return m.ShaSentryfile
	}
	return nil
}

func (m *CMsgClientRequestMachineAuth) GetLockAccountAction() int32 {
	if m != nil && m.LockAccountAction != nil {
		return *m.LockAccountAction
	}
	return 0
}

func (m *CMsgClientRequestMachineAuth) GetOtpType() uint32 {
	if m != nil && m.OtpType != nil {
		return *m.OtpType
	}
	return 0
}

func (m *CMsgClientRequestMachineAuth) GetOtpIdentifier() string {
	if m != nil && m.OtpIdentifier != nil {
		return *m.OtpIdentifier
	}
	return ""
}

func (m *CMsgClientRequestMachineAuth) GetOtpSharedsecret() []byte {
	if m != nil {
		return m.OtpSharedsecret
	}
	return nil
}

func (m *CMsgClientRequestMachineAuth) GetOtpValue() uint32 {
	if m != nil && m.OtpValue != nil {
		return *m.OtpValue
	}
	return 0
}

func (m *CMsgClientRequestMachineAuth) GetMachineName() string {
	if m != nil && m.MachineName != nil {
		return *m.MachineName
	}
	return ""
}

func (m *CMsgClientRequestMachineAuth) GetMachineNameUserchosen() string {
	if m != nil && m.MachineNameUserchosen != nil {
		return *m.MachineNameUserchosen
	}
	return ""
}

type CMsgClientRequestMachineAuthResponse struct {
	Eresult          *uint32 `protobuf:"varint,1,opt,name=eresult" json:"eresult,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CMsgClientRequestMachineAuthResponse) Reset()         { *m = CMsgClientRequestMachineAuthResponse{} }
func (m *CMsgClientRequestMachineAuthResponse) String() string { return proto.CompactTextString(m) }
func (*CMsgClientRequestMachineAuthResponse) ProtoMessage()    {}

func (m *CMsgClientRequestMachineAuthResponse) GetEresult() uint32 {
	if m != nil && m.Eresult != nil {
		return *m.Eresult
	}
	return 0
}
