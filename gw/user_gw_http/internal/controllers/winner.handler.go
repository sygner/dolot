package controllers

import (
	"dolott_user_gw_http/internal/services"

	"github.com/gofiber/fiber/v2"
)

type (
	WinnerHandler interface {
		GetWinnersByGameId(*fiber.Ctx) error
	}
	winnerHandler struct {
		winnerService services.WinnerService
	}
)

func NewWinnerHandler(winnerService services.WinnerService) WinnerHandler {
	return &winnerHandler{
		winnerService: winnerService,
	}
}

func (c *winnerHandler) GetWinnersByGameId(ctx *fiber.Ctx) error {
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
