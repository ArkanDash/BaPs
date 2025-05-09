// Code generated by gucooing DO NOT EDIT.

package proto

import (
	"encoding/json"

	"github.com/gucooing/BaPs/pkg/mx"
)

type MultiFloorRaidEndBattleResponse struct{
	message ProtoMessage
	ResponsePacket

    MultiFloorRaidDB *MultiFloorRaidDB
    ParcelResultDB *ParcelResultDB
}

func (x *MultiFloorRaidEndBattleResponse) String() string {
	j, _ := json.Marshal(x)
	return string(j)
}

func (x *MultiFloorRaidEndBattleResponse) ProtoReflect() Message {
	return x
}

func (x *MultiFloorRaidEndBattleResponse) GetProtocol() int32 {
	return mx.Protocol_MultiFloorRaid_EndBattle
}

func (x *MultiFloorRaidEndBattleResponse) SetSessionKey(base *BasePacket) {
	if x == nil {
		return
	}
	x.BasePacket = base
}

