package steamlang

import (
	"fmt"
	"sort"
	"strings"
)

type EPersonaState int32

const (
	EPersonaState_Offline        EPersonaState = 0
	EPersonaState_Online         EPersonaState = 1
	EPersonaState_Busy           EPersonaState = 2
	EPersonaState_Away           EPersonaState = 3
	EPersonaState_Snooze         EPersonaState = 4
	EPersonaState_LookingToTrade EPersonaState = 5
	EPersonaState_LookingToPlay  EPersonaState = 6
	EPersonaState_Max            EPersonaState = 7
)

var EPersonaState_name = map[EPersonaState]string{
	0: "EPersonaState_Offline",
	1: "EPersonaState_Online",
	2: "EPersonaState_Busy",
	3: "EPersonaState_Away",
	4: "EPersonaState_Snooze",
	5: "EPersonaState_LookingToTrade",
	6: "EPersonaState_LookingToPlay",
	7: "EPersonaState_Max",
}

func (e EPersonaState) String() string {
	if s, ok := EPersonaState_name[e]; ok {
		return s
	}
	var flags []string
	for k, v := range EPersonaState_name {
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
