// Code auto-generated by gucooing & Hiro420 DO NOT EDIT.

package proto

import (
	"github.com/bytedance/sonic"
	"github.com/gucooing/BaPs/protocol/mx"
)

type ConquestMainStoryConquerWithBattleStartRequest struct {
	*RequestPacket

	EventContentId    int64
	Difficulty        StageDifficulty
	TileUniqueId      int64
	EchelonNumber     int64
	ClanAssistUseInfo *ClanAssistUseInfo
}

func (x *ConquestMainStoryConquerWithBattleStartRequest) String() string {
	j, _ := sonic.MarshalString(x)
	return j
}

func (x *ConquestMainStoryConquerWithBattleStartRequest) SetPacket(packet mx.Message) {
	if x == nil {
		return
	}
	x.RequestPacket = packet.(*RequestPacket)
}
