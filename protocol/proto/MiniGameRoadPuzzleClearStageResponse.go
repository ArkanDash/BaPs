// Code auto-generated by gucooing & Hiro420 DO NOT EDIT.

package proto

import (
	"github.com/bytedance/sonic"
	"github.com/gucooing/BaPs/protocol/mx"
)

type MiniGameRoadPuzzleClearStageResponse struct {
	*ResponsePacket

	IsSkip         bool
	ParcelResultDB *ParcelResultDB
}

func (x *MiniGameRoadPuzzleClearStageResponse) String() string {
	j, _ := sonic.MarshalString(x)
	return j
}

func (x *MiniGameRoadPuzzleClearStageResponse) SetPacket(packet mx.Message) {
	if x == nil {
		return
	}
	x.ResponsePacket = packet.(*ResponsePacket)
}
