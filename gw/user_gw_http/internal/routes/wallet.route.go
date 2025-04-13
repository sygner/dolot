package routes

import (
	"dolott_user_gw_http/internal/controllers"
	"dolott_user_gw_http/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func WalletGroup(app fiber.Router, walletController controllers.WalletController, middleware middleware.MiddlewareService) {
	walletGroup := app.Group("/wlt")
	walletGroup.Use(middleware.VerificationMiddleware)

	walletGroup.Get("/coin/all", walletController.GetAllCoins)
	walletGroup.Get("/coin/id/:coinId", walletController.GetCoinById)
	walletGroup.Get("/coin/symbol/:symbol", walletController.GetCoinBySymbol)

	walletGroup.Get("/wallet/build/:coinId", walletController.CreateWallet)
	walletGroup.Get("/wallet/balance/:coinId", walletController.GetBalanceByCoinIdAndUserId)
	walletGroup.Get("/wallets", walletController.GetWalletsByUserId)
	walletGroup.Get("/wallet/address/:address", walletController.GetWalletByAddress)
	walletGroup.Get("/wallet/sid/:sid", walletController.GetWalletBySid)
	walletGroup.Get("/wallet/id/:wallet_id", walletController.GetWalletByWalletId)

	walletGroup.Get("/transaction/txid/:txId", walletController.GetTransactionByTxId)
	walletGroup.Get("/transaction/wallet/:wallet_id", walletController.GetTransactionsByWalletIdAndUserId)
	walletGroup.Post("/transaction/p/wallet", walletController.GetTransactionsByWalletIdAndUserIdAndPagination)
	walletGroup.Post("/transaction/p/usr", walletController.GetTransactionsByUserIdAndPagination)
	walletGroup.Get("/transaction/usr", walletController.GetTransactionsByUserId)
	walletGroup.Get("/transaction/pre", walletController.GetPreTransactionDetail)
	walletGroup.Post("/transaction/add", walletController.AddTransaction)
}
