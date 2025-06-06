// Code auto-generated by gucooing & Hiro420 DO NOT EDIT.
package proto

type ConquestDisplayType int32

const (
	ConquestDisplayType_None                     ConquestDisplayType = 0
	ConquestDisplayType_TileConquered            ConquestDisplayType = 1
	ConquestDisplayType_TileUpgraded             ConquestDisplayType = 2
	ConquestDisplayType_UnexpectedEvent          ConquestDisplayType = 3
	ConquestDisplayType_BossOpen                 ConquestDisplayType = 4
	ConquestDisplayType_PropAnimation            ConquestDisplayType = 5
	ConquestDisplayType_PropAnimationAndBlock    ConquestDisplayType = 6
	ConquestDisplayType_PropAnimationHoldAndPlay ConquestDisplayType = 7
	ConquestDisplayType_Operator                 ConquestDisplayType = 8
	ConquestDisplayType_StepComplete             ConquestDisplayType = 9
	ConquestDisplayType_MassErosion              ConquestDisplayType = 10
	ConquestDisplayType_Erosion                  ConquestDisplayType = 11
	ConquestDisplayType_ErosionRemove            ConquestDisplayType = 12
	ConquestDisplayType_CheckTileErosion         ConquestDisplayType = 13
	ConquestDisplayType_StepOpen                 ConquestDisplayType = 14
	ConquestDisplayType_BossClear                ConquestDisplayType = 15
	ConquestDisplayType_HideConquestUI           ConquestDisplayType = 16
	ConquestDisplayType_ShowConquestUI           ConquestDisplayType = 17
	ConquestDisplayType_HideHexaUI               ConquestDisplayType = 18
	ConquestDisplayType_ShowHexaUI               ConquestDisplayType = 19
	ConquestDisplayType_StepObjectComplete       ConquestDisplayType = 20
	ConquestDisplayType_CameraSetting            ConquestDisplayType = 21
	ConquestDisplayType_PlayMapEnterScenario     ConquestDisplayType = 22
	ConquestDisplayType_ShowTileConquerReward    ConquestDisplayType = 23
)

var (
	ConquestDisplayType_name = map[int32]string{
		0:  "None",
		1:  "TileConquered",
		2:  "TileUpgraded",
		3:  "UnexpectedEvent",
		4:  "BossOpen",
		5:  "PropAnimation",
		6:  "PropAnimationAndBlock",
		7:  "PropAnimationHoldAndPlay",
		8:  "Operator",
		9:  "StepComplete",
		10: "MassErosion",
		11: "Erosion",
		12: "ErosionRemove",
		13: "CheckTileErosion",
		14: "StepOpen",
		15: "BossClear",
		16: "HideConquestUI",
		17: "ShowConquestUI",
		18: "HideHexaUI",
		19: "ShowHexaUI",
		20: "StepObjectComplete",
		21: "CameraSetting",
		22: "PlayMapEnterScenario",
		23: "ShowTileConquerReward",
	}
	ConquestDisplayType_value = map[string]int32{
		"None":                     0,
		"TileConquered":            1,
		"TileUpgraded":             2,
		"UnexpectedEvent":          3,
		"BossOpen":                 4,
		"PropAnimation":            5,
		"PropAnimationAndBlock":    6,
		"PropAnimationHoldAndPlay": 7,
		"Operator":                 8,
		"StepComplete":             9,
		"MassErosion":              10,
		"Erosion":                  11,
		"ErosionRemove":            12,
		"CheckTileErosion":         13,
		"StepOpen":                 14,
		"BossClear":                15,
		"HideConquestUI":           16,
		"ShowConquestUI":           17,
		"HideHexaUI":               18,
		"ShowHexaUI":               19,
		"StepObjectComplete":       20,
		"CameraSetting":            21,
		"PlayMapEnterScenario":     22,
		"ShowTileConquerReward":    23,
	}
)

func (x ConquestDisplayType) String() string {
	return ConquestDisplayType_name[int32(x)]
}

func (x ConquestDisplayType) Value(sr string) ConquestDisplayType {
	return ConquestDisplayType(ConquestDisplayType_value[sr])
}
