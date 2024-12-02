package models

type AddGameDTO struct {
	Name        string `json:"name"`
	GameTypeInt int32  `json:"game_type"`
	StartTime   string `json:"start_time"`
	EndTime     string `json:"end_time"`
	CreatorId   int32
}

type AddGameResultDTO struct {
	GameId string `json:"game_id"`
	Result string `json:"result"`
}
