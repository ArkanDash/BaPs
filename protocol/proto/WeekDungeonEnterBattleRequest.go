// Code auto-generated by gucooing & Hiro420 DO NOT EDIT.

package proto

import (
	"github.com/bytedance/sonic"
	"github.com/gucooing/BaPs/protocol/mx"
)

type WeekDungeonEnterBattleRequest struct {
	*RequestPacket

	StageUniqueId int64
	EchelonIndex  int64
}

func (x *WeekDungeonEnterBattleRequest) String() string {
	j, _ := sonic.MarshalString(x)
	return j
}

func (x *WeekDungeonEnterBattleRequest) SetPacket(packet mx.Message) {
	if x == nil {
		return
	}
	x.RequestPacket = packet.(*RequestPacket)
}
