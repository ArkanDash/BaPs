// Code generated by gucooing DO NOT EDIT.

package proto

type ShopProductType int32

const (
	ShopProductType_None    = 0
	ShopProductType_General = 1
	ShopProductType_Refresh = 2
)

var (
	ShopProductType_name = map[int32]string{
		0: "None",
		1: "General",
		2: "Refresh",
	}
	ShopProductType_value = map[string]int32{
		"None":    0,
		"General": 1,
		"Refresh": 2,
	}
)

func (x ShopProductType) String() string {
	return ShopProductType_name[int32(x)]
}
