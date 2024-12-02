package models

type AddUserChoiceDTO struct {
	Id                 string    `json:"id"`
	UserId             int32     `json:"user_id"`
	GameId             string    `json:"game_id"`
	ChosenMainNumbers  [][]int32 `json:"chosen_main_numbers"`
	ChosenBonusNumbers [][]int32 `json:"chosen_bonus_numbers"`
}
