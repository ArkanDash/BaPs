// Code auto-generated by gucooing & Hiro420 DO NOT EDIT.

package proto

import (
	"github.com/bytedance/sonic"
	"github.com/gucooing/BaPs/protocol/mx"
)

type ContentsValueChangeDB struct {
	ContentsChangeType ContentsChangeType
}

func (x *ContentsValueChangeDB) String() string {
	j, _ := sonic.MarshalString(x)
	return j
}

func (x *ContentsValueChangeDB) SetPacket(packet mx.Message) {

}
