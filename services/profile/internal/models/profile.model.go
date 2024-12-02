package models

import "time"

type Profile struct {
	UserId        int32     `json:"user_id"`
	Sid           string    `json:"sid"`
	Username      string    `json:"username"`
	Score         float32   `json:"score"`
	Impression    int32     `json:"impression"`
	Rank          int32     `json:"rank"`
	GamesQuantity int32     `json:"games_quantity"`
	WonGames      int32     `json:"won_games"`
	LostGames     int32     `json:"lost_games"`
	CreatedAt     time.Time `json:"created_at"`
}
