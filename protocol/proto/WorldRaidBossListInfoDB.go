// Code auto-generated by gucooing & Hiro420 DO NOT EDIT.

package proto

import (
	"github.com/bytedance/sonic"
)

type WorldRaidBossListInfoDB struct {
	GroupId      int64
	WorldBossDB  *WorldRaidWorldBossDB
	LocalBossDBs []*WorldRaidLocalBossDB
}

func (x *WorldRaidBossListInfoDB) String() string {
	j, _ := sonic.MarshalString(x)
	return j
}
