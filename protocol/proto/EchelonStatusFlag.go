// Code generated by gucooing DO NOT EDIT.

package proto

type EchelonStatusFlag int32

const (
	EchelonStatusFlag_None         = 0
	EchelonStatusFlag_BeforeDeploy = 1
	EchelonStatusFlag_OnDuty       = 2
)

var (
	EchelonStatusFlag_name = map[int32]string{
		0: "None",
		1: "BeforeDeploy",
		2: "OnDuty",
	}
	EchelonStatusFlag_value = map[string]int32{
		"None":         0,
		"BeforeDeploy": 1,
		"OnDuty":       2,
	}
)

func (x EchelonStatusFlag) String() string {
	return EchelonStatusFlag_name[int32(x)]
}
