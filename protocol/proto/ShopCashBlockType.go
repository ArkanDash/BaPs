// Code auto-generated by gucooing & Hiro420 DO NOT EDIT.
package proto

type ShopCashBlockType int64

const (
	ShopCashBlockType_All        ShopCashBlockType = -1
	ShopCashBlockType_AppStore   ShopCashBlockType = -2
	ShopCashBlockType_GooglePlay ShopCashBlockType = -3
	ShopCashBlockType_None       ShopCashBlockType = -9999
)

var (
	ShopCashBlockType_name = map[int64]string{
		-1:    "All",
		-2:    "AppStore",
		-3:    "GooglePlay",
		-9999: "None",
	}
	ShopCashBlockType_value = map[string]int64{
		"All":        -1,
		"AppStore":   -2,
		"GooglePlay": -3,
		"None":       -9999,
	}
)

func (x ShopCashBlockType) String() string {
	return ShopCashBlockType_name[int64(x)]
}

func (x ShopCashBlockType) Value(sr string) ShopCashBlockType {
	return ShopCashBlockType(ShopCashBlockType_value[sr])
}
