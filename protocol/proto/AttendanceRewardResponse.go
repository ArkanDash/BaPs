// Code auto-generated by gucooing & Hiro420 DO NOT EDIT.

package proto

import (
	"github.com/bytedance/sonic"
	"github.com/gucooing/BaPs/protocol/mx"
)

type AttendanceRewardResponse struct {
	*ResponsePacket

	AttendanceBookRewards []*AttendanceBookReward
	AttendanceHistoryDBs  []*AttendanceHistoryDB
	ParcelResultDB        *ParcelResultDB
}

func (x *AttendanceRewardResponse) String() string {
	j, _ := sonic.MarshalString(x)
	return j
}

func (x *AttendanceRewardResponse) SetPacket(packet mx.Message) {
	if x == nil {
		return
	}
	x.ResponsePacket = packet.(*ResponsePacket)
}
