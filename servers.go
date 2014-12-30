package gosteam

import (
	"bufio"
	"fmt"
	. "github.com/smithfox/gosteam/apath"
	"github.com/smithfox/gosteam/netutil"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var gCMServers []*netutil.PortAddr = []*netutil.PortAddr{}

func LoadCMServers() error {
	ccs := []*netutil.PortAddr{}

	fpath := filepath.Join(AppPath(), "bots", "servers.txt")

	f, err := os.OpenFile(fpath, os.O_RDONLY, 0660)
	if err != nil {
		fmt.Printf("open file:%s, err=%v\n", fpath, err)
		return err
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.Replace(line, " ", "", 0)
		line = strings.Replace(line, ",", "", 0)
		line = strings.Replace(line, "\"", "", 0)
		if line == "" {
			continue
		}
		if strings.HasPrefix(line, "#") {
			continue
		}
		if strings.HasPrefix(line, "//") {
			continue
		}
		addr := netutil.ParsePortAddr(line)
		if addr != nil {
			ccs = append(ccs, addr)
		} else {
			return fmt.Errorf("invalid server line:%s\n", line)
		}
	}
	gCMServers = ccs
	return scanner.Err()
}

func UpdateCMServers(ss []*netutil.PortAddr) {
	if len(ss) > 0 {
		gCMServers = ss
	}
}

func ShowCMServers() {
	for _, s := range gCMServers {
		fmt.Println(s.String())
	}
}

func GetRandomCM() *netutil.PortAddr {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	addr := gCMServers[rng.Int31n(int32(len(gCMServers)))]

	if addr == nil {
		panic("invalid address in CMServers slice")
	}
	return addr
}
