// Code auto-generated by gucooing & Hiro420 DO NOT EDIT.

package proto

import (
	"github.com/bytedance/sonic"
)

type MiniGameDreamMakerParameterDB struct {
	ParameterType DreamMakerParameterType
	BaseAmount    int64
	CurrentAmount int64
}

func (x *MiniGameDreamMakerParameterDB) String() string {
	j, _ := sonic.MarshalString(x)
	return j
}
