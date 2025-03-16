package models

import "time"

type Ticket struct {
	ID         int32      `json:"id"`
	Signature  string     `json:"signature"`
	UserId     int32      `json:"user_id"`
	TicketType string     `json:"ticket_type"`
	Status     string     `json:"status"`
	Used       bool       `json:"used"`
	UsedAt     *time.Time `json:"used_at"`
	GameId     *string    `json:"game_id"`
	CreatedAt  time.Time  `json:"created_at"`
}

type Tickets struct {
	Tickets []Ticket `json:"tickets"`
	Total   *int32   `json:"total"`
}
