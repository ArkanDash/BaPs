// Code auto-generated by gucooing & Hiro420 DO NOT EDIT.

package proto

import (
	"github.com/bytedance/sonic"
	"github.com/gucooing/BaPs/protocol/mx"
)

type KillLogCollection struct {
	*KillLog
}

func (x *KillLogCollection) String() string {
	j, _ := sonic.MarshalString(x)
	return j
}

func (x *KillLogCollection) SetPacket(packet mx.Message) {

}
