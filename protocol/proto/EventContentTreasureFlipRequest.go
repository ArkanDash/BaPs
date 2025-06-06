// Code auto-generated by gucooing & Hiro420 DO NOT EDIT.

package proto

import (
	"github.com/bytedance/sonic"
	"github.com/gucooing/BaPs/protocol/mx"
)

type EventContentTreasureFlipRequest struct {
	*RequestPacket

	EventContentId int64
	Round          int32
	Cells          []*EventContentTreasureCell
}

func (x *EventContentTreasureFlipRequest) String() string {
	j, _ := sonic.MarshalString(x)
	return j
}

func (x *EventContentTreasureFlipRequest) SetPacket(packet mx.Message) {
	if x == nil {
		return
	}
	x.RequestPacket = packet.(*RequestPacket)
}
