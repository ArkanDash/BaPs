// Code auto-generated by gucooing & Hiro420 DO NOT EDIT.

package proto

import (
	"github.com/bytedance/sonic"
	"github.com/gucooing/BaPs/protocol/mx"
)

type MiniGameDreamMakerResetRequest struct {
	*RequestPacket

	EventContentId int64
}

func (x *MiniGameDreamMakerResetRequest) String() string {
	j, _ := sonic.MarshalString(x)
	return j
}

func (x *MiniGameDreamMakerResetRequest) SetPacket(packet mx.Message) {
	if x == nil {
		return
	}
	x.RequestPacket = packet.(*RequestPacket)
}
