package models

type AddTicketDTO struct {
	UserId     int32
	TicketType string `json:"ticket_type"`
}

type AddTicketsDTO struct {
	AddTickets []AddTicketDTO `json:"tickets"`
}

type UseTicketDTO struct {
	GameId string `json:"game_id"`
	Amount int32  `json:"amount"`
}

type BuyTicketDTO struct {
	WalletId     int32 `json:"wallet_id"`
	TotalTickets int32 `json:"total_tickets"`
}
