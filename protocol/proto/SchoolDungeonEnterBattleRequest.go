// Code auto-generated by gucooing & Hiro420 DO NOT EDIT.

package proto

import (
	"github.com/bytedance/sonic"
	"github.com/gucooing/BaPs/protocol/mx"
)

type SchoolDungeonEnterBattleRequest struct {
	*RequestPacket

	StageUniqueId int64
}

func (x *SchoolDungeonEnterBattleRequest) String() string {
	j, _ := sonic.MarshalString(x)
	return j
}

func (x *SchoolDungeonEnterBattleRequest) SetPacket(packet mx.Message) {
	if x == nil {
		return
	}
	x.RequestPacket = packet.(*RequestPacket)
}
