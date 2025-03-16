package models

type UpdateTotalPaidDTO struct {
	GameId    string `json:"game_id"`
	TotalPaid uint64 `json:"total_paid"`
}
