// Code auto-generated by gucooing & Hiro420 DO NOT EDIT.

package proto

import (
	"github.com/bytedance/sonic"
	"github.com/gucooing/BaPs/protocol/mx"
)

type ScenarioGroupHistoryUpdateResponse struct {
	*ResponsePacket

	ScenarioGroupHistoryDB *ScenarioGroupHistoryDB
}

func (x *ScenarioGroupHistoryUpdateResponse) String() string {
	j, _ := sonic.MarshalString(x)
	return j
}

func (x *ScenarioGroupHistoryUpdateResponse) SetPacket(packet mx.Message) {
	if x == nil {
		return
	}
	x.ResponsePacket = packet.(*ResponsePacket)
}
