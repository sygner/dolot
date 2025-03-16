package models

import "time"

type Winners struct {
	Id           int32            `db:"id" json:"id"`
	GameId       string           `db:"game_id" json:"game_id"`
	GameType     int32            `db:"game_type" json:"game_type"`
	Divisions    []DivisionResult `db:"divisions" json:"divisions"`
	ResultNumber string           `db:"result_number" json:"result_number"`
	Prize        uint32           `db:"prize" json:"prize"`
	JackPot      bool             `db:"jackpot" json:"jackpot"`
	TotalPaid    *string          `db:"total_paid" json:"total_paid"`
	CreatedAt    time.Time        `db:"created_at" json:"created_at"`
}

type WinnersCount struct {
	Id           int32                 `db:"id" json:"id"`
	GameId       string                `db:"game_id" json:"game_id"`
	GameType     int32                 `db:"game_type" json:"game_type"`
	Divisions    []DivisionResultCount `db:"divisions" json:"divisions"`
	ResultNumber string                `db:"result_number" json:"result_number"`
	Prize        uint32                `db:"prize" json:"prize"`
	JackPot      bool                  `db:"jackpot" json:"jackpot"`
	TotalPaid    *string               `db:"total_paid" json:"total_paid"`
	CreatedAt    time.Time             `db:"created_at" json:"created_at"`
}
