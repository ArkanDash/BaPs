// Code auto-generated by gucooing & Hiro420 DO NOT EDIT.

package proto

import (
	"github.com/bytedance/sonic"
	"github.com/gucooing/BaPs/protocol/mx"
)

type EquipmentSetting struct {
	*ValueType

	IsValid  bool
	ServerId int64
	UniqueId int64
	Level    int32
	Tier     int32
}

func (x *EquipmentSetting) String() string {
	j, _ := sonic.MarshalString(x)
	return j
}

func (x *EquipmentSetting) SetPacket(packet mx.Message) {
	if x == nil {
		return
	}
	x.ValueType = packet.(*ValueType)
}
