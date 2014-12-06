package steamlang

import (
	"fmt"
	"sort"
	"strings"
)

type EAccountFlags int32

const (
	EAccountFlags_NormalUser                 EAccountFlags = 0
	EAccountFlags_PersonaNameSet             EAccountFlags = 1
	EAccountFlags_Unbannable                 EAccountFlags = 2
	EAccountFlags_PasswordSet                EAccountFlags = 4
	EAccountFlags_Support                    EAccountFlags = 8
	EAccountFlags_Admin                      EAccountFlags = 16
	EAccountFlags_Supervisor                 EAccountFlags = 32
	EAccountFlags_AppEditor                  EAccountFlags = 64
	EAccountFlags_HWIDSet                    EAccountFlags = 128
	EAccountFlags_PersonalQASet              EAccountFlags = 256
	EAccountFlags_VacBeta                    EAccountFlags = 512
	EAccountFlags_Debug                      EAccountFlags = 1024
	EAccountFlags_Disabled                   EAccountFlags = 2048
	EAccountFlags_LimitedUser                EAccountFlags = 4096
	EAccountFlags_LimitedUserForce           EAccountFlags = 8192
	EAccountFlags_EmailValidated             EAccountFlags = 16384
	EAccountFlags_MarketingTreatment         EAccountFlags = 32768
	EAccountFlags_OGGInviteOptOut            EAccountFlags = 65536
	EAccountFlags_ForcePasswordChange        EAccountFlags = 131072
	EAccountFlags_ForceEmailVerification     EAccountFlags = 262144
	EAccountFlags_LogonExtraSecurity         EAccountFlags = 524288
	EAccountFlags_LogonExtraSecurityDisabled EAccountFlags = 1048576
	EAccountFlags_Steam2MigrationComplete    EAccountFlags = 2097152
	EAccountFlags_NeedLogs                   EAccountFlags = 4194304
	EAccountFlags_Lockdown                   EAccountFlags = 8388608
	EAccountFlags_MasterAppEditor            EAccountFlags = 16777216
	EAccountFlags_BannedFromWebAPI           EAccountFlags = 33554432
	EAccountFlags_ClansOnlyFromFriends       EAccountFlags = 67108864
	EAccountFlags_GlobalModerator            EAccountFlags = 134217728
)

var EAccountFlags_name = map[EAccountFlags]string{
	0:         "EAccountFlags_NormalUser",
	1:         "EAccountFlags_PersonaNameSet",
	2:         "EAccountFlags_Unbannable",
	4:         "EAccountFlags_PasswordSet",
	8:         "EAccountFlags_Support",
	16:        "EAccountFlags_Admin",
	32:        "EAccountFlags_Supervisor",
	64:        "EAccountFlags_AppEditor",
	128:       "EAccountFlags_HWIDSet",
	256:       "EAccountFlags_PersonalQASet",
	512:       "EAccountFlags_VacBeta",
	1024:      "EAccountFlags_Debug",
	2048:      "EAccountFlags_Disabled",
	4096:      "EAccountFlags_LimitedUser",
	8192:      "EAccountFlags_LimitedUserForce",
	16384:     "EAccountFlags_EmailValidated",
	32768:     "EAccountFlags_MarketingTreatment",
	65536:     "EAccountFlags_OGGInviteOptOut",
	131072:    "EAccountFlags_ForcePasswordChange",
	262144:    "EAccountFlags_ForceEmailVerification",
	524288:    "EAccountFlags_LogonExtraSecurity",
	1048576:   "EAccountFlags_LogonExtraSecurityDisabled",
	2097152:   "EAccountFlags_Steam2MigrationComplete",
	4194304:   "EAccountFlags_NeedLogs",
	8388608:   "EAccountFlags_Lockdown",
	16777216:  "EAccountFlags_MasterAppEditor",
	33554432:  "EAccountFlags_BannedFromWebAPI",
	67108864:  "EAccountFlags_ClansOnlyFromFriends",
	134217728: "EAccountFlags_GlobalModerator",
}

func (e EAccountFlags) String() string {
	if s, ok := EAccountFlags_name[e]; ok {
		return s
	}
	var flags []string
	for k, v := range EAccountFlags_name {
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
