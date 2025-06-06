// Code auto-generated by gucooing & Hiro420 DO NOT EDIT.

package proto

import (
	"github.com/bytedance/sonic"
	"github.com/gucooing/BaPs/protocol/mx"
)

type EventContentMapMoveResponse struct {
	*ResponsePacket

	SaveDataDB                *EventContentMainStageSaveDB
	EchelonEntityId           int64
	StrategyObject            *Strategy
	StrategyObjectParcelInfos []*ParcelInfo
}

func (x *EventContentMapMoveResponse) String() string {
	j, _ := sonic.MarshalString(x)
	return j
}

func (x *EventContentMapMoveResponse) SetPacket(packet mx.Message) {
	if x == nil {
		return
	}
	x.ResponsePacket = packet.(*ResponsePacket)
}
