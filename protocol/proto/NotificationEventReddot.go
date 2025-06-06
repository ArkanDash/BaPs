// Code auto-generated by gucooing & Hiro420 DO NOT EDIT.
package proto

type NotificationEventReddot int32

const (
	NotificationEventReddot_StagePointReward        NotificationEventReddot = 0
	NotificationEventReddot_MissionComplete         NotificationEventReddot = 1
	NotificationEventReddot_MiniGameMissionComplete NotificationEventReddot = 2
	NotificationEventReddot_WorldRaidReward         NotificationEventReddot = 3
	NotificationEventReddot_ConquestCalculateReward NotificationEventReddot = 4
	NotificationEventReddot_DiceRaceLapReward       NotificationEventReddot = 5
)

var (
	NotificationEventReddot_name = map[int32]string{
		0: "StagePointReward",
		1: "MissionComplete",
		2: "MiniGameMissionComplete",
		3: "WorldRaidReward",
		4: "ConquestCalculateReward",
		5: "DiceRaceLapReward",
	}
	NotificationEventReddot_value = map[string]int32{
		"StagePointReward":        0,
		"MissionComplete":         1,
		"MiniGameMissionComplete": 2,
		"WorldRaidReward":         3,
		"ConquestCalculateReward": 4,
		"DiceRaceLapReward":       5,
	}
)

func (x NotificationEventReddot) String() string {
	return NotificationEventReddot_name[int32(x)]
}

func (x NotificationEventReddot) Value(sr string) NotificationEventReddot {
	return NotificationEventReddot(NotificationEventReddot_value[sr])
}
