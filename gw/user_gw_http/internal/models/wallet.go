package models

import "time"

type Coin struct {
	CoinId int32  `json:"coin_id"`
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}

type Wallet struct {
	Id      int32   `json:"wallet_id"`
	Sid     string  `json:"sid"`
	UserId  int32   `json:"user_id"`
	CoinId  int32   `json:"coin_id"`
	Balance float64 `json:"balance"`
	Address string  `json:"address"`
	// PublicKey string  `json:"public_key"`
	// PrivateKey string  `json:"private_key"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Transaction struct {
	TxId          string    `json:"tx_id"`
	CurrencyId    int32     `json:"currency_id"`
	CurrencyName  string    `json:"currency_name"`
	FromAddress   string    `json:"from_address"`
	ToAddress     string    `json:"to_address"`
	FromWalletId  int32     `json:"from_wallet_id"`
	FromPublicKey string    `json:"from_public_key"`
	Amount        float64   `json:"amount"`
	TransactionAt time.Time `json:"transaction_at"`
}
