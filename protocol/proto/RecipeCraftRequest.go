// Code auto-generated by gucooing & Hiro420 DO NOT EDIT.

package proto

import (
	"github.com/bytedance/sonic"
	"github.com/gucooing/BaPs/protocol/mx"
)

type RecipeCraftRequest struct {
	*RequestPacket

	RecipeCraftUniqueId      int64
	RecipeIngredientUniqueId int64
}

func (x *RecipeCraftRequest) String() string {
	j, _ := sonic.MarshalString(x)
	return j
}

func (x *RecipeCraftRequest) SetPacket(packet mx.Message) {
	if x == nil {
		return
	}
	x.RequestPacket = packet.(*RequestPacket)
}
