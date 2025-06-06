// Code auto-generated by gucooing & Hiro420 DO NOT EDIT.

package proto

import (
	"github.com/bytedance/sonic"
	"github.com/gucooing/BaPs/protocol/mx"
)

type EventContentTreasureNextRoundResponse struct {
	*ResponsePacket

	BoardHistoryDB *EventContentTreasureHistoryDB
	HiddenImage    *EventContentTreasureCell
}

func (x *EventContentTreasureNextRoundResponse) String() string {
	j, _ := sonic.MarshalString(x)
	return j
}

func (x *EventContentTreasureNextRoundResponse) SetPacket(packet mx.Message) {
	if x == nil {
		return
	}
	x.ResponsePacket = packet.(*ResponsePacket)
}
