// Code auto-generated by gucooing & Hiro420 DO NOT EDIT.

package proto

import (
	"github.com/bytedance/sonic"
)

type BeforehandGachaSnapshotDB struct {
	ShopUniqueId int64
	GoodsId      int64
	LastIndex    int64
	LastResults  []int64
	SavedIndex   int64
	SavedResults []int64
	PickedIndex  int64
}

func (x *BeforehandGachaSnapshotDB) String() string {
	j, _ := sonic.MarshalString(x)
	return j
}
