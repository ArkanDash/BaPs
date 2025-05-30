// Code auto-generated by gucooing & Hiro420 DO NOT EDIT.

package proto

import (
	"github.com/bytedance/sonic"
	"github.com/gucooing/BaPs/protocol/mx"
)

type EliminateRaidEndBattleResponse struct {
	*ResponsePacket

	RankingPoint        int64
	BestRankingPoint    int64
	ClearTimePoint      int64
	HPPercentScorePoint int64
	DefaultClearPoint   int64
	ParcelResultDB      *ParcelResultDB
}

func (x *EliminateRaidEndBattleResponse) String() string {
	j, _ := sonic.MarshalString(x)
	return j
}

func (x *EliminateRaidEndBattleResponse) SetPacket(packet mx.Message) {
	if x == nil {
		return
	}
	x.ResponsePacket = packet.(*ResponsePacket)
}
