// Code generated by gucooing DO NOT EDIT.

package proto

type HeroSummaryDetailFlag int32

const (
	HeroSummaryDetailFlag_None             = 0
	HeroSummaryDetailFlag_BattleProperty   = 2
	HeroSummaryDetailFlag_BattleStatistics = 4
	HeroSummaryDetailFlag_NumericLogs      = 8
	HeroSummaryDetailFlag_StatSnapshot     = 16
	HeroSummaryDetailFlag_Default          = 14
	HeroSummaryDetailFlag_All              = 30
)

var (
	HeroSummaryDetailFlag_name = map[int32]string{
		0:  "None",
		2:  "BattleProperty",
		4:  "BattleStatistics",
		8:  "NumericLogs",
		16: "StatSnapshot",
		14: "Default",
		30: "All",
	}
	HeroSummaryDetailFlag_value = map[string]int32{
		"None":             0,
		"BattleProperty":   2,
		"BattleStatistics": 4,
		"NumericLogs":      8,
		"StatSnapshot":     16,
		"Default":          14,
		"All":              30,
	}
)

func (x HeroSummaryDetailFlag) String() string {
	return HeroSummaryDetailFlag_name[int32(x)]
}
