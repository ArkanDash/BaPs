// Code auto-generated by gucooing & Hiro420 DO NOT EDIT.
package proto

type RankingSearchType int32

const (
	RankingSearchType_None  RankingSearchType = 0
	RankingSearchType_Rank  RankingSearchType = 1
	RankingSearchType_Score RankingSearchType = 2
)

var (
	RankingSearchType_name = map[int32]string{
		0: "None",
		1: "Rank",
		2: "Score",
	}
	RankingSearchType_value = map[string]int32{
		"None":  0,
		"Rank":  1,
		"Score": 2,
	}
)

func (x RankingSearchType) String() string {
	return RankingSearchType_name[int32(x)]
}

func (x RankingSearchType) Value(sr string) RankingSearchType {
	return RankingSearchType(RankingSearchType_value[sr])
}
