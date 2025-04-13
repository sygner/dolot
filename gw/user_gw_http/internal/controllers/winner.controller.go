package controllers

import (
	"dolott_user_gw_http/internal/models"
	"dolott_user_gw_http/internal/services"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type (
	WinnerController interface {
		GetWinnersByGameId(*fiber.Ctx) error
		GetWinnersByGameIdCount(*fiber.Ctx) error
		UpdateTotalPaidUsers(*fiber.Ctx) error
	}
	winnerController struct {
		winnerService services.WinnerService
	}
)

func NewWinnerController(winnerService services.WinnerService) WinnerController {
	return &winnerController{
		winnerService: winnerService,
	}
}

func (c *winnerController) GetWinnersByGameId(ctx *fiber.Ctx) error {
	gameId := ctx.Params("game_id")
	if gameId == "" {
		return ctx.Status(400).JSON(map[string]interface{}{
			"message": "game id cannot be empty, error code #113",
			"success": false,
		})
	}
	res, err := c.winnerService.GetWinnersByGameId(gameId)
	if err != nil {
		return ctx.Status(err.ErrorToHttpStatus()).JSON(err.ErrorToJsonMessage())
	}
	return ctx.JSON(map[string]interface{}{
		"data":    res,
		"success": true,
	})
}

func (c *winnerController) GetWinnersByGameIdCount(ctx *fiber.Ctx) error {
	gameId := ctx.Params("game_id")
	if gameId == "" {
		return ctx.Status(400).JSON(map[string]interface{}{
			"message": "game id cannot be empty, error code #113",
			"success": false,
		})
	}
	res, err := c.winnerService.GetWinnersByGameIdCount(gameId)
	if err != nil {
		return ctx.Status(err.ErrorToHttpStatus()).JSON(err.ErrorToJsonMessage())
	}
	return ctx.JSON(map[string]interface{}{
		"data":    res,
		"success": true,
	})
}

func (c *winnerController) UpdateTotalPaidUsers(ctx *fiber.Ctx) error {
	updateTotalPaidDTO := new(models.UpdateTotalPaidDTO)
	if err := ctx.BodyParser(updateTotalPaidDTO); err != nil {
		return ctx.Status(400).JSON(map[string]interface{}{
			"message": "wrong body, error code #589-1",
			"success": false,
		})
	}
	err := c.winnerService.UpdateTotalPaidUsers(updateTotalPaidDTO.GameId, updateTotalPaidDTO.TotalPaid)
	if err != nil {
		return ctx.Status(err.ErrorToHttpStatus()).JSON(err.ErrorToJsonMessage())
	}
	fmt.Println("DDDD2 ", updateTotalPaidDTO)
	return ctx.JSON(map[string]interface{}{
		"message": "Updated",
		"success": true,
	})
}
