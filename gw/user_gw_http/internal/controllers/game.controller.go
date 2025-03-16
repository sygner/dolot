package controllers

import (
	"dolott_user_gw_http/internal/models"
	"dolott_user_gw_http/internal/services"
	"net/http"
	"os"
	"path/filepath"

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
		GetAllGameTypes(*fiber.Ctx) error
		DownloadGameTypeFile(*fiber.Ctx) error
		ChangeGameDetailCalculation(*fiber.Ctx) error
		GetAllUserPreviousGames(*fiber.Ctx) error
		GetAllUserPreviousGamesByGameType(*fiber.Ctx) error
		GetAllUserChoiceDivisionsByGameId(*fiber.Ctx) error
		GetAllUsersChoiceDivisionsByGameId(*fiber.Ctx) error
		UpdateGamePrizeByGameId(*fiber.Ctx) error
	}
	gameHandler struct {
		gameService     services.GameService
		fileStoragePath string
	}
)

func NewGameHandler(gameService services.GameService, fileStoragePath string) GameHandler {
	return &gameHandler{
		gameService:     gameService,
		fileStoragePath: fileStoragePath,
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

func (c *gameHandler) DownloadGameTypeFile(ctx *fiber.Ctx) error {
	fileName := ctx.Query("path", "")
	if fileName == "" {
		return ctx.Status(http.StatusBadRequest).SendString("File name is required")
	}

	realPath := filepath.Join(c.fileStoragePath+"/pictures/", fileName)

	absPath, err := filepath.Abs(realPath)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).SendString("Invalid file path")
	}

	if _, err := os.Stat(absPath); os.IsNotExist(err) {
		return ctx.Status(http.StatusNotFound).SendString("File not found")
	}

	return ctx.SendFile(absPath)
}

func (c *gameHandler) GetAllGameTypes(ctx *fiber.Ctx) error {
	res, err := c.gameService.GetAllGameTypes()
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

func (c *gameHandler) ChangeGameDetailCalculation(ctx *fiber.Ctx) error {
	changeGameDetailCalculationDTO := new(models.ChangeGameDetailCalculationDTO)
	if err := ctx.BodyParser(changeGameDetailCalculationDTO); err != nil {
		return ctx.Status(400).JSON(map[string]interface{}{
			"message": "wrong body, error code #176",
			"success": false,
		})
	}

	if changeGameDetailCalculationDTO.PrizeReward == nil {
		*changeGameDetailCalculationDTO.PrizeReward = 0
	}

	if changeGameDetailCalculationDTO.TokenBurn == nil {
		*changeGameDetailCalculationDTO.TokenBurn = 0
	}

	gameType := ctx.QueryInt("game_type")

	res, err := c.gameService.UpdateGameTypeDetail(int32(gameType), changeGameDetailCalculationDTO.DayName, *changeGameDetailCalculationDTO.PrizeReward, *changeGameDetailCalculationDTO.TokenBurn, changeGameDetailCalculationDTO.AutoCompute)

	if err != nil {
		return ctx.Status(err.ErrorToHttpStatus()).JSON(err.ErrorToJsonMessage())
	}
	return ctx.JSON(map[string]interface{}{
		"data":    res,
		"success": true,
	})
}

func (c *gameHandler) GetAllUserPreviousGames(ctx *fiber.Ctx) error {
	paginationDTO := new(models.Pagination)
	if err := ctx.BodyParser(paginationDTO); err != nil {
		return ctx.Status(400).JSON(map[string]interface{}{
			"message": "wrong body, error code #177",
			"success": false,
		})
	}

	localData := ctx.Locals("user_data")
	if localData == nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(map[string]interface{}{
			"message": "internal issue, error code #178",
			"success": false,
		})
	}
	userData := localData.(models.Token)

	res, err := c.gameService.GetAllUserPreviousGames(userData.UserId, paginationDTO)
	if err != nil {
		return ctx.Status(err.ErrorToHttpStatus()).JSON(err.ErrorToJsonMessage())
	}

	return ctx.JSON(map[string]interface{}{
		"data":    res,
		"success": true,
	})
}

func (c *gameHandler) GetAllUserPreviousGamesByGameType(ctx *fiber.Ctx) error {
	getAllUserPreviousGamesByGameTypeDTO := new(models.GetAllUserPreviousGamesByGameTypeDTO)
	if err := ctx.BodyParser(getAllUserPreviousGamesByGameTypeDTO); err != nil {
		return ctx.Status(400).JSON(map[string]interface{}{
			"message": "wrong body, error code #183",
			"success": false,
		})
	}

	localData := ctx.Locals("user_data")
	if localData == nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(map[string]interface{}{
			"message": "internal issue, error code #184",
			"success": false,
		})
	}
	userData := localData.(models.Token)
	gameType := ""
	switch getAllUserPreviousGamesByGameTypeDTO.GameType {
	case 0:
		gameType = "lotto"
	case 1:
		gameType = "ozlotto"
	case 2:
		gameType = "powerball"
	case 3:
		gameType = "american_powerball"
	}
	res, err := c.gameService.GetAllUserPreviousGamesByGameType(userData.UserId, gameType, &getAllUserPreviousGamesByGameTypeDTO.Pagination)
	if err != nil {
		return ctx.Status(err.ErrorToHttpStatus()).JSON(err.ErrorToJsonMessage())
	}

	return ctx.JSON(map[string]interface{}{
		"data":    res,
		"success": true,
	})
}

func (c *gameHandler) GetAllUserChoiceDivisionsByGameId(ctx *fiber.Ctx) error {

	gameId := ctx.Params("game_id")
	if gameId == "" {
		return ctx.Status(400).JSON(map[string]interface{}{
			"message": "game id cannot be empty, error code #179",
			"success": false,
		})
	}

	localData := ctx.Locals("user_data")
	if localData == nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(map[string]interface{}{
			"message": "internal issue, error code #180",
			"success": false,
		})
	}
	userData := localData.(models.Token)

	res, err := c.gameService.GetAllUserChoiceDivisionsByGameId(userData.UserId, gameId)
	if err != nil {
		return ctx.Status(err.ErrorToHttpStatus()).JSON(err.ErrorToJsonMessage())
	}

	return ctx.JSON(map[string]interface{}{
		"data":    res,
		"success": true,
	})
}

func (c *gameHandler) GetAllUsersChoiceDivisionsByGameId(ctx *fiber.Ctx) error {
	gameId := ctx.Params("game_id")
	if gameId == "" {
		return ctx.Status(400).JSON(map[string]interface{}{
			"message": "game id cannot be empty, error code #181",
			"success": false,
		})
	}

	res, err := c.gameService.GetAllUsersChoiceDivisionsByGameId(gameId)
	if err != nil {
		return ctx.Status(err.ErrorToHttpStatus()).JSON(err.ErrorToJsonMessage())
	}

	return ctx.JSON(map[string]interface{}{
		"data":    res,
		"success": true,
	})
}

func (c *gameHandler) UpdateGamePrizeByGameId(ctx *fiber.Ctx) error {
	updateGamePrizeDTO := new(models.UpdateGamePrizeDTO)
	if err := ctx.BodyParser(updateGamePrizeDTO); err != nil {
		return ctx.Status(400).JSON(map[string]interface{}{
			"message": "wrong body, error code #182",
			"success": false,
		})
	}

	err := c.gameService.UpdateGamePrizeByGameId(updateGamePrizeDTO.GameId, updateGamePrizeDTO.Prize, updateGamePrizeDTO.AutoCompute)
	if err != nil {
		return ctx.Status(err.ErrorToHttpStatus()).JSON(err.ErrorToJsonMessage())
	}

	return ctx.JSON(map[string]interface{}{
		"message": "updated",
		"success": true,
	})
}
