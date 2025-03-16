package models

type AddTicketDTO struct {
	Signature  string `json:"singature"`
	UserId     int32  `json:"user_id"`
	TicketType string `json:"ticket_type"`
	Status     string `json:"status"`
}

type Pagination struct {
	Offset int32 `json:"offset"`
	Limit  int32 `json:"limit"`
	Total  bool  `json:"total"`
}
