package controllers

import (
	"dolott_user_gw_http/internal/models"
	"dolott_user_gw_http/internal/services"

	"github.com/gofiber/fiber/v2"
)

type (
	GameHandler interface {
		GetGameByGameId(*fiber.Ctx) error
		AddGame(*fiber.Ctx) error
		GetNextGamesByGameType(*fiber.Ctx) error
		DeleteGameByGameId(*fiber.Ctx) error
		CheckGameExistsByGameId(*fiber.Ctx) error
		GetGamesByCreatorId(*fiber.Ctx) error
		AddResultByGameId(*fiber.Ctx) error
		GetAllNextGames(*fiber.Ctx) error
		GetAllPreviousGames(*fiber.Ctx) error
		GetAllGames(*fiber.Ctx) error
	}
	gameHandler struct {
		gameService services.GameService
	}
)

func NewGameHandler(gameService services.GameService) GameHandler {
	return &gameHandler{
		gameService: gameService,
	}
}

func (c *gameHandler) GetGameByGameId(ctx *fiber.Ctx) error {
	gameId := ctx.Params("game_id")
	if gameId == "" {
		return ctx.Status(400).JSON(map[string]interface{}{
			"message": "game id cannot be empty, error code #114",
			"success": false,
		})
	}

	res, err := c.gameService.GetGameByGameId(gameId)
	if err != nil {
		return ctx.Status(err.ErrorToHttpStatus()).JSON(err.ErrorToJsonMessage())
	}
	return ctx.JSON(map[string]interface{}{
		"data":    res,
		"success": true,
	})
}

func (c *gameHandler) AddGame(ctx *fiber.Ctx) error {
	addGameDTO := new(models.AddGameDTO)
	if err := ctx.BodyParser(addGameDTO); err != nil {
		return ctx.Status(400).JSON(map[string]interface{}{
			"message": "wrong body, error code #115",
			"success": false,
		})
	}

	localData := ctx.Locals("user_data")
	if localData == nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(map[string]interface{}{
			"message": "internal issue, error code #116",
			"success": false,
		})
	}
	userData := localData.(models.Token)
	addGameDTO.CreatorId = userData.UserId

	res, err := c.gameService.AddGame(addGameDTO)
	if err != nil {
		return ctx.Status(err.ErrorToHttpStatus()).JSON(err.ErrorToJsonMessage())
	}
	return ctx.JSON(map[string]interface{}{
		"data":    res,
		"success": true,
	})
}

func (c *gameHandler) GetNextGamesByGameType(ctx *fiber.Ctx) error {
	gameType := ctx.QueryInt("game_type")
	limit := ctx.QueryInt("limit")
	res, err := c.gameService.GetNextGamesByGameType(int32(gameType), int32(limit))
	if err != nil {
		return ctx.Status(err.ErrorToHttpStatus()).JSON(err.ErrorToJsonMessage())
	}

	return ctx.JSON(map[string]interface{}{
		"data":    res,
		"success": true,
	})
}

func (c *gameHandler) GetAllNextGames(ctx *fiber.Ctx) error {
	res, err := c.gameService.GetAllNextGames()
	if err != nil {
		return ctx.Status(err.ErrorToHttpStatus()).JSON(err.ErrorToJsonMessage())
	}

	return ctx.JSON(map[string]interface{}{
		"data":    res,
		"success": true,
	})
}

func (c *gameHandler) GetAllPreviousGames(ctx *fiber.Ctx) error {
	paginationDTO := new(models.Pagination)
	if err := ctx.BodyParser(paginationDTO); err != nil {
		return ctx.Status(400).JSON(map[string]interface{}{
			"message": "wrong body, error code #121",
			"success": false,
		})
	}
	res, err := c.gameService.GetAllPreviousGames(paginationDTO)
	if err != nil {
		return ctx.Status(err.ErrorToHttpStatus()).JSON(err.ErrorToJsonMessage())
	}

	return ctx.JSON(map[string]interface{}{
		"data":    res,
		"success": true,
	})
}

func (c *gameHandler) GetAllGames(ctx *fiber.Ctx) error {
	paginationDTO := new(models.Pagination)
	if err := ctx.BodyParser(paginationDTO); err != nil {
		return ctx.Status(400).JSON(map[string]interface{}{
			"message": "wrong body, error code #121",
			"success": false,
		})
	}
	res, err := c.gameService.GetAllGames(paginationDTO)
	if err != nil {
		return ctx.Status(err.ErrorToHttpStatus()).JSON(err.ErrorToJsonMessage())
	}

	return ctx.JSON(map[string]interface{}{
		"data":    res,
		"success": true,
	})
}

func (c *gameHandler) DeleteGameByGameId(ctx *fiber.Ctx) error {
	gameId := ctx.Params("game_id")
	if gameId == "" {
		return ctx.Status(400).JSON(map[string]interface{}{
			"message": "game id cannot be empty, error code #117",
			"success": false,
		})
	}
	err := c.gameService.DeleteGameByGameId(gameId)
	if err != nil {
		return ctx.Status(err.ErrorToHttpStatus()).JSON(err.ErrorToJsonMessage())
	}
	return ctx.JSON(map[string]interface{}{
		"message": "game deleted",
		"success": true,
	})
}

func (c *gameHandler) CheckGameExistsByGameId(ctx *fiber.Ctx) error {
	gameId := ctx.Params("game_id")
	if gameId == "" {
		return ctx.Status(400).JSON(map[string]interface{}{
			"message": "game id cannot be empty, error code #118",
			"success": false,
		})
	}

	err := c.gameService.CheckGameExistsGameId(gameId)
	if err != nil {
		if err.Code == 5 {
			return ctx.JSON(map[string]interface{}{
				"data":    false,
				"success": true,
			})
		} else {
			return ctx.Status(err.ErrorToHttpStatus()).JSON(err.ErrorToJsonMessage())
		}
	}
	return ctx.JSON(map[string]interface{}{
		"data":    true,
		"success": true,
	})
}

func (c *gameHandler) GetGamesByCreatorId(ctx *fiber.Ctx) error {
	userId := ctx.QueryInt("id")
	paginationDTO := new(models.Pagination)
	if err := ctx.BodyParser(paginationDTO); err != nil {
		return ctx.Status(400).JSON(map[string]interface{}{
			"message": "wrong body, error code #119",
			"success": false,
		})
	}
	res, err := c.gameService.GetGamesByCreatorId(int32(userId), paginationDTO)
	if err != nil {
		return ctx.Status(err.ErrorToHttpStatus()).JSON(err.ErrorToJsonMessage())
	}
	return ctx.JSON(map[string]interface{}{
		"data":    res,
		"success": true,
	})
}

func (c *gameHandler) AddResultByGameId(ctx *fiber.Ctx) error {
	addGameResultDTO := new(models.AddGameResultDTO)
	if err := ctx.BodyParser(addGameResultDTO); err != nil {
		return ctx.Status(400).JSON(map[string]interface{}{
			"message": "wrong body, error code #120",
			"success": false,
		})
	}
	res, err := c.gameService.AddResultByGameId(addGameResultDTO.GameId, addGameResultDTO.Result)
	if err != nil {
		return ctx.Status(err.ErrorToHttpStatus()).JSON(err.ErrorToJsonMessage())
	}
	return ctx.JSON(map[string]interface{}{
		"data":    res,
		"success": true,
	})
}
