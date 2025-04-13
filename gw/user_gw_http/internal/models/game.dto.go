package models

type AddGameDTO struct {
	Name        string  `json:"name"`
	GameTypeInt int32   `json:"game_type"`
	StartTime   string  `json:"start_time"`
	EndTime     string  `json:"end_time"`
	Prize       *uint32 `json:"prize,omitempty"`
	AutoCompute bool    `json:"auto_compute"`
	CreatorId   int32
}

type AddGameResultDTO struct {
	GameId string `json:"game_id"`
	Result string `json:"result"`
}

type ChangeGameDetailCalculationDTO struct {
	DayName     *string `json:"day_name,omitempty"`
	PrizeReward *int32  `json:"prize_reward,omitempty"`
	TokenBurn   *int32  `json:"token_burn,omitempty"`
	AutoCompute bool    `json:"auto_compute"`
}

type UpdateGamePrizeDTO struct {
	Prize       *uint32 `json:"prize,omitempty"`
	GameId      string  `json:"game_id"`
	AutoCompute bool    `json:"auto_compute"`
}

type GetAllUserPreviousGamesByGameTypeDTO struct {
	GameType   int32      `json:"game_type"`
	Pagination Pagination `json:"pagination"`
}

type GetUserGamesByTimesAndGameTypesDTO struct {
	GameType  *int32 `json:"game_type,omitempty"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}
