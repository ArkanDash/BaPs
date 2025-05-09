// Code generated by gucooing DO NOT EDIT.

package proto

type CampaignState int32

const (
	CampaignState_BeforeStart      = 0
	CampaignState_BeginPlayerPhase = 1
	CampaignState_PlayerPhase      = 2
	CampaignState_EndPlayerPhase   = 3
	CampaignState_BeginEnemyPhase  = 4
	CampaignState_EnemyPhase       = 5
	CampaignState_EndEnemyPhase    = 6
	CampaignState_Win              = 7
	CampaignState_Lose             = 8
	CampaignState_StrategySkip     = 9
)

var (
	CampaignState_name = map[int32]string{
		0: "BeforeStart",
		1: "BeginPlayerPhase",
		2: "PlayerPhase",
		3: "EndPlayerPhase",
		4: "BeginEnemyPhase",
		5: "EnemyPhase",
		6: "EndEnemyPhase",
		7: "Win",
		8: "Lose",
		9: "StrategySkip",
	}
	CampaignState_value = map[string]int32{
		"BeforeStart":      0,
		"BeginPlayerPhase": 1,
		"PlayerPhase":      2,
		"EndPlayerPhase":   3,
		"BeginEnemyPhase":  4,
		"EnemyPhase":       5,
		"EndEnemyPhase":    6,
		"Win":              7,
		"Lose":             8,
		"StrategySkip":     9,
	}
)

func (x CampaignState) String() string {
	return CampaignState_name[int32(x)]
}
