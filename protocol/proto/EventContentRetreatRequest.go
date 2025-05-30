// Code auto-generated by gucooing & Hiro420 DO NOT EDIT.

package proto

import (
	"github.com/bytedance/sonic"
	"github.com/gucooing/BaPs/protocol/mx"
)

type EventContentRetreatRequest struct {
	*RequestPacket

	EventContentId int64
	StageUniqueId  int64
}

func (x *EventContentRetreatRequest) String() string {
	j, _ := sonic.MarshalString(x)
	return j
}

func (x *EventContentRetreatRequest) SetPacket(packet mx.Message) {
	if x == nil {
		return
	}
	x.RequestPacket = packet.(*RequestPacket)
}
