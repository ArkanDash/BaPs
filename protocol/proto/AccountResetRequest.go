// Code auto-generated by gucooing & Hiro420 DO NOT EDIT.

package proto

import (
	"github.com/bytedance/sonic"
	"github.com/gucooing/BaPs/protocol/mx"
)

type AccountResetRequest struct {
	*RequestPacket

	DevId string
}

func (x *AccountResetRequest) String() string {
	j, _ := sonic.MarshalString(x)
	return j
}

func (x *AccountResetRequest) SetPacket(packet mx.Message) {
	if x == nil {
		return
	}
	x.RequestPacket = packet.(*RequestPacket)
}
