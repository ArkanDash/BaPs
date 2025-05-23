// Code auto-generated by gucooing & Hiro420 DO NOT EDIT.

package proto

import (
	"github.com/bytedance/sonic"
	"github.com/gucooing/BaPs/protocol/mx"
)

type EventContentEnterSubStageRequest struct {
	*RequestPacket

	EventContentId              int64
	StageUniqueId               int64
	LastEnterStageEchelonNumber int64
}

func (x *EventContentEnterSubStageRequest) String() string {
	j, _ := sonic.MarshalString(x)
	return j
}

func (x *EventContentEnterSubStageRequest) SetPacket(packet mx.Message) {
	if x == nil {
		return
	}
	x.RequestPacket = packet.(*RequestPacket)
}
