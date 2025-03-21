// Code generated by gucooing DO NOT EDIT.

package proto

import (
	"encoding/json"

	"github.com/gucooing/BaPs/pkg/mx"
)

type CharacterGearListResponse struct{
	message ProtoMessage
	ResponsePacket

    GearDBs []*GearDB
}

func (x *CharacterGearListResponse) String() string {
	j, _ := json.Marshal(x)
	return string(j)
}

func (x *CharacterGearListResponse) ProtoReflect() Message {
	return x
}

func (x *CharacterGearListResponse) GetProtocol() int32 {
	return mx.Protocol_CharacterGear_List
}

func (x *CharacterGearListResponse) SetSessionKey(base *BasePacket) {
	if x == nil {
		return
	}
	x.BasePacket = base
}

