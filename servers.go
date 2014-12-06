package gosteam

import (
	"github.com/smithfox/gosteam/netutil"
	"math/rand"
	"time"
)

var northAmericanServers = []string{
	//North American Servers
	// Qwest, Seattle
	"72.165.61.174:27017",
	"72.165.61.174:27018",
	"72.165.61.175:27017",
	"72.165.61.175:27018",
	"72.165.61.176:27017",
	"72.165.61.176:27018",
	"72.165.61.185:27017",
	"72.165.61.185:27018",
	"72.165.61.187:27017",
	"72.165.61.187:27018",
	"72.165.61.188:27017",
	"72.165.61.188:27018",
	// Highwinds, Kaysville
	"209.197.29.196:27017",
	"209.197.29.197:27017",
}
var eruopeServers = []string{
	//Europe Servers
	// Inteliquent, Luxembourg, cm-[01-04].lux.valve.net
	"146.66.152.12:27017",
	"146.66.152.12:27018",
	"146.66.152.12:27019",
	"146.66.152.13:27017",
	"146.66.152.13:27018",
	"146.66.152.13:27019",
	"146.66.152.14:27017",
	"146.66.152.14:27018",
	"146.66.152.14:27019",
	"146.66.152.15:27017",
	"146.66.152.15:27018",
	"146.66.152.15:27019",
}
var otherServers = []string{
/* Highwinds, Netherlands (not live)
"81.171.115.5":27017",
"81.171.115.5":27018",
"81.171.115.5":27019",
"81.171.115.6":27017",
"81.171.115.6":27018",
"81.171.115.6":27019",
"81.171.115.7":27017",
"81.171.115.7":27018",
"81.171.115.7":27019",
"81.171.115.8":27017",
"81.171.115.8":27018",
"81.171.115.8":27019",*/
/* Starhub, Singapore (non-optimal route)
"103.28.54.10":27017",
"103.28.54.11":27017,*/
}

var gCMServers []*netutil.PortAddr = []*netutil.PortAddr{}

func InitCMServers() {
	ccs := []*netutil.PortAddr{}
	for _, cmserver := range northAmericanServers {
		addr := netutil.ParsePortAddr(cmserver)
		if addr != nil {
			ccs = append(ccs, addr)
		}
	}
	for _, cmserver := range eruopeServers {
		addr := netutil.ParsePortAddr(cmserver)
		if addr != nil {
			ccs = append(ccs, addr)
		}
	}
	for _, cmserver := range otherServers {
		addr := netutil.ParsePortAddr(cmserver)
		if addr != nil {
			ccs = append(ccs, addr)
		}
	}
	gCMServers = ccs
}

func UpdateCMServers(ss []*netutil.PortAddr) {
	if len(ss) > 0 {
		gCMServers = ss
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

/*
func GetRandomNorthAmericaCM() *netutil.PortAddr {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	addr := netutil.ParsePortAddr(northAmericanServers[rng.Int31n(int32(len(northAmericanServers)))])
	if addr == nil {
		panic("invalid address in CMServers slice")
	}
	return addr
}

func GetRandomEuropeCM() *netutil.PortAddr {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	addr := netutil.ParsePortAddr(eruopeServers[rng.Int31n(int32(len(eruopeServers)))])
	if addr == nil {
		panic("invalid address in CMServers slice")
	}
	return addr
}
*/
