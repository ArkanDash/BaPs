// Code auto-generated by gucooing & Hiro420 DO NOT EDIT.

package proto

import (
	"github.com/bytedance/sonic"
	"github.com/gucooing/BaPs/protocol/mx"
)

type CafeTravelRequest struct {
	*RequestPacket

	TargetAccountId          int64
	CurrentVisitingAccountId int64
}

func (x *CafeTravelRequest) String() string {
	j, _ := sonic.MarshalString(x)
	return j
}

func (x *CafeTravelRequest) SetPacket(packet mx.Message) {
	if x == nil {
		return
	}
	x.RequestPacket = packet.(*RequestPacket)
}
