// Code auto-generated by gucooing & Hiro420 DO NOT EDIT.
package proto

type LayoutKind int32

const (
	LayoutKind_Sequential LayoutKind = 0
	LayoutKind_Explicit   LayoutKind = 2
	LayoutKind_Auto       LayoutKind = 3
)

var (
	LayoutKind_name = map[int32]string{
		0: "Sequential",
		2: "Explicit",
		3: "Auto",
	}
	LayoutKind_value = map[string]int32{
		"Sequential": 0,
		"Explicit":   2,
		"Auto":       3,
	}
)

func (x LayoutKind) String() string {
	return LayoutKind_name[int32(x)]
}

func (x LayoutKind) Value(sr string) LayoutKind {
	return LayoutKind(LayoutKind_value[sr])
}
