package models

type AddUserChoiceDTO struct {
	UserId             int32
	GameId             string    `json:"game_id"`
	ChosenMainNumbers  [][]int32 `json:"chosen_main_numbers"`
	ChosenBonusNumbers [][]int32 `json:"chosen_bonus_numbers"`
	BoughtPrice        int32
	ShouldReturn       bool `json:"should_return"`
}

type AddUserChoicesDTO struct {
	UserChoices []AddUserChoiceDTO `json:"user_choices"`
}

type GetUserChoicesByUserIdAndTimeRangeDTO struct {
	UserId    int32
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}
