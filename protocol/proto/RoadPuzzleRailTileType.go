// Code auto-generated by gucooing & Hiro420 DO NOT EDIT.
package proto

type RoadPuzzleRailTileType int32

const (
	RoadPuzzleRailTileType_None       RoadPuzzleRailTileType = 0
	RoadPuzzleRailTileType_Straight   RoadPuzzleRailTileType = 1
	RoadPuzzleRailTileType_CurveBig   RoadPuzzleRailTileType = 2
	RoadPuzzleRailTileType_CurveSmall RoadPuzzleRailTileType = 3
)

var (
	RoadPuzzleRailTileType_name = map[int32]string{
		0: "None",
		1: "Straight",
		2: "CurveBig",
		3: "CurveSmall",
	}
	RoadPuzzleRailTileType_value = map[string]int32{
		"None":       0,
		"Straight":   1,
		"CurveBig":   2,
		"CurveSmall": 3,
	}
)

func (x RoadPuzzleRailTileType) String() string {
	return RoadPuzzleRailTileType_name[int32(x)]
}

func (x RoadPuzzleRailTileType) Value(sr string) RoadPuzzleRailTileType {
	return RoadPuzzleRailTileType(RoadPuzzleRailTileType_value[sr])
}
