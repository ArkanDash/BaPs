// Code auto-generated by gucooing & Hiro420 DO NOT EDIT.

package proto

import (
	"github.com/bytedance/sonic"
)

type PurchaseOrderDB struct {
	ShopCashId      int64
	StatusCode      PurchaseStatusCode
	PurchaseOrderId int64
}

func (x *PurchaseOrderDB) String() string {
	j, _ := sonic.MarshalString(x)
	return j
}
