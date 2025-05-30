// Code auto-generated by gucooing & Hiro420 DO NOT EDIT.

package proto

import (
	"github.com/bytedance/sonic"
	"github.com/gucooing/BaPs/protocol/mx"
)

type RaidEnterBattleRequest struct {
	*RequestPacket

	RaidServerId  int64
	RaidUniqueId  int64
	IsPractice    bool
	EchelonId     int64
	AssistUseInfo *ClanAssistUseInfo
}

func (x *RaidEnterBattleRequest) String() string {
	j, _ := sonic.MarshalString(x)
	return j
}

func (x *RaidEnterBattleRequest) SetPacket(packet mx.Message) {
	if x == nil {
		return
	}
	x.RequestPacket = packet.(*RequestPacket)
}
