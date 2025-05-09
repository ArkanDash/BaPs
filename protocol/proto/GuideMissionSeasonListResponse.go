// Code generated by gucooing DO NOT EDIT.

package proto

import (
	"encoding/json"

	"github.com/gucooing/BaPs/pkg/mx"
)

type GuideMissionSeasonListResponse struct{
	message ProtoMessage
	ResponsePacket

    GuideMissionSeasonDBs []*GuideMissionSeasonDB
}

func (x *GuideMissionSeasonListResponse) String() string {
	j, _ := json.Marshal(x)
	return string(j)
}

func (x *GuideMissionSeasonListResponse) ProtoReflect() Message {
	return x
}

func (x *GuideMissionSeasonListResponse) GetProtocol() int32 {
	return mx.Protocol_Mission_GuideMissionSeasonList
}


func (x *GuideMissionSeasonListResponse) SetSessionKey(base *BasePacket) {
	if x == nil {
		return
	}
	x.BasePacket = base
}

