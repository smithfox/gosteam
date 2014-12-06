package steamlang

import (
	"fmt"
	"sort"
	"strings"
)

type EAccountType int32

const (
	EAccountType_Invalid        EAccountType = 0
	EAccountType_Individual     EAccountType = 1
	EAccountType_Multiseat      EAccountType = 2
	EAccountType_GameServer     EAccountType = 3
	EAccountType_AnonGameServer EAccountType = 4
	EAccountType_Pending        EAccountType = 5
	EAccountType_ContentServer  EAccountType = 6
	EAccountType_Clan           EAccountType = 7
	EAccountType_Chat           EAccountType = 8
	EAccountType_ConsoleUser    EAccountType = 9
	EAccountType_AnonUser       EAccountType = 10
	EAccountType_Max            EAccountType = 11
)

var EAccountType_name = map[EAccountType]string{
	0:  "EAccountType_Invalid",
	1:  "EAccountType_Individual",
	2:  "EAccountType_Multiseat",
	3:  "EAccountType_GameServer",
	4:  "EAccountType_AnonGameServer",
	5:  "EAccountType_Pending",
	6:  "EAccountType_ContentServer",
	7:  "EAccountType_Clan",
	8:  "EAccountType_Chat",
	9:  "EAccountType_ConsoleUser",
	10: "EAccountType_AnonUser",
	11: "EAccountType_Max",
}

func (e EAccountType) String() string {
	if s, ok := EAccountType_name[e]; ok {
		return s
	}
	var flags []string
	for k, v := range EAccountType_name {
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
