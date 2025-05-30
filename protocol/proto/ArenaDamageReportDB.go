// Code auto-generated by gucooing & Hiro420 DO NOT EDIT.

package proto

import (
	"github.com/bytedance/sonic"
	"github.com/gucooing/BaPs/protocol/mx"
)

type ArenaDamageReportDB struct {
	ArenaBattleServerId   int64
	WinnerAccountServerId int64
	AttackerUserDB        *ArenaUserDB
	DefenderUserDB        *ArenaUserDB
	BattleEndTime         mx.MxTime
	AttackerDamageReport  map[int64]int64
	DefenderDamageReport  map[int64]int64
}

func (x *ArenaDamageReportDB) String() string {
	j, _ := sonic.MarshalString(x)
	return j
}
