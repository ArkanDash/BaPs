// Code generated by gucooing DO NOT EDIT.

package proto

type ConsumeResultDB struct {
	RemovedItemServerIds                    []int64
	RemovedEquipmentServerIds               []int64
	RemovedFurnitureServerIds               []int64
	UsedItemServerIdAndRemainingCounts      map[int64]int64
	UsedEquipmentServerIdAndRemainingCounts map[int64]int64
	UsedFurnitureServerIdAndRemainingCounts map[int64]int64
}
