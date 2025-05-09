// Code generated by gucooing DO NOT EDIT.

package proto

import (
	"encoding/json"

	"github.com/gucooing/BaPs/pkg/mx"
)

type CafeRelocateFurnitureResponse struct{
	message ProtoMessage
	ResponsePacket

    CafeDB *CafeDB
    RelocatedFurnitureDB *FurnitureDB
}

func (x *CafeRelocateFurnitureResponse) String() string {
	j, _ := json.Marshal(x)
	return string(j)
}

func (x *CafeRelocateFurnitureResponse) ProtoReflect() Message {
	return x
}

func (x *CafeRelocateFurnitureResponse) GetProtocol() int32 {
	return mx.Protocol_Cafe_Relocate
}

func (x *CafeRelocateFurnitureResponse) SetSessionKey(base *BasePacket) {
	if x == nil {
		return
	}
	x.BasePacket = base
}

