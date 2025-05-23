// Code auto-generated by gucooing & Hiro420 DO NOT EDIT.

package proto

import (
	"github.com/bytedance/sonic"
	"github.com/gucooing/BaPs/protocol/mx"
)

type ClanPermitResponse struct {
	*ResponsePacket

	ClanDB       *ClanDB
	ClanMemberDB *ClanMemberDB
}

func (x *ClanPermitResponse) String() string {
	j, _ := sonic.MarshalString(x)
	return j
}

func (x *ClanPermitResponse) SetPacket(packet mx.Message) {
	if x == nil {
		return
	}
	x.ResponsePacket = packet.(*ResponsePacket)
}
