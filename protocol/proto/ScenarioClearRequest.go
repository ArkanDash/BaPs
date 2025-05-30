// Code auto-generated by gucooing & Hiro420 DO NOT EDIT.

package proto

import (
	"github.com/bytedance/sonic"
	"github.com/gucooing/BaPs/protocol/mx"
)

type ScenarioClearRequest struct {
	*RequestPacket

	ScenarioId    int64
	BattleSummary *BattleSummary
}

func (x *ScenarioClearRequest) String() string {
	j, _ := sonic.MarshalString(x)
	return j
}

func (x *ScenarioClearRequest) SetPacket(packet mx.Message) {
	if x == nil {
		return
	}
	x.RequestPacket = packet.(*RequestPacket)
}
