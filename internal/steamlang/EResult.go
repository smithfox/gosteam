package steamlang

import (
	"fmt"
	"sort"
	"strings"
)

type EResult int32

const (
	EResult_Invalid                                 EResult = 0
	EResult_OK                                      EResult = 1
	EResult_Fail                                    EResult = 2
	EResult_NoConnection                            EResult = 3
	EResult_InvalidPassword                         EResult = 5
	EResult_LoggedInElsewhere                       EResult = 6
	EResult_InvalidProtocolVer                      EResult = 7
	EResult_InvalidParam                            EResult = 8
	EResult_FileNotFound                            EResult = 9
	EResult_Busy                                    EResult = 10
	EResult_InvalidState                            EResult = 11
	EResult_InvalidName                             EResult = 12
	EResult_InvalidEmail                            EResult = 13
	EResult_DuplicateName                           EResult = 14
	EResult_AccessDenied                            EResult = 15
	EResult_Timeout                                 EResult = 16
	EResult_Banned                                  EResult = 17
	EResult_AccountNotFound                         EResult = 18
	EResult_InvalidSteamID                          EResult = 19
	EResult_ServiceUnavailable                      EResult = 20
	EResult_NotLoggedOn                             EResult = 21
	EResult_Pending                                 EResult = 22
	EResult_EncryptionFailure                       EResult = 23
	EResult_InsufficientPrivilege                   EResult = 24
	EResult_LimitExceeded                           EResult = 25
	EResult_Revoked                                 EResult = 26
	EResult_Expired                                 EResult = 27
	EResult_AlreadyRedeemed                         EResult = 28
	EResult_DuplicateRequest                        EResult = 29
	EResult_AlreadyOwned                            EResult = 30
	EResult_IPNotFound                              EResult = 31
	EResult_PersistFailed                           EResult = 32
	EResult_LockingFailed                           EResult = 33
	EResult_LogonSessionReplaced                    EResult = 34
	EResult_ConnectFailed                           EResult = 35
	EResult_HandshakeFailed                         EResult = 36
	EResult_IOFailure                               EResult = 37
	EResult_RemoteDisconnect                        EResult = 38
	EResult_ShoppingCartNotFound                    EResult = 39
	EResult_Blocked                                 EResult = 40
	EResult_Ignored                                 EResult = 41
	EResult_NoMatch                                 EResult = 42
	EResult_AccountDisabled                         EResult = 43
	EResult_ServiceReadOnly                         EResult = 44
	EResult_AccountNotFeatured                      EResult = 45
	EResult_AdministratorOK                         EResult = 46
	EResult_ContentVersion                          EResult = 47
	EResult_TryAnotherCM                            EResult = 48
	EResult_PasswordRequiredToKickSession           EResult = 49
	EResult_AlreadyLoggedInElsewhere                EResult = 50
	EResult_Suspended                               EResult = 51
	EResult_Cancelled                               EResult = 52
	EResult_DataCorruption                          EResult = 53
	EResult_DiskFull                                EResult = 54
	EResult_RemoteCallFailed                        EResult = 55
	EResult_PasswordNotSet                          EResult = 56
	EResult_ExternalAccountUnlinked                 EResult = 57
	EResult_PSNTicketInvalid                        EResult = 58
	EResult_ExternalAccountAlreadyLinked            EResult = 59
	EResult_RemoteFileConflict                      EResult = 60
	EResult_IllegalPassword                         EResult = 61
	EResult_SameAsPreviousValue                     EResult = 62
	EResult_AccountLogonDenied                      EResult = 63
	EResult_CannotUseOldPassword                    EResult = 64
	EResult_InvalidLoginAuthCode                    EResult = 65
	EResult_AccountLogonDeniedNoMailSent            EResult = 66
	EResult_HardwareNotCapableOfIPT                 EResult = 67
	EResult_IPTInitError                            EResult = 68
	EResult_ParentalControlRestricted               EResult = 69
	EResult_FacebookQueryError                      EResult = 70
	EResult_ExpiredLoginAuthCode                    EResult = 71
	EResult_IPLoginRestrictionFailed                EResult = 72
	EResult_AccountLocked                           EResult = 73
	EResult_AccountLogonDeniedVerifiedEmailRequired EResult = 74
	EResult_NoMatchingURL                           EResult = 75
	EResult_BadResponse                             EResult = 76
	EResult_RequirePasswordReEntry                  EResult = 77
	EResult_ValueOutOfRange                         EResult = 78
	EResult_UnexpectedError                         EResult = 79
	EResult_Disabled                                EResult = 80
	EResult_InvalidCEGSubmission                    EResult = 81
	EResult_RestrictedDevice                        EResult = 82
	EResult_RegionLocked                            EResult = 83
	EResult_RateLimitExceeded                       EResult = 84
	EResult_AccountLogonDeniedNeedTwoFactorCode     EResult = 85
	EResult_ItemOrEntryHasBeenDeleted               EResult = 86
)

var EResult_name = map[EResult]string{
	0:  "EResult_Invalid",
	1:  "EResult_OK",
	2:  "EResult_Fail",
	3:  "EResult_NoConnection",
	5:  "EResult_InvalidPassword",
	6:  "EResult_LoggedInElsewhere",
	7:  "EResult_InvalidProtocolVer",
	8:  "EResult_InvalidParam",
	9:  "EResult_FileNotFound",
	10: "EResult_Busy",
	11: "EResult_InvalidState",
	12: "EResult_InvalidName",
	13: "EResult_InvalidEmail",
	14: "EResult_DuplicateName",
	15: "EResult_AccessDenied",
	16: "EResult_Timeout",
	17: "EResult_Banned",
	18: "EResult_AccountNotFound",
	19: "EResult_InvalidSteamID",
	20: "EResult_ServiceUnavailable",
	21: "EResult_NotLoggedOn",
	22: "EResult_Pending",
	23: "EResult_EncryptionFailure",
	24: "EResult_InsufficientPrivilege",
	25: "EResult_LimitExceeded",
	26: "EResult_Revoked",
	27: "EResult_Expired",
	28: "EResult_AlreadyRedeemed",
	29: "EResult_DuplicateRequest",
	30: "EResult_AlreadyOwned",
	31: "EResult_IPNotFound",
	32: "EResult_PersistFailed",
	33: "EResult_LockingFailed",
	34: "EResult_LogonSessionReplaced",
	35: "EResult_ConnectFailed",
	36: "EResult_HandshakeFailed",
	37: "EResult_IOFailure",
	38: "EResult_RemoteDisconnect",
	39: "EResult_ShoppingCartNotFound",
	40: "EResult_Blocked",
	41: "EResult_Ignored",
	42: "EResult_NoMatch",
	43: "EResult_AccountDisabled",
	44: "EResult_ServiceReadOnly",
	45: "EResult_AccountNotFeatured",
	46: "EResult_AdministratorOK",
	47: "EResult_ContentVersion",
	48: "EResult_TryAnotherCM",
	49: "EResult_PasswordRequiredToKickSession",
	50: "EResult_AlreadyLoggedInElsewhere",
	51: "EResult_Suspended",
	52: "EResult_Cancelled",
	53: "EResult_DataCorruption",
	54: "EResult_DiskFull",
	55: "EResult_RemoteCallFailed",
	56: "EResult_PasswordNotSet",
	57: "EResult_ExternalAccountUnlinked",
	58: "EResult_PSNTicketInvalid",
	59: "EResult_ExternalAccountAlreadyLinked",
	60: "EResult_RemoteFileConflict",
	61: "EResult_IllegalPassword",
	62: "EResult_SameAsPreviousValue",
	63: "EResult_AccountLogonDenied",
	64: "EResult_CannotUseOldPassword",
	65: "EResult_InvalidLoginAuthCode",
	66: "EResult_AccountLogonDeniedNoMailSent",
	67: "EResult_HardwareNotCapableOfIPT",
	68: "EResult_IPTInitError",
	69: "EResult_ParentalControlRestricted",
	70: "EResult_FacebookQueryError",
	71: "EResult_ExpiredLoginAuthCode",
	72: "EResult_IPLoginRestrictionFailed",
	73: "EResult_AccountLocked",
	74: "EResult_AccountLogonDeniedVerifiedEmailRequired",
	75: "EResult_NoMatchingURL",
	76: "EResult_BadResponse",
	77: "EResult_RequirePasswordReEntry",
	78: "EResult_ValueOutOfRange",
	79: "EResult_UnexpectedError",
	80: "EResult_Disabled",
	81: "EResult_InvalidCEGSubmission",
	82: "EResult_RestrictedDevice",
	83: "EResult_RegionLocked",
	84: "EResult_RateLimitExceeded",
	85: "EResult_AccountLogonDeniedNeedTwoFactorCode",
	86: "EResult_ItemOrEntryHasBeenDeleted",
}

func (e EResult) String() string {
	if s, ok := EResult_name[e]; ok {
		return s
	}
	var flags []string
	for k, v := range EResult_name {
		if e&k != 0 {
			flags = append(flags, v)
		}
	}
	if len(flags) == 0 {
		return fmt.Sprintf("%d", e)
	}
	sort.Strings(flags)
	return strings.Join(flags, " | ")
}
