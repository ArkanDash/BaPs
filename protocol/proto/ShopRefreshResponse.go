// Code auto-generated by gucooing & Hiro420 DO NOT EDIT.

package proto

import (
	"github.com/bytedance/sonic"
	"github.com/gucooing/BaPs/protocol/mx"
)

type ShopRefreshResponse struct {
	*ResponsePacket

	ParcelResultDB *ParcelResultDB
	ShopInfoDB     *ShopInfoDB
}

func (x *ShopRefreshResponse) String() string {
	j, _ := sonic.MarshalString(x)
	return j
}

func (x *ShopRefreshResponse) SetPacket(packet mx.Message) {
	if x == nil {
		return
	}
	x.ResponsePacket = packet.(*ResponsePacket)
}
