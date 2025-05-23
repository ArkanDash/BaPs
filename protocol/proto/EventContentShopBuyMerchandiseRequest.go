// Code auto-generated by gucooing & Hiro420 DO NOT EDIT.

package proto

import (
	"github.com/bytedance/sonic"
	"github.com/gucooing/BaPs/protocol/mx"
)

type EventContentShopBuyMerchandiseRequest struct {
	*RequestPacket

	EventContentId       int64
	IsRefreshMerchandise bool
	ShopUniqueId         int64
	GoodsUniqueId        int64
	PurchaseCount        int64
}

func (x *EventContentShopBuyMerchandiseRequest) String() string {
	j, _ := sonic.MarshalString(x)
	return j
}

func (x *EventContentShopBuyMerchandiseRequest) SetPacket(packet mx.Message) {
	if x == nil {
		return
	}
	x.RequestPacket = packet.(*RequestPacket)
}
