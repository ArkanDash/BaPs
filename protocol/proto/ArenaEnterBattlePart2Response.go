// Code generated by gucooing DO NOT EDIT.

package proto

import (
	"encoding/json"

	"github.com/gucooing/BaPs/pkg/mx"
)

type ArenaEnterBattlePart2Response struct{
	message ProtoMessage
	ResponsePacket

    ArenaBattleDB *ArenaBattleDB
    ArenaPlayerInfoDB *ArenaPlayerInfoDB
    AccountCurrencyDB *AccountCurrencyDB
    VictoryRewards *ParcelResultDB
    SeasonRewards *ParcelResultDB
    AllTimeRewards *ParcelResultDB
}

func (x *ArenaEnterBattlePart2Response) String() string {
	j, _ := json.Marshal(x)
	return string(j)
}

func (x *ArenaEnterBattlePart2Response) ProtoReflect() Message {
	return x
}

func (x *ArenaEnterBattlePart2Response) GetProtocol() int32 {
	return mx.Protocol_Arena_EnterBattlePart2
}

func (x *ArenaEnterBattlePart2Response) SetSessionKey(base *BasePacket) {
	if x == nil {
		return
	}
	x.BasePacket = base
}

