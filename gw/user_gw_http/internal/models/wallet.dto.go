package models

type AddTransactionDTO struct {
	Amount           float64 `json:"amount"`
	FromWalletId     int32   `json:"from_wallet_id"`
	FromWalletUserId int32
	ToWalletAddress  string `json:"to_wallet_address"`
	CoinId           int32  `json:"coin_id"`
}

type CreateWalletDTO struct {
	UserId int32 `json:"user_id"`
	CoinId int32 `json:"coin_id"`
}

type GetBalanceByCoinIdAndUserIdDTO struct {
	UserId int32 `json:"user_id"`
	CoinId int32 `json:"coin_id"`
}

type GetTransactionsByUserIdAndPaginationDTO struct {
	WalletId   *int32      `json:"wallet_id,omitempty"`
	Pagination *Pagination `json:"pagination"`
}
