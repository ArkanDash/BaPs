// Code generated by gucooing DO NOT EDIT.

package proto

type CharacterDB struct {
	Type                   ParcelType
	ServerId               int64
	UniqueId               int64
	StarGrade              int32
	Level                  int32
	Exp                    int64
	FavorRank              int32
	FavorExp               int64
	PublicSkillLevel       int32
	ExSkillLevel           int32
	PassiveSkillLevel      int32
	ExtraPassiveSkillLevel int32
	LeaderSkillLevel       int32
	IsNew                  bool
	IsLocked               bool
	IsFavorite             bool
	EquipmentServerIds     []int64
	PotentialStats         map[int32]int32
}
