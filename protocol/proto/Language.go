// Code generated by gucooing DO NOT EDIT.

package proto

type Language int32

const (
	Language_Kr = 0
	Language_Jp = 1
	Language_Th = 2
	Language_Tw = 3
	Language_En = 4
)

var (
	Language_name = map[int32]string{
		0: "Kr",
		1: "Jp",
		2: "Th",
		3: "Tw",
		4: "En",
	}
	Language_value = map[string]int32{
		"Kr": 0,
		"Jp": 1,
		"Th": 2,
		"Tw": 3,
		"En": 4,
	}
)

func (x Language) String() string {
	return Language_name[int32(x)]
}
