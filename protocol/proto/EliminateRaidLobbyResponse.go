// Code generated by gucooing DO NOT EDIT.

package proto

import (
	"encoding/json"

	"github.com/gucooing/BaPs/pkg/mx"
)

type EliminateRaidLobbyResponse struct{
	message ProtoMessage
	ResponsePacket

    SeasonType RaidSeasonType
    RaidGiveUpDB *RaidGiveUpDB
    RaidLobbyInfoDB *EliminateRaidLobbyInfoDB
    ParcelResultDB *ParcelResultDB
}

func (x *EliminateRaidLobbyResponse) String() string {
	j, _ := json.Marshal(x)
	return string(j)
}

func (x *EliminateRaidLobbyResponse) ProtoReflect() Message {
	return x
}

func (x *EliminateRaidLobbyResponse) GetProtocol() int32 {
	return mx.Protocol_EliminateRaid_Lobby
}

func (x *EliminateRaidLobbyResponse) SetSessionKey(base *BasePacket) {
	if x == nil {
		return
	}
	x.BasePacket = base
}

