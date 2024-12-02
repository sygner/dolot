package models

import "time"

type AddGameDTO struct {
	Id               string `db:"id" json:"id"`
	Name             string `db:"name" json:"name"`
	GameTypeInt      int32
	GameTypeString   string    `db:"game_type" json:"game_type"`
	NumMainNumbers   int32     `db:"num_main_numbers" json:"num_main_numbers"`
	NumBonusNumbers  *int32    `db:"num_bonus_numbers" json:"num_bonus_numbers"`
	MainNumberRange  int32     `db:"main_number_range" json:"main_number_range"`
	BonusNumberRange *int32    `db:"bonus_number_range" json:"bonus_number_range"`
	StartTime        time.Time `db:"start_time" json:"start_time"`
	EndTime          time.Time `db:"end_time" json:"end_time"`
	CreatorId        int32     `db:"creator_id" json:"creator_id"`
	Result           *string   `db:"result" json:"result"`
}
