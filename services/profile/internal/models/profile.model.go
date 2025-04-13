package models

import "time"

type Profile struct {
	UserId        int32     `json:"user_id"`
	Sid           string    `json:"sid"`
	Username      string    `json:"username"`
	Score         float32   `json:"score"`
	Impression    int32     `json:"impression"`
	DCoin         int32     `json:"d_coin"`
	Rank          int32     `json:"rank"`
	GamesQuantity int32     `json:"games_quantity"`
	WonGames      int32     `json:"won_games"`
	LostGames     int32     `json:"lost_games"`
	CreatedAt     time.Time `json:"created_at"`
	HighestRank   int32     `json:"highest_rank"`
}

type Ranking struct {
	TotalRanking           uint32 `json:"total_ranking"`
	IndividualRanking      uint32 `json:"individual_ranking"`
	SeasonRanking          uint32 `json:"season_ranking"`
	MonthRanking           uint32 `json:"month_ranking"`
	SeasonRankChangesCount uint32 `json:"season_rank_changes_count"`
	MonthRankChangesCount  uint32 `json:"month_rank_changes_count"`
	AllRankChangesCount    uint32 `json:"all_rank_changes_count"`
}
