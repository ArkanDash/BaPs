// Code generated by gucooing DO NOT EDIT.

package proto

import (
    "github.com/gucooing/BaPs/pkg/mx"
)

type CafeDB struct{
    CafeDBId int64
    CafeId int64
    AccountId int64
    CafeRank int32
    LastUpdate mx.MxTime
    LastSummonDate mx.MxTime
    IsNew bool
    CafeVisitCharacterDBs map[int64]*CafeCharacterDB
    FurnitureDBs []*FurnitureDB
    ProductionAppliedTime mx.MxTime
    ProductionDB *CafeProductionDB
}