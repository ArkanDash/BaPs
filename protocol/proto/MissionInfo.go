// Code generated by gucooing DO NOT EDIT.

package proto

import (
	"time"
)

type MissionInfo struct {
	Id                          int64
	Category                    MissionCategory
	ResetType                   MissionResetType
	ToastDisplayType            MissionToastDisplayConditionType
	Description                 uint32
	IsVisible                   bool
	IsLimited                   bool
	StartDate                   time.Time
	StartableEndDate            time.Time
	EndDate                     time.Time
	EndDday                     int64
	AccountState                AccountState
	AccountLevel                int64
	PreMissionIds               []int64
	NextMissionId               int64
	SuddenMissionContentTypes   []SuddenMissionContentType
	CompleteConditionType       MissionCompleteConditionType
	CompleteConditionCount      int64
	CompleteConditionParameters []int64
	RewardIcon                  string
	Rewards                     []*ParcelInfo
	DateAutoRefer               ContentType
	ToastImagePath              string
	DisplayOrder                int64
	HasFollowingMission         bool
	Shortcuts                   []string
	ChallengeStageId            int64
}
