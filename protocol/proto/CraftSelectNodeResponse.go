// Code auto-generated by gucooing & Hiro420 DO NOT EDIT.

package proto

import (
	"github.com/bytedance/sonic"
	"github.com/gucooing/BaPs/protocol/mx"
)

type CraftSelectNodeResponse struct {
	*ResponsePacket

	SelectedNodeDB *CraftNodeDB
}

func (x *CraftSelectNodeResponse) String() string {
	j, _ := sonic.MarshalString(x)
	return j
}

func (x *CraftSelectNodeResponse) SetPacket(packet mx.Message) {
	if x == nil {
		return
	}
	x.ResponsePacket = packet.(*ResponsePacket)
}
