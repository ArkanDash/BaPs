// Code auto-generated by gucooing & Hiro420 DO NOT EDIT.

package proto

import (
	"github.com/bytedance/sonic"
	"github.com/gucooing/BaPs/protocol/mx"
)

type CafeClearPresetRequest struct {
	*RequestPacket

	SlotId int32
}

func (x *CafeClearPresetRequest) String() string {
	j, _ := sonic.MarshalString(x)
	return j
}

func (x *CafeClearPresetRequest) SetPacket(packet mx.Message) {
	if x == nil {
		return
	}
	x.RequestPacket = packet.(*RequestPacket)
}
