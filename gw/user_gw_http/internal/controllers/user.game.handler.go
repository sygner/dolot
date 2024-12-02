package controllers

import (
	"dolott_user_gw_http/internal/models"
	"dolott_user_gw_http/internal/services"

	"github.com/gofiber/fiber/v2"
)

type (
	UserGameHandler interface {
		AddUserChoice(*fiber.Ctx) error
		GetUserChoicesByUserId(*fiber.Ctx) error
		GetUserChoicesByUserIdAndTimeRange(*fiber.Ctx) error
		GetUserChoicesByGameIdAndPagination(*fiber.Ctx) error
		GetAllUserGames(*fiber.Ctx) error
	}
	userGameHandler struct {
		userGameService services.UserGameService
	}
)

func NewUserGameHandler(userGameService services.UserGameService) UserGameHandler {
	return &userGameHandler{
		userGameService: userGameService,
	}
}

func (c *userGameHandler) AddUserChoice(ctx *fiber.Ctx) error {
	addUserChoiceDTO := new(models.AddUserChoiceDTO)
	if err := ctx.BodyParser(addUserChoiceDTO); err != nil {
		return ctx.Status(400).JSON(map[string]interface{}{
			"message": "wrong body, error code #121",
			"success": false,
		})
	}

	localData := ctx.Locals("user_data")
	if localData == nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(map[string]interface{}{
			"message": "internal issue, error code #122",
			"success": false,
		})
	}
	userData := localData.(models.Token)
	addUserChoiceDTO.UserId = userData.UserId
	res, err := c.userGameService.AddUserChoice(addUserChoiceDTO)
	if err != nil {
		return ctx.Status(err.ErrorToHttpStatus()).JSON(err.ErrorToJsonMessage())
	}
	return ctx.JSON(map[string]interface{}{
		"data":    res,
		"success": true,
	})
}

func (c *userGameHandler) GetUserChoicesByUserId(ctx *fiber.Ctx) error {
	paginationDTO := new(models.Pagination)
	if err := ctx.BodyParser(paginationDTO); err != nil {
		return ctx.Status(400).JSON(map[string]interface{}{
			"message": "wrong body, error code #122",
			"success": false,
		})
	}
	localData := ctx.Locals("user_data")
	if localData == nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(map[string]interface{}{
			"message": "internal issue, error code #126",
			"success": false,
		})
	}
	userData := localData.(models.Token)
	res, err := c.userGameService.GetUserChoicesByUserId(userData.UserId, paginationDTO)
	if err != nil {
		return ctx.Status(err.ErrorToHttpStatus()).JSON(err.ErrorToJsonMessage())
	}
	return ctx.JSON(map[string]interface{}{
		"data":    res,
		"success": true,
	})
}

func (c *userGameHandler) GetUserChoicesByUserIdAndTimeRange(ctx *fiber.Ctx) error {
	userChoicesByUserIdAndTimeRangeDTO := new(models.GetUserChoicesByUserIdAndTimeRangeDTO)
	if err := ctx.BodyParser(userChoicesByUserIdAndTimeRangeDTO); err != nil {
		return ctx.Status(400).JSON(map[string]interface{}{
			"message": "wrong body, error code #123",
			"success": false,
		})
	}
	localData := ctx.Locals("user_data")
	if localData == nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(map[string]interface{}{
			"message": "internal issue, error code #127",
			"success": false,
		})
	}
	userData := localData.(models.Token)
	userChoicesByUserIdAndTimeRangeDTO.UserId = userData.UserId
	res, err := c.userGameService.GetUserChoicesByUserIdAndTimeRange(userChoicesByUserIdAndTimeRangeDTO)
	if err != nil {
		return ctx.Status(err.ErrorToHttpStatus()).JSON(err.ErrorToJsonMessage())
	}
	return ctx.JSON(map[string]interface{}{
		"data":    res,
		"success": true,
	})
}

func (c *userGameHandler) GetUserChoicesByGameIdAndPagination(ctx *fiber.Ctx) error {
	gameId := ctx.Params("game_id")
	if gameId == "" {
		return ctx.Status(400).JSON(map[string]interface{}{
			"message": "game id cannot be empty, error code #124",
			"success": false,
		})
	}
	paginationDTO := new(models.Pagination)
	if err := ctx.BodyParser(paginationDTO); err != nil {
		return ctx.Status(400).JSON(map[string]interface{}{
			"message": "wrong body, error code #125",
			"success": false,
		})
	}
	res, err := c.userGameService.GetUserChoicesByGameIdAndPagination(gameId, paginationDTO)
	if err != nil {
		return ctx.Status(err.ErrorToHttpStatus()).JSON(err.ErrorToJsonMessage())
	}
	return ctx.JSON(map[string]interface{}{
		"data":    res,
		"success": true,
	})
}

func (c *userGameHandler) GetAllUserGames(ctx *fiber.Ctx) error {
	localData := ctx.Locals("user_data")
	if localData == nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(map[string]interface{}{
			"message": "internal issue, error code #133",
			"success": false,
		})
	}
	userData := localData.(models.Token)
	res, err := c.userGameService.GetAllUserGames(userData.UserId)
	if err != nil {
		return ctx.Status(err.ErrorToHttpStatus()).JSON(err.ErrorToJsonMessage())
	}
	return ctx.JSON(map[string]interface{}{
		"data":    res,
		"success": true,
	})
}
