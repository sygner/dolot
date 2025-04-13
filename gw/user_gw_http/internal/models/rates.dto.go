package models

type UpdateRatesDTO struct {
	TicketBuyRate            *float64 `json:"ticket_buy_rate,omitempty"`
	TransactionTaxPercentage *float64 `json:"transaction_tax_percentage,omitempty"`
	ImpressionExchangeRate   *uint32  `json:"impression_exchange_rate,omitempty"`
}
