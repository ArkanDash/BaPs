// Code auto-generated by gucooing & Hiro420 DO NOT EDIT.

package proto

import (
	"github.com/bytedance/sonic"
	"github.com/gucooing/BaPs/protocol/mx"
)

type GuideMissionSeasonListResponse struct {
	*ResponsePacket

	GuideMissionSeasonDBs []*GuideMissionSeasonDB
}

func (x *GuideMissionSeasonListResponse) String() string {
	j, _ := sonic.MarshalString(x)
	return j
}

func (x *GuideMissionSeasonListResponse) SetPacket(packet mx.Message) {
	if x == nil {
		return
	}
	x.ResponsePacket = packet.(*ResponsePacket)
}
