package protobuf

import (
	"github.com/golang/protobuf/proto"
)

type CMsgClientLogon struct {
	ProtocolVersion                   *uint32 `protobuf:"varint,1,opt,name=protocol_version" json:"protocol_version,omitempty"`
	ObfustucatedPrivateIp             *uint32 `protobuf:"varint,2,opt,name=obfustucated_private_ip" json:"obfustucated_private_ip,omitempty"`
	CellId                            *uint32 `protobuf:"varint,3,opt,name=cell_id" json:"cell_id,omitempty"`
	LastSessionId                     *uint32 `protobuf:"varint,4,opt,name=last_session_id" json:"last_session_id,omitempty"`
	ClientPackageVersion              *uint32 `protobuf:"varint,5,opt,name=client_package_version" json:"client_package_version,omitempty"`
	ClientLanguage                    *string `protobuf:"bytes,6,opt,name=client_language" json:"client_language,omitempty"`
	ClientOsType                      *uint32 `protobuf:"varint,7,opt,name=client_os_type" json:"client_os_type,omitempty"`
	ShouldRememberPassword            *bool   `protobuf:"varint,8,opt,name=should_remember_password,def=0" json:"should_remember_password,omitempty"`
	WineVersion                       *string `protobuf:"bytes,9,opt,name=wine_version" json:"wine_version,omitempty"`
	PingMsFromCellSearch              *uint32 `protobuf:"varint,10,opt,name=ping_ms_from_cell_search" json:"ping_ms_from_cell_search,omitempty"`
	PublicIp                          *uint32 `protobuf:"varint,20,opt,name=public_ip" json:"public_ip,omitempty"`
	QosLevel                          *uint32 `protobuf:"varint,21,opt,name=qos_level" json:"qos_level,omitempty"`
	ClientSuppliedSteamId             *uint64 `protobuf:"fixed64,22,opt,name=client_supplied_steam_id" json:"client_supplied_steam_id,omitempty"`
	MachineId                         []byte  `protobuf:"bytes,30,opt,name=machine_id" json:"machine_id,omitempty"`
	LauncherType                      *uint32 `protobuf:"varint,31,opt,name=launcher_type,def=0" json:"launcher_type,omitempty"`
	UiMode                            *uint32 `protobuf:"varint,32,opt,name=ui_mode,def=0" json:"ui_mode,omitempty"`
	Steam2AuthTicket                  []byte  `protobuf:"bytes,41,opt,name=steam2_auth_ticket" json:"steam2_auth_ticket,omitempty"`
	EmailAddress                      *string `protobuf:"bytes,42,opt,name=email_address" json:"email_address,omitempty"`
	Rtime32AccountCreation            *uint32 `protobuf:"fixed32,43,opt,name=rtime32_account_creation" json:"rtime32_account_creation,omitempty"`
	AccountName                       *string `protobuf:"bytes,50,opt,name=account_name" json:"account_name,omitempty"`
	Password                          *string `protobuf:"bytes,51,opt,name=password" json:"password,omitempty"`
	GameServerToken                   *string `protobuf:"bytes,52,opt,name=game_server_token" json:"game_server_token,omitempty"`
	LoginKey                          *string `protobuf:"bytes,60,opt,name=login_key" json:"login_key,omitempty"`
	WasConvertedDeprecatedMsg         *bool   `protobuf:"varint,70,opt,name=was_converted_deprecated_msg,def=0" json:"was_converted_deprecated_msg,omitempty"`
	AnonUserTargetAccountName         *string `protobuf:"bytes,80,opt,name=anon_user_target_account_name" json:"anon_user_target_account_name,omitempty"`
	ResolvedUserSteamId               *uint64 `protobuf:"fixed64,81,opt,name=resolved_user_steam_id" json:"resolved_user_steam_id,omitempty"`
	EresultSentryfile                 *int32  `protobuf:"varint,82,opt,name=eresult_sentryfile" json:"eresult_sentryfile,omitempty"`
	ShaSentryfile                     []byte  `protobuf:"bytes,83,opt,name=sha_sentryfile" json:"sha_sentryfile,omitempty"`
	AuthCode                          *string `protobuf:"bytes,84,opt,name=auth_code" json:"auth_code,omitempty"`
	OtpType                           *int32  `protobuf:"varint,85,opt,name=otp_type" json:"otp_type,omitempty"`
	OtpValue                          *uint32 `protobuf:"varint,86,opt,name=otp_value" json:"otp_value,omitempty"`
	OtpIdentifier                     *string `protobuf:"bytes,87,opt,name=otp_identifier" json:"otp_identifier,omitempty"`
	Steam2TicketRequest               *bool   `protobuf:"varint,88,opt,name=steam2_ticket_request" json:"steam2_ticket_request,omitempty"`
	SonyPsnTicket                     []byte  `protobuf:"bytes,90,opt,name=sony_psn_ticket" json:"sony_psn_ticket,omitempty"`
	SonyPsnServiceId                  *string `protobuf:"bytes,91,opt,name=sony_psn_service_id" json:"sony_psn_service_id,omitempty"`
	CreateNewPsnLinkedAccountIfNeeded *bool   `protobuf:"varint,92,opt,name=create_new_psn_linked_account_if_needed,def=0" json:"create_new_psn_linked_account_if_needed,omitempty"`
	SonyPsnName                       *string `protobuf:"bytes,93,opt,name=sony_psn_name" json:"sony_psn_name,omitempty"`
	GameServerAppId                   *int32  `protobuf:"varint,94,opt,name=game_server_app_id" json:"game_server_app_id,omitempty"`
	SteamguardDontRememberComputer    *bool   `protobuf:"varint,95,opt,name=steamguard_dont_remember_computer" json:"steamguard_dont_remember_computer,omitempty"`
	MachineName                       *string `protobuf:"bytes,96,opt,name=machine_name" json:"machine_name,omitempty"`
	MachineNameUserchosen             *string `protobuf:"bytes,97,opt,name=machine_name_userchosen" json:"machine_name_userchosen,omitempty"`
	CountryOverride                   *string `protobuf:"bytes,98,opt,name=country_override" json:"country_override,omitempty"`
	IsSteamBox                        *bool   `protobuf:"varint,99,opt,name=is_steam_box" json:"is_steam_box,omitempty"`
	ClientInstanceId                  *uint64 `protobuf:"varint,100,opt,name=client_instance_id" json:"client_instance_id,omitempty"`
	TwoFactorCode                     *string `protobuf:"bytes,101,opt,name=two_factor_code" json:"two_factor_code,omitempty"`
	XXX_unrecognized                  []byte  `json:"-"`
}

func (m *CMsgClientLogon) Reset()         { *m = CMsgClientLogon{} }
func (m *CMsgClientLogon) String() string { return proto.CompactTextString(m) }
func (*CMsgClientLogon) ProtoMessage()    {}

const Default_CMsgClientLogon_ShouldRememberPassword bool = false
const Default_CMsgClientLogon_LauncherType uint32 = 0
const Default_CMsgClientLogon_UiMode uint32 = 0
const Default_CMsgClientLogon_WasConvertedDeprecatedMsg bool = false
const Default_CMsgClientLogon_CreateNewPsnLinkedAccountIfNeeded bool = false

func (m *CMsgClientLogon) GetProtocolVersion() uint32 {
	if m != nil && m.ProtocolVersion != nil {
		return *m.ProtocolVersion
	}
	return 0
}

func (m *CMsgClientLogon) GetObfustucatedPrivateIp() uint32 {
	if m != nil && m.ObfustucatedPrivateIp != nil {
		return *m.ObfustucatedPrivateIp
	}
	return 0
}

func (m *CMsgClientLogon) GetCellId() uint32 {
	if m != nil && m.CellId != nil {
		return *m.CellId
	}
	return 0
}

func (m *CMsgClientLogon) GetLastSessionId() uint32 {
	if m != nil && m.LastSessionId != nil {
		return *m.LastSessionId
	}
	return 0
}

func (m *CMsgClientLogon) GetClientPackageVersion() uint32 {
	if m != nil && m.ClientPackageVersion != nil {
		return *m.ClientPackageVersion
	}
	return 0
}

func (m *CMsgClientLogon) GetClientLanguage() string {
	if m != nil && m.ClientLanguage != nil {
		return *m.ClientLanguage
	}
	return ""
}

func (m *CMsgClientLogon) GetClientOsType() uint32 {
	if m != nil && m.ClientOsType != nil {
		return *m.ClientOsType
	}
	return 0
}

func (m *CMsgClientLogon) GetShouldRememberPassword() bool {
	if m != nil && m.ShouldRememberPassword != nil {
		return *m.ShouldRememberPassword
	}
	return Default_CMsgClientLogon_ShouldRememberPassword
}

func (m *CMsgClientLogon) GetWineVersion() string {
	if m != nil && m.WineVersion != nil {
		return *m.WineVersion
	}
	return ""
}

func (m *CMsgClientLogon) GetPingMsFromCellSearch() uint32 {
	if m != nil && m.PingMsFromCellSearch != nil {
		return *m.PingMsFromCellSearch
	}
	return 0
}

func (m *CMsgClientLogon) GetPublicIp() uint32 {
	if m != nil && m.PublicIp != nil {
		return *m.PublicIp
	}
	return 0
}

func (m *CMsgClientLogon) GetQosLevel() uint32 {
	if m != nil && m.QosLevel != nil {
		return *m.QosLevel
	}
	return 0
}

func (m *CMsgClientLogon) GetClientSuppliedSteamId() uint64 {
	if m != nil && m.ClientSuppliedSteamId != nil {
		return *m.ClientSuppliedSteamId
	}
	return 0
}

func (m *CMsgClientLogon) GetMachineId() []byte {
	if m != nil {
		return m.MachineId
	}
	return nil
}

func (m *CMsgClientLogon) GetLauncherType() uint32 {
	if m != nil && m.LauncherType != nil {
		return *m.LauncherType
	}
	return Default_CMsgClientLogon_LauncherType
}

func (m *CMsgClientLogon) GetUiMode() uint32 {
	if m != nil && m.UiMode != nil {
		return *m.UiMode
	}
	return Default_CMsgClientLogon_UiMode
}

func (m *CMsgClientLogon) GetSteam2AuthTicket() []byte {
	if m != nil {
		return m.Steam2AuthTicket
	}
	return nil
}

func (m *CMsgClientLogon) GetEmailAddress() string {
	if m != nil && m.EmailAddress != nil {
		return *m.EmailAddress
	}
	return ""
}

func (m *CMsgClientLogon) GetRtime32AccountCreation() uint32 {
	if m != nil && m.Rtime32AccountCreation != nil {
		return *m.Rtime32AccountCreation
	}
	return 0
}

func (m *CMsgClientLogon) GetAccountName() string {
	if m != nil && m.AccountName != nil {
		return *m.AccountName
	}
	return ""
}

func (m *CMsgClientLogon) GetPassword() string {
	if m != nil && m.Password != nil {
		return *m.Password
	}
	return ""
}

func (m *CMsgClientLogon) GetGameServerToken() string {
	if m != nil && m.GameServerToken != nil {
		return *m.GameServerToken
	}
	return ""
}

func (m *CMsgClientLogon) GetLoginKey() string {
	if m != nil && m.LoginKey != nil {
		return *m.LoginKey
	}
	return ""
}

func (m *CMsgClientLogon) GetWasConvertedDeprecatedMsg() bool {
	if m != nil && m.WasConvertedDeprecatedMsg != nil {
		return *m.WasConvertedDeprecatedMsg
	}
	return Default_CMsgClientLogon_WasConvertedDeprecatedMsg
}

func (m *CMsgClientLogon) GetAnonUserTargetAccountName() string {
	if m != nil && m.AnonUserTargetAccountName != nil {
		return *m.AnonUserTargetAccountName
	}
	return ""
}

func (m *CMsgClientLogon) GetResolvedUserSteamId() uint64 {
	if m != nil && m.ResolvedUserSteamId != nil {
		return *m.ResolvedUserSteamId
	}
	return 0
}

func (m *CMsgClientLogon) GetEresultSentryfile() int32 {
	if m != nil && m.EresultSentryfile != nil {
		return *m.EresultSentryfile
	}
	return 0
}

func (m *CMsgClientLogon) GetShaSentryfile() []byte {
	if m != nil {
		return m.ShaSentryfile
	}
	return nil
}

func (m *CMsgClientLogon) GetAuthCode() string {
	if m != nil && m.AuthCode != nil {
		return *m.AuthCode
	}
	return ""
}

func (m *CMsgClientLogon) GetOtpType() int32 {
	if m != nil && m.OtpType != nil {
		return *m.OtpType
	}
	return 0
}

func (m *CMsgClientLogon) GetOtpValue() uint32 {
	if m != nil && m.OtpValue != nil {
		return *m.OtpValue
	}
	return 0
}

func (m *CMsgClientLogon) GetOtpIdentifier() string {
	if m != nil && m.OtpIdentifier != nil {
		return *m.OtpIdentifier
	}
	return ""
}

func (m *CMsgClientLogon) GetSteam2TicketRequest() bool {
	if m != nil && m.Steam2TicketRequest != nil {
		return *m.Steam2TicketRequest
	}
	return false
}

func (m *CMsgClientLogon) GetSonyPsnTicket() []byte {
	if m != nil {
		return m.SonyPsnTicket
	}
	return nil
}

func (m *CMsgClientLogon) GetSonyPsnServiceId() string {
	if m != nil && m.SonyPsnServiceId != nil {
		return *m.SonyPsnServiceId
	}
	return ""
}

func (m *CMsgClientLogon) GetCreateNewPsnLinkedAccountIfNeeded() bool {
	if m != nil && m.CreateNewPsnLinkedAccountIfNeeded != nil {
		return *m.CreateNewPsnLinkedAccountIfNeeded
	}
	return Default_CMsgClientLogon_CreateNewPsnLinkedAccountIfNeeded
}

func (m *CMsgClientLogon) GetSonyPsnName() string {
	if m != nil && m.SonyPsnName != nil {
		return *m.SonyPsnName
	}
	return ""
}

func (m *CMsgClientLogon) GetGameServerAppId() int32 {
	if m != nil && m.GameServerAppId != nil {
		return *m.GameServerAppId
	}
	return 0
}

func (m *CMsgClientLogon) GetSteamguardDontRememberComputer() bool {
	if m != nil && m.SteamguardDontRememberComputer != nil {
		return *m.SteamguardDontRememberComputer
	}
	return false
}

func (m *CMsgClientLogon) GetMachineName() string {
	if m != nil && m.MachineName != nil {
		return *m.MachineName
	}
	return ""
}

func (m *CMsgClientLogon) GetMachineNameUserchosen() string {
	if m != nil && m.MachineNameUserchosen != nil {
		return *m.MachineNameUserchosen
	}
	return ""
}

func (m *CMsgClientLogon) GetCountryOverride() string {
	if m != nil && m.CountryOverride != nil {
		return *m.CountryOverride
	}
	return ""
}

func (m *CMsgClientLogon) GetIsSteamBox() bool {
	if m != nil && m.IsSteamBox != nil {
		return *m.IsSteamBox
	}
	return false
}

func (m *CMsgClientLogon) GetClientInstanceId() uint64 {
	if m != nil && m.ClientInstanceId != nil {
		return *m.ClientInstanceId
	}
	return 0
}

func (m *CMsgClientLogon) GetTwoFactorCode() string {
	if m != nil && m.TwoFactorCode != nil {
		return *m.TwoFactorCode
	}
	return ""
}

type CMsgClientLogonResponse struct {
	Eresult                     *int32  `protobuf:"varint,1,opt,name=eresult,def=2" json:"eresult,omitempty"`
	OutOfGameHeartbeatSeconds   *int32  `protobuf:"varint,2,opt,name=out_of_game_heartbeat_seconds" json:"out_of_game_heartbeat_seconds,omitempty"`
	InGameHeartbeatSeconds      *int32  `protobuf:"varint,3,opt,name=in_game_heartbeat_seconds" json:"in_game_heartbeat_seconds,omitempty"`
	PublicIp                    *uint32 `protobuf:"varint,4,opt,name=public_ip" json:"public_ip,omitempty"`
	Rtime32ServerTime           *uint32 `protobuf:"fixed32,5,opt,name=rtime32_server_time" json:"rtime32_server_time,omitempty"`
	AccountFlags                *uint32 `protobuf:"varint,6,opt,name=account_flags" json:"account_flags,omitempty"`
	CellId                      *uint32 `protobuf:"varint,7,opt,name=cell_id" json:"cell_id,omitempty"`
	EmailDomain                 *string `protobuf:"bytes,8,opt,name=email_domain" json:"email_domain,omitempty"`
	Steam2Ticket                []byte  `protobuf:"bytes,9,opt,name=steam2_ticket" json:"steam2_ticket,omitempty"`
	EresultExtended             *int32  `protobuf:"varint,10,opt,name=eresult_extended" json:"eresult_extended,omitempty"`
	WebapiAuthenticateUserNonce *string `protobuf:"bytes,11,opt,name=webapi_authenticate_user_nonce" json:"webapi_authenticate_user_nonce,omitempty"`
	CellIdPingThreshold         *uint32 `protobuf:"varint,12,opt,name=cell_id_ping_threshold" json:"cell_id_ping_threshold,omitempty"`
	UsePics                     *bool   `protobuf:"varint,13,opt,name=use_pics" json:"use_pics,omitempty"`
	VanityUrl                   *string `protobuf:"bytes,14,opt,name=vanity_url" json:"vanity_url,omitempty"`
	ClientSuppliedSteamid       *uint64 `protobuf:"fixed64,20,opt,name=client_supplied_steamid" json:"client_supplied_steamid,omitempty"`
	IpCountryCode               *string `protobuf:"bytes,21,opt,name=ip_country_code" json:"ip_country_code,omitempty"`
	ParentalSettings            []byte  `protobuf:"bytes,22,opt,name=parental_settings" json:"parental_settings,omitempty"`
	ParentalSettingSignature    []byte  `protobuf:"bytes,23,opt,name=parental_setting_signature" json:"parental_setting_signature,omitempty"`
	CountLoginfailuresToMigrate *int32  `protobuf:"varint,24,opt,name=count_loginfailures_to_migrate" json:"count_loginfailures_to_migrate,omitempty"`
	CountDisconnectsToMigrate   *int32  `protobuf:"varint,25,opt,name=count_disconnects_to_migrate" json:"count_disconnects_to_migrate,omitempty"`
	OgsDataReportTimeWindow     *int32  `protobuf:"varint,26,opt,name=ogs_data_report_time_window" json:"ogs_data_report_time_window,omitempty"`
	ClientInstanceId            *uint64 `protobuf:"varint,27,opt,name=client_instance_id" json:"client_instance_id,omitempty"`
	XXX_unrecognized            []byte  `json:"-"`
}

func (m *CMsgClientLogonResponse) Reset()         { *m = CMsgClientLogonResponse{} }
func (m *CMsgClientLogonResponse) String() string { return proto.CompactTextString(m) }
func (*CMsgClientLogonResponse) ProtoMessage()    {}

const Default_CMsgClientLogonResponse_Eresult int32 = 2

func (m *CMsgClientLogonResponse) GetEresult() int32 {
	if m != nil && m.Eresult != nil {
		return *m.Eresult
	}
	return Default_CMsgClientLogonResponse_Eresult
}

func (m *CMsgClientLogonResponse) GetOutOfGameHeartbeatSeconds() int32 {
	if m != nil && m.OutOfGameHeartbeatSeconds != nil {
		return *m.OutOfGameHeartbeatSeconds
	}
	return 0
}

func (m *CMsgClientLogonResponse) GetInGameHeartbeatSeconds() int32 {
	if m != nil && m.InGameHeartbeatSeconds != nil {
		return *m.InGameHeartbeatSeconds
	}
	return 0
}

func (m *CMsgClientLogonResponse) GetPublicIp() uint32 {
	if m != nil && m.PublicIp != nil {
		return *m.PublicIp
	}
	return 0
}

func (m *CMsgClientLogonResponse) GetRtime32ServerTime() uint32 {
	if m != nil && m.Rtime32ServerTime != nil {
		return *m.Rtime32ServerTime
	}
	return 0
}

func (m *CMsgClientLogonResponse) GetAccountFlags() uint32 {
	if m != nil && m.AccountFlags != nil {
		return *m.AccountFlags
	}
	return 0
}

func (m *CMsgClientLogonResponse) GetCellId() uint32 {
	if m != nil && m.CellId != nil {
		return *m.CellId
	}
	return 0
}

func (m *CMsgClientLogonResponse) GetEmailDomain() string {
	if m != nil && m.EmailDomain != nil {
		return *m.EmailDomain
	}
	return ""
}

func (m *CMsgClientLogonResponse) GetSteam2Ticket() []byte {
	if m != nil {
		return m.Steam2Ticket
	}
	return nil
}

func (m *CMsgClientLogonResponse) GetEresultExtended() int32 {
	if m != nil && m.EresultExtended != nil {
		return *m.EresultExtended
	}
	return 0
}

func (m *CMsgClientLogonResponse) GetWebapiAuthenticateUserNonce() string {
	if m != nil && m.WebapiAuthenticateUserNonce != nil {
		return *m.WebapiAuthenticateUserNonce
	}
	return ""
}

func (m *CMsgClientLogonResponse) GetCellIdPingThreshold() uint32 {
	if m != nil && m.CellIdPingThreshold != nil {
		return *m.CellIdPingThreshold
	}
	return 0
}

func (m *CMsgClientLogonResponse) GetUsePics() bool {
	if m != nil && m.UsePics != nil {
		return *m.UsePics
	}
	return false
}

func (m *CMsgClientLogonResponse) GetVanityUrl() string {
	if m != nil && m.VanityUrl != nil {
		return *m.VanityUrl
	}
	return ""
}

func (m *CMsgClientLogonResponse) GetClientSuppliedSteamid() uint64 {
	if m != nil && m.ClientSuppliedSteamid != nil {
		return *m.ClientSuppliedSteamid
	}
	return 0
}

func (m *CMsgClientLogonResponse) GetIpCountryCode() string {
	if m != nil && m.IpCountryCode != nil {
		return *m.IpCountryCode
	}
	return ""
}

func (m *CMsgClientLogonResponse) GetParentalSettings() []byte {
	if m != nil {
		return m.ParentalSettings
	}
	return nil
}

func (m *CMsgClientLogonResponse) GetParentalSettingSignature() []byte {
	if m != nil {
		return m.ParentalSettingSignature
	}
	return nil
}

func (m *CMsgClientLogonResponse) GetCountLoginfailuresToMigrate() int32 {
	if m != nil && m.CountLoginfailuresToMigrate != nil {
		return *m.CountLoginfailuresToMigrate
	}
	return 0
}

func (m *CMsgClientLogonResponse) GetCountDisconnectsToMigrate() int32 {
	if m != nil && m.CountDisconnectsToMigrate != nil {
		return *m.CountDisconnectsToMigrate
	}
	return 0
}

func (m *CMsgClientLogonResponse) GetOgsDataReportTimeWindow() int32 {
	if m != nil && m.OgsDataReportTimeWindow != nil {
		return *m.OgsDataReportTimeWindow
	}
	return 0
}

func (m *CMsgClientLogonResponse) GetClientInstanceId() uint64 {
	if m != nil && m.ClientInstanceId != nil {
		return *m.ClientInstanceId
	}
	return 0
}
