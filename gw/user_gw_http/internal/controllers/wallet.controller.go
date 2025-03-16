package controllers

import (
	"dolott_user_gw_http/internal/models"
	"dolott_user_gw_http/internal/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type (
	WalletController interface {
		GetAllCoins(*fiber.Ctx) error
		GetCoinById(*fiber.Ctx) error
		GetCoinBySymbol(*fiber.Ctx) error
		CreateWallet(*fiber.Ctx) error
		GetBalanceByCoinIdAndUserId(*fiber.Ctx) error
		GetWalletsByUserId(*fiber.Ctx) error
		GetWalletByAddress(*fiber.Ctx) error
		GetWalletBySid(*fiber.Ctx) error
		GetWalletByWalletId(*fiber.Ctx) error
		GetTransactionByTxId(*fiber.Ctx) error
		GetTransactionsByWalletIdAndUserId(*fiber.Ctx) error
		GetTransactionsByUserId(*fiber.Ctx) error
		AddTransaction(*fiber.Ctx) error
	}
	walletController struct {
		walletService services.WalletService
	}
)

func NewWalletController(walletService services.WalletService) WalletController {
	return &walletController{
		walletService: walletService,
	}
}

func (c *walletController) GetAllCoins(ctx *fiber.Ctx) error {
	res, err := c.walletService.GetAllCoins()
	if err != nil {
		return ctx.Status(err.ErrorToHttpStatus()).JSON(err.ErrorToJsonMessage())
	}
	return ctx.JSON(map[string]interface{}{
		"data":    res,
		"success": true,
	})
}

func (c *walletController) GetCoinById(ctx *fiber.Ctx) error {
	coinIdString := ctx.Params("coinId")
	if coinIdString == "" {
		return ctx.Status(400).JSON(map[string]interface{}{
			"message": "coin id cannot be empty, error code #134",
			"success": false,
		})
	}

	coinIdInt, rerr := strconv.Atoi(coinIdString)
	if rerr != nil {
		return ctx.Status(400).JSON(map[string]interface{}{
			"message": "invalid coin id, must be a number, error code #135",
			"success": false,
		})
	}

	coinId := int32(coinIdInt)

	res, err := c.walletService.GetCoinById(coinId)
	if err != nil {
		return ctx.Status(err.ErrorToHttpStatus()).JSON(err.ErrorToJsonMessage())
	}

	return ctx.JSON(map[string]interface{}{
		"data":    res,
		"success": true,
	})
}

func (c *walletController) GetCoinBySymbol(ctx *fiber.Ctx) error {
	symbol := ctx.Params("symbol")
	if symbol == "" {
		return ctx.Status(400).JSON(map[string]interface{}{
			"message": "symbol cannot be empty, error code #136",
			"success": false,
		})
	}

	res, err := c.walletService.GetCoinBySymbol(symbol)
	if err != nil {
		return ctx.Status(err.ErrorToHttpStatus()).JSON(err.ErrorToJsonMessage())
	}

	return ctx.JSON(map[string]interface{}{
		"data":    res,
		"success": true,
	})
}

func (c *walletController) CreateWallet(ctx *fiber.Ctx) error {
	coinIdString := ctx.Params("coinId")
	if coinIdString == "" {
		return ctx.Status(400).JSON(map[string]interface{}{
			"message": "coin id cannot be empty, error code #137",
			"success": false,
		})
	}

	coinIdInt, rerr := strconv.Atoi(coinIdString)
	if rerr != nil {
		return ctx.Status(400).JSON(map[string]interface{}{
			"message": "invalid coin id, must be a number, error code #138",
			"success": false,
		})
	}

	coinId := int32(coinIdInt)

	localData := ctx.Locals("user_data")
	if localData == nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(map[string]interface{}{
			"message": "internal issue, error code #139",
			"success": false,
		})
	}
	userData := localData.(models.Token)

	res, err := c.walletService.CreateWallet(userData.UserId, coinId)
	if err != nil {
		return ctx.Status(err.ErrorToHttpStatus()).JSON(err.ErrorToJsonMessage())
	}

	return ctx.JSON(map[string]interface{}{
		"data":    res,
		"success": true,
	})
}

func (c *walletController) GetBalanceByCoinIdAndUserId(ctx *fiber.Ctx) error {
	coinIdString := ctx.Params("coinId")
	if coinIdString == "" {
		return ctx.Status(400).JSON(map[string]interface{}{
			"message": "coin id cannot be empty, error code #140",
			"success": false,
		})
	}

	coinIdInt, rerr := strconv.Atoi(coinIdString)
	if rerr != nil {
		return ctx.Status(400).JSON(map[string]interface{}{
			"message": "invalid coin id, must be a number, error code #141",
			"success": false,
		})
	}

	coinId := int32(coinIdInt)

	localData := ctx.Locals("user_data")
	if localData == nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(map[string]interface{}{
			"message": "internal issue, error code #142",
			"success": false,
		})
	}
	userData := localData.(models.Token)

	res, err := c.walletService.GetBalanceByCoinIdAndUserId(userData.UserId, coinId)
	if err != nil {
		return ctx.Status(err.ErrorToHttpStatus()).JSON(err.ErrorToJsonMessage())
	}

	return ctx.JSON(map[string]interface{}{
		"data":    res,
		"success": true,
	})
}

func (c *walletController) GetWalletsByUserId(ctx *fiber.Ctx) error {
	localData := ctx.Locals("user_data")
	if localData == nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(map[string]interface{}{
			"message": "internal issue, error code #143",
			"success": false,
		})
	}
	userData := localData.(models.Token)

	res, err := c.walletService.GetWalletsByUserId(userData.UserId)
	if err != nil {
		return ctx.Status(err.ErrorToHttpStatus()).JSON(err.ErrorToJsonMessage())
	}

	return ctx.JSON(map[string]interface{}{
		"data":    res,
		"success": true,
	})
}

func (c *walletController) GetWalletByAddress(ctx *fiber.Ctx) error {
	address := ctx.Params("address")
	if address == "" {
		return ctx.Status(400).JSON(map[string]interface{}{
			"message": "address cannot be empty, error code #144",
			"success": false,
		})
	}

	res, err := c.walletService.GetWalletByAddress(address)
	if err != nil {
		return ctx.Status(err.ErrorToHttpStatus()).JSON(err.ErrorToJsonMessage())
	}

	return ctx.JSON(map[string]interface{}{
		"data":    res,
		"success": true,
	})
}

func (c *walletController) GetWalletBySid(ctx *fiber.Ctx) error {
	sid := ctx.Params("sid")
	if sid == "" {
		return ctx.Status(400).JSON(map[string]interface{}{
			"message": "sid cannot be empty, error code #145",
			"success": false,
		})
	}

	res, err := c.walletService.GetWalletBySid(sid)
	if err != nil {
		return ctx.Status(err.ErrorToHttpStatus()).JSON(err.ErrorToJsonMessage())
	}

	return ctx.JSON(map[string]interface{}{
		"data":    res,
		"success": true,
	})
}

func (c *walletController) GetWalletByWalletId(ctx *fiber.Ctx) error {
	walletIdString := ctx.Params("wallet_id")
	if walletIdString == "" {
		return ctx.Status(400).JSON(map[string]interface{}{
			"message": "wallet id cannot be empty, error code #146",
			"success": false,
		})
	}

	walletIdInt, rerr := strconv.Atoi(walletIdString)
	if rerr != nil {
		return ctx.Status(400).JSON(map[string]interface{}{
			"message": "invalid coin id, must be a number, error code #147",
			"success": false,
		})
	}

	walletId := int32(walletIdInt)

	res, err := c.walletService.GetWalletByWalletId(walletId)
	if err != nil {
		return ctx.Status(err.ErrorToHttpStatus()).JSON(err.ErrorToJsonMessage())
	}

	return ctx.JSON(map[string]interface{}{
		"data":    res,
		"success": true,
	})
}

func (c *walletController) GetTransactionByTxId(ctx *fiber.Ctx) error {
	tx := ctx.Params("txId")
	if tx == "" {
		return ctx.Status(400).JSON(map[string]interface{}{
			"message": "tx cannot be empty, error code #148",
			"success": false,
		})
	}

	res, err := c.walletService.GetTransactionByTxId(tx)
	if err != nil {
		return ctx.Status(err.ErrorToHttpStatus()).JSON(err.ErrorToJsonMessage())
	}

	return ctx.JSON(map[string]interface{}{
		"data":    res,
		"success": true,
	})
}

func (c *walletController) GetTransactionsByWalletIdAndUserId(ctx *fiber.Ctx) error {
	walletIdString := ctx.Params("wallet_id")
	if walletIdString == "" {
		return ctx.Status(400).JSON(map[string]interface{}{
			"message": "wallet id cannot be empty, error code #149",
			"success": false,
		})
	}

	walletIdInt, rerr := strconv.Atoi(walletIdString)
	if rerr != nil {
		return ctx.Status(400).JSON(map[string]interface{}{
			"message": "invalid coin id, must be a number, error code #150",
			"success": false,
		})
	}

	walletId := int32(walletIdInt)

	localData := ctx.Locals("user_data")
	if localData == nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(map[string]interface{}{
			"message": "internal issue, error code #151",
			"success": false,
		})
	}
	userData := localData.(models.Token)

	res, err := c.walletService.GetTransactionsByWalletIdAndUserId(walletId, userData.UserId)
	if err != nil {
		return ctx.Status(err.ErrorToHttpStatus()).JSON(err.ErrorToJsonMessage())
	}

	return ctx.JSON(map[string]interface{}{
		"data":    res,
		"success": true,
	})
}

func (c *walletController) GetTransactionsByUserId(ctx *fiber.Ctx) error {
	localData := ctx.Locals("user_data")
	if localData == nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(map[string]interface{}{
			"message": "internal issue, error code #152",
			"success": false,
		})
	}
	userData := localData.(models.Token)
	res, err := c.walletService.GetTransactionsByUserId(userData.UserId)
	if err != nil {
		return ctx.Status(err.ErrorToHttpStatus()).JSON(err.ErrorToJsonMessage())
	}

	return ctx.JSON(map[string]interface{}{
		"data":    res,
		"success": true,
	})
}

func (c *walletController) AddTransaction(ctx *fiber.Ctx) error {
	addTransactionDTO := new(models.AddTransactionDTO)
	if err := ctx.BodyParser(addTransactionDTO); err != nil {
		return ctx.Status(400).JSON(map[string]interface{}{
			"message": "wrong body, error code #101",
			"success": false,
		})
	}

	localData := ctx.Locals("user_data")
	if localData == nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(map[string]interface{}{
			"message": "internal issue, error code #153",
			"success": false,
		})
	}
	userData := localData.(models.Token)
	addTransactionDTO.FromWalletUserId = userData.UserId

	res, err := c.walletService.AddTransaction(addTransactionDTO)
	if err != nil {
		return ctx.Status(err.ErrorToHttpStatus()).JSON(err.ErrorToJsonMessage())
	}

	return ctx.JSON(map[string]interface{}{
		"data":    res,
		"success": true,
	})

}
