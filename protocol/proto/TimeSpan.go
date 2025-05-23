// Code auto-generated by gucooing & Hiro420 DO NOT EDIT.

package proto

import (
	"github.com/bytedance/sonic"
	"github.com/gucooing/BaPs/protocol/mx"
)

type TimeSpan struct {
	*ValueType

	Ticks             int64
	Days              int32
	Hours             int32
	Milliseconds      int32
	Microseconds      int32
	Nanoseconds       int32
	Minutes           int32
	Seconds           int32
	TotalDays         float64
	TotalHours        float64
	TotalMilliseconds float64
	TotalMicroseconds float64
	TotalNanoseconds  float64
	TotalMinutes      float64
	TotalSeconds      float64
}

func (x *TimeSpan) String() string {
	j, _ := sonic.MarshalString(x)
	return j
}

func (x *TimeSpan) SetPacket(packet mx.Message) {
	if x == nil {
		return
	}
	x.ValueType = packet.(*ValueType)
}
