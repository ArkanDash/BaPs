// Code auto-generated by gucooing & Hiro420 DO NOT EDIT.

package proto

import (
	"github.com/bytedance/sonic"
	"github.com/gucooing/BaPs/protocol/mx"
)

type MiniGameTableBoardEncounterInputRequest struct {
	*RequestPacket

	EventContentId int64
	ObjectServerId int64
	EncounterStage int32
	SelectedIndex  int32
}

func (x *MiniGameTableBoardEncounterInputRequest) String() string {
	j, _ := sonic.MarshalString(x)
	return j
}

func (x *MiniGameTableBoardEncounterInputRequest) SetPacket(packet mx.Message) {
	if x == nil {
		return
	}
	x.RequestPacket = packet.(*RequestPacket)
}
