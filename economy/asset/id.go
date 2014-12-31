package asset

import (
	. "github.com/smithfox/gosteam/steamid"
	"sync"
)

type AoidT struct {
	AssetId    int64
	OriginalId int64
	Did        int
	LastTime   int64
}

type SteamAoids struct {
	lock    sync.RWMutex
	steamid SteamId
	bya     map[int64]*AoidT
	byo     map[int64]*AoidT
}

func (m *SteamAoids) GetAssetIdByOid(oid int64) int64 {
	if oid == 0 {
		return 0
	}

	if m.byo == nil {
		m.lock.Lock()
		m.byo = make(map[int64]*AoidT)
		m.lock.Unlock()
	}

	m.lock.RLock()
	ss, _ := m.byo[oid]
	m.lock.RUnlock()

	if ss == nil {
		return 0
	} else {
		if ss.AssetId == 0 && ss.OriginalId == 0 {
			m.lock.Lock()
			delete(m.byo, oid)
			m.lock.Unlock()
			return 0
		}
		return ss.AssetId
	}
}

var gAllSteamAoidsLock sync.RWMutex
var gAllSteamAoids map[SteamId]*SteamAoids = make(map[SteamId]*SteamAoids)

func GetSteamAssetIdByOid(steamid SteamId, oid int64) int64 {
	if steamid == 0 || oid == 0 {
		return 0
	}
	var aid int64 = 0

	gAllSteamAoidsLock.RLock()
	u, _ := gAllSteamAoids[steamid]
	gAllSteamAoidsLock.RUnlock()

	if u != nil {
		aid = u.GetAssetIdByOid(oid)
	}

	return aid
}

func RemoveAsset(steamid SteamId, aid int64) {
	if steamid == 0 || aid == 0 {
		return
	}

	gAllSteamAoidsLock.Lock()
	u, _ := gAllSteamAoids[steamid]
	if u != nil {
		dd, _ := u.bya[aid]
		if dd != nil {
			dd.OriginalId = 0
			dd.AssetId = 0
			dd.Did = 0
			dd.LastTime = 0
		}
		delete(u.bya, aid)
	}
	gAllSteamAoidsLock.Unlock()
}

func AddAsset(steamid SteamId, aid int64) {
}

func UpateSteamAoids(steamid SteamId, lasttime int64, aoidis ...*AoidT) {
	if steamid == 0 || len(aoidis) == 0 {
		return
	}
	gAllSteamAoidsLock.Lock()
	_updateSteamAoids(steamid, lasttime, aoidis...)
	gAllSteamAoidsLock.Unlock()
}

func _updateSteamAoids(steamid SteamId, lasttime int64, aoidis ...*AoidT) {
	u, _ := gAllSteamAoids[steamid]

	if u == nil {
		u = &SteamAoids{
			steamid: steamid,
			bya:     make(map[int64]*AoidT),
			byo:     make(map[int64]*AoidT),
		}

		gAllSteamAoids[steamid] = u
	}

	for _, a := range aoidis {
		d, _ := u.bya[a.AssetId]
		if d == nil {
			u.bya[a.AssetId] = a
			u.byo[a.OriginalId] = a
		}
	}
}
