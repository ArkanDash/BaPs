// Code auto-generated by gucooing & Hiro420 DO NOT EDIT.

package proto

import (
	"github.com/bytedance/sonic"
	"github.com/gucooing/BaPs/protocol/mx"
)

type EventContentShopListRequest struct {
	*RequestPacket

	EventContentId int64
	CategoryList   []ShopCategoryType
}

func (x *EventContentShopListRequest) String() string {
	j, _ := sonic.MarshalString(x)
	return j
}

func (x *EventContentShopListRequest) SetPacket(packet mx.Message) {
	if x == nil {
		return
	}
	x.RequestPacket = packet.(*RequestPacket)
}
