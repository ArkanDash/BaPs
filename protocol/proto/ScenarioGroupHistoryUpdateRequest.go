// Code auto-generated by gucooing & Hiro420 DO NOT EDIT.

package proto

import (
	"github.com/bytedance/sonic"
	"github.com/gucooing/BaPs/protocol/mx"
)

type ScenarioGroupHistoryUpdateRequest struct {
	*RequestPacket

	ScenarioGroupUniqueId int64
	ScenarioType          int64
}

func (x *ScenarioGroupHistoryUpdateRequest) String() string {
	j, _ := sonic.MarshalString(x)
	return j
}

func (x *ScenarioGroupHistoryUpdateRequest) SetPacket(packet mx.Message) {
	if x == nil {
		return
	}
	x.RequestPacket = packet.(*RequestPacket)
}
