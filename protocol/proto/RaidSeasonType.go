// Code generated by gucooing DO NOT EDIT.

package proto

type RaidSeasonType int32

const (
	RaidSeasonType_None       = 0
	RaidSeasonType_Open       = 1
	RaidSeasonType_Close      = 2
	RaidSeasonType_Settlement = 3
)

var (
	RaidSeasonType_name = map[int32]string{
		0: "None",
		1: "Open",
		2: "Close",
		3: "Settlement",
	}
	RaidSeasonType_value = map[string]int32{
		"None":       0,
		"Open":       1,
		"Close":      2,
		"Settlement": 3,
	}
)

func (x RaidSeasonType) String() string {
	return RaidSeasonType_name[int32(x)]
}
