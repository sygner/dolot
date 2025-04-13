package controllers

import (
	"dolott_user_gw_http/internal/admin"
	"dolott_user_gw_http/internal/constants"
	"dolott_user_gw_http/internal/models"
	"dolott_user_gw_http/internal/services"

	"github.com/gofiber/fiber/v2"
)

type (
	ProfileController interface {
		GetProfileByUsername(*fiber.Ctx) error
		GetSelfProfile(*fiber.Ctx) error
		UpdateProfile(*fiber.Ctx) error
		GetProfileBySid(*fiber.Ctx) error
		ChangeUserImpression(*fiber.Ctx) error
		ImpressionExchange(*fiber.Ctx) error
		SearchUsername(*fiber.Ctx) error
		GetAllUserRanking(*fiber.Ctx) error
		GetUserLeaderBoard(cx *fiber.Ctx) error
		UpdateUserRank(*fiber.Ctx) error
		ChangeImpressionAndDCoin(*fiber.Ctx) error
	}
	profileController struct {
		profileService services.ProfileService
		walletService  services.WalletService
	}
)

func NewProfileController(profileService services.ProfileService, walletService services.WalletService) ProfileController {
	return &profileController{
		profileService: profileService,
		walletService:  walletService,
	}
}

func (c *profileController) GetProfileByUsername(ctx *fiber.Ctx) error {
	username := ctx.Params("username")
	if username == "" {
		return ctx.Status(400).JSON(map[string]interface{}{
			"message": "username cannot be empty, error code #128",
			"success": false,
		})
	}
	res, err := c.profileService.GetProfileUsername(username)
	if err != nil {
		return ctx.Status(err.ErrorToHttpStatus()).JSON(err.ErrorToJsonMessage())
	}
	return ctx.JSON(map[string]interface{}{
		"data":    res,
		"success": true,
	})
}
func (c *profileController) GetProfileBySid(ctx *fiber.Ctx) error {
	sid := ctx.Params("sid")
	if sid == "" {
		return ctx.Status(400).JSON(map[string]interface{}{
			"message": "sid cannot be empty, error code #132",
			"success": false,
		})
	}
	res, err := c.profileService.GetProfileSid(sid)
	if err != nil {
		return ctx.Status(err.ErrorToHttpStatus()).JSON(err.ErrorToJsonMessage())
	}
	return ctx.JSON(map[string]interface{}{
		"data":    res,
		"success": true,
	})
}

func (c *profileController) GetSelfProfile(ctx *fiber.Ctx) error {
	localData := ctx.Locals("user_data")
	if localData == nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(map[string]interface{}{
			"message": "internal issue, error code #129",
			"success": false,
		})
	}
	userData := localData.(models.Token)
	res, err := c.profileService.GetProfileByUserId(userData.UserId)
	if err != nil {
		return ctx.Status(err.ErrorToHttpStatus()).JSON(err.ErrorToJsonMessage())
	}

	wallet, err := c.walletService.GetWalletsByUserId(userData.UserId)
	if err != nil {
		return ctx.Status(err.ErrorToHttpStatus()).JSON(err.ErrorToJsonMessage())
	}

	price, err := constants.GetLUNCPriceCoinPaprika()
	if err != nil {
		price = 0
	}
	// Format the float to avoid scientific notation
	return ctx.JSON(map[string]interface{}{
		"data":                     res,
		"wallet":                   wallet,
		"lunc_price":               price,
		"ticket_rate":              admin.TICKET_BUY_RATE,
		"impression_exchange_rate": admin.IMPRESSION_EXCHANGE_RATE,
		"success":                  true,
	})
}

func (c *profileController) UpdateProfile(ctx *fiber.Ctx) error {
	updateProfileDTO := new(models.UpdateProfileDTO)
	if err := ctx.BodyParser(updateProfileDTO); err != nil {
		return ctx.Status(400).JSON(map[string]interface{}{
			"message": "wrong body, error code #130",
			"success": false,
		})
	}
	localData := ctx.Locals("user_data")
	if localData == nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(map[string]interface{}{
			"message": "internal issue, error code #131",
			"success": false,
		})
	}
	userData := localData.(models.Token)
	err := c.profileService.UpdateProfile(userData.UserId, updateProfileDTO.Username)
	if err != nil {
		return ctx.Status(err.ErrorToHttpStatus()).JSON(err.ErrorToJsonMessage())
	}
	return ctx.JSON(map[string]interface{}{
		"message": "profile updated",
		"success": true,
	})
}

func (c *profileController) ChangeUserImpression(ctx *fiber.Ctx) error {
	amount := ctx.QueryInt("amount", 0)
	if amount == 0 {
		return ctx.Status(400).JSON(map[string]interface{}{
			"message": "amount cannot be empty, error code #199",
			"success": false,
		})
	}
	localData := ctx.Locals("user_data")
	if localData == nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(map[string]interface{}{
			"message": "internal issue, error code #200",
			"success": false,
		})
	}
	userData := localData.(models.Token)

	err := c.profileService.ChangeUserImpression(userData.UserId, int32(amount), true)
	if err != nil {
		return ctx.Status(err.ErrorToHttpStatus()).JSON(err.ErrorToJsonMessage())
	}
	return ctx.JSON(map[string]interface{}{
		"message": "impression updated",
		"success": true,
	})
}
func (c *profileController) ImpressionExchange(ctx *fiber.Ctx) error {
	impressionExchangeDTO := new(models.ImpressionExchangeDTO)
	if err := ctx.BodyParser(impressionExchangeDTO); err != nil {
		return ctx.Status(400).JSON(map[string]interface{}{
			"message": "wrong body, error code #197",
			"success": false,
		})
	}
	localData := ctx.Locals("user_data")
	if localData == nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(map[string]interface{}{
			"message": "internal issue, error code #198",
			"success": false,
		})
	}
	userData := localData.(models.Token)
	res, err := c.profileService.ImpressionExchange(userData.UserId, impressionExchangeDTO.Impression, impressionExchangeDTO.ToCoin)
	if err != nil {
		return ctx.Status(err.ErrorToHttpStatus()).JSON(err.ErrorToJsonMessage())
	}
	return ctx.JSON(map[string]interface{}{
		"data":    res,
		"success": true,
	})
}

func (c *profileController) SearchUsername(ctx *fiber.Ctx) error {
	username := ctx.Query("username", "")
	if username == "" {
		return ctx.Status(400).JSON(map[string]interface{}{
			"message": "username cannot be empty, error code #201",
			"success": false,
		})
	}
	if username == "" || len(username) < 3 {
		return ctx.Status(400).JSON(map[string]interface{}{
			"message": "username must be at least 3 characters, error code #202",
			"success": false,
		})
	}

	res, err := c.profileService.SearchUsername(username)
	if err != nil {
		return ctx.Status(err.ErrorToHttpStatus()).JSON(err.ErrorToJsonMessage())
	}
	return ctx.JSON(map[string]interface{}{
		"data":    res,
		"success": true,
	})
}

func (c *profileController) GetAllUserRanking(ctx *fiber.Ctx) error {
	localData := ctx.Locals("user_data")
	if localData == nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(map[string]interface{}{
			"message": "internal issue, error code #203",
			"success": false,
		})
	}
	userData := localData.(models.Token)
	res, err := c.profileService.GetAllUserRanking(userData.UserId)
	if err != nil {
		return ctx.Status(err.ErrorToHttpStatus()).JSON(err.ErrorToJsonMessage())
	}
	return ctx.JSON(map[string]interface{}{
		"data":    res,
		"success": true,
	})
}

func (c *profileController) GetUserLeaderBoard(ctx *fiber.Ctx) error {
	localData := ctx.Locals("user_data")
	if localData == nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(map[string]interface{}{
			"message": "internal issue, error code #204",
			"success": false,
		})
	}
	userData := localData.(models.Token)
	res, err := c.profileService.GetUserLeaderBoard(userData.UserId)
	if err != nil {
		return ctx.Status(err.ErrorToHttpStatus()).JSON(err.ErrorToJsonMessage())
	}
	return ctx.JSON(map[string]interface{}{
		"data":    res,
		"success": true,
	})
}

func (c *profileController) UpdateUserRank(ctx *fiber.Ctx) error {
	amount := ctx.QueryInt("amount", 0)
	increment := ctx.QueryBool("increment", false)

	localData := ctx.Locals("user_data")
	if localData == nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(map[string]interface{}{
			"message": "internal issue, error code #204",
			"success": false,
		})
	}
	userData := localData.(models.Token)

	err := c.profileService.ChangeUserRank(userData.UserId, int32(amount), increment)
	if err != nil {
		return ctx.Status(err.ErrorToHttpStatus()).JSON(err.ErrorToJsonMessage())
	}
	return ctx.JSON(map[string]interface{}{
		"message": "rank changed",
		"success": true,
	})
}

func (c *profileController) ChangeImpressionAndDCoin(ctx *fiber.Ctx) error {
	dCoinAmount := ctx.QueryInt("d_coin_amount", 0)
	ImpressionAmount := ctx.QueryInt("impression_amount", 0)
	increment := ctx.QueryBool("increment", false)

	localData := ctx.Locals("user_data")
	if localData == nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(map[string]interface{}{
			"message": "internal issue, error code #204",
			"success": false,
		})
	}
	userData := localData.(models.Token)

	err := c.profileService.ChangeImpressionAndDCoin(userData.UserId, int32(ImpressionAmount), int32(dCoinAmount), increment)
	if err != nil {
		return ctx.Status(err.ErrorToHttpStatus()).JSON(err.ErrorToJsonMessage())
	}
	return ctx.JSON(map[string]interface{}{
		"message": "impression changed",
		"success": true,
	})
}
