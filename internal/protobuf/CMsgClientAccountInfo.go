package protobuf

import (
	"github.com/golang/protobuf/proto"
)

type CMsgClientAccountInfo struct {
	PersonaName                     *string `protobuf:"bytes,1,opt,name=persona_name" json:"persona_name,omitempty"`
	IpCountry                       *string `protobuf:"bytes,2,opt,name=ip_country" json:"ip_country,omitempty"`
	SaltPassword                    []byte  `protobuf:"bytes,3,opt,name=salt_password" json:"salt_password,omitempty"`
	ShaDigest_Password              []byte  `protobuf:"bytes,4,opt,name=sha_digest_Password" json:"sha_digest_Password,omitempty"`
	CountAuthedComputers            *int32  `protobuf:"varint,5,opt,name=count_authed_computers" json:"count_authed_computers,omitempty"`
	LockedWithIpt                   *bool   `protobuf:"varint,6,opt,name=locked_with_ipt" json:"locked_with_ipt,omitempty"`
	AccountFlags                    *uint32 `protobuf:"varint,7,opt,name=account_flags" json:"account_flags,omitempty"`
	FacebookId                      *uint64 `protobuf:"varint,8,opt,name=facebook_id" json:"facebook_id,omitempty"`
	FacebookName                    *string `protobuf:"bytes,9,opt,name=facebook_name" json:"facebook_name,omitempty"`
	SteamGuardProvider              *int32  `protobuf:"varint,10,opt,name=steam_guard_provider" json:"steam_guard_provider,omitempty"`
	SteamguardRequireCodeDefault    *bool   `protobuf:"varint,11,opt,name=steamguard_require_code_default" json:"steamguard_require_code_default,omitempty"`
	SteamguardShowProviders         *bool   `protobuf:"varint,12,opt,name=steamguard_show_providers" json:"steamguard_show_providers,omitempty"`
	SteamguardCanUseMobileProvider  *bool   `protobuf:"varint,13,opt,name=steamguard_can_use_mobile_provider" json:"steamguard_can_use_mobile_provider,omitempty"`
	SteamguardNotifyNewmachines     *bool   `protobuf:"varint,14,opt,name=steamguard_notify_newmachines" json:"steamguard_notify_newmachines,omitempty"`
	SteamguardMachineNameUserChosen *string `protobuf:"bytes,15,opt,name=steamguard_machine_name_user_chosen" json:"steamguard_machine_name_user_chosen,omitempty"`
	XXX_unrecognized                []byte  `json:"-"`
}

func (m *CMsgClientAccountInfo) Reset()         { *m = CMsgClientAccountInfo{} }
func (m *CMsgClientAccountInfo) String() string { return proto.CompactTextString(m) }
func (*CMsgClientAccountInfo) ProtoMessage()    {}

func (m *CMsgClientAccountInfo) GetPersonaName() string {
	if m != nil && m.PersonaName != nil {
		return *m.PersonaName
	}
	return ""
}

func (m *CMsgClientAccountInfo) GetIpCountry() string {
	if m != nil && m.IpCountry != nil {
		return *m.IpCountry
	}
	return ""
}

func (m *CMsgClientAccountInfo) GetSaltPassword() []byte {
	if m != nil {
		return m.SaltPassword
	}
	return nil
}

func (m *CMsgClientAccountInfo) GetShaDigest_Password() []byte {
	if m != nil {
		return m.ShaDigest_Password
	}
	return nil
}

func (m *CMsgClientAccountInfo) GetCountAuthedComputers() int32 {
	if m != nil && m.CountAuthedComputers != nil {
		return *m.CountAuthedComputers
	}
	return 0
}

func (m *CMsgClientAccountInfo) GetLockedWithIpt() bool {
	if m != nil && m.LockedWithIpt != nil {
		return *m.LockedWithIpt
	}
	return false
}

func (m *CMsgClientAccountInfo) GetAccountFlags() uint32 {
	if m != nil && m.AccountFlags != nil {
		return *m.AccountFlags
	}
	return 0
}

func (m *CMsgClientAccountInfo) GetFacebookId() uint64 {
	if m != nil && m.FacebookId != nil {
		return *m.FacebookId
	}
	return 0
}

func (m *CMsgClientAccountInfo) GetFacebookName() string {
	if m != nil && m.FacebookName != nil {
		return *m.FacebookName
	}
	return ""
}

func (m *CMsgClientAccountInfo) GetSteamGuardProvider() int32 {
	if m != nil && m.SteamGuardProvider != nil {
		return *m.SteamGuardProvider
	}
	return 0
}

func (m *CMsgClientAccountInfo) GetSteamguardRequireCodeDefault() bool {
	if m != nil && m.SteamguardRequireCodeDefault != nil {
		return *m.SteamguardRequireCodeDefault
	}
	return false
}

func (m *CMsgClientAccountInfo) GetSteamguardShowProviders() bool {
	if m != nil && m.SteamguardShowProviders != nil {
		return *m.SteamguardShowProviders
	}
	return false
}

func (m *CMsgClientAccountInfo) GetSteamguardCanUseMobileProvider() bool {
	if m != nil && m.SteamguardCanUseMobileProvider != nil {
		return *m.SteamguardCanUseMobileProvider
	}
	return false
}

func (m *CMsgClientAccountInfo) GetSteamguardNotifyNewmachines() bool {
	if m != nil && m.SteamguardNotifyNewmachines != nil {
		return *m.SteamguardNotifyNewmachines
	}
	return false
}

func (m *CMsgClientAccountInfo) GetSteamguardMachineNameUserChosen() string {
	if m != nil && m.SteamguardMachineNameUserChosen != nil {
		return *m.SteamguardMachineNameUserChosen
	}
	return ""
}
