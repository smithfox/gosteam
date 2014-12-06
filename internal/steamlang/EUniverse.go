package steamlang

import (
	"fmt"
	"sort"
	"strings"
)

type EUniverse int32

const (
	EUniverse_Invalid  EUniverse = 0
	EUniverse_Public   EUniverse = 1
	EUniverse_Beta     EUniverse = 2
	EUniverse_Internal EUniverse = 3
	EUniverse_Dev      EUniverse = 4
	EUniverse_RC       EUniverse = 5 // Deprecated: Universe no longer exists
	EUniverse_Max      EUniverse = 5
)

var EUniverse_name = map[EUniverse]string{
	0: "EUniverse_Invalid",
	1: "EUniverse_Public",
	2: "EUniverse_Beta",
	3: "EUniverse_Internal",
	4: "EUniverse_Dev",
	5: "EUniverse_RC",
}

func (e EUniverse) String() string {
	if s, ok := EUniverse_name[e]; ok {
		return s
	}

	var flags []string
	for k, v := range EUniverse_name {
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
