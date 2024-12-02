package services

import (
	"dolott_game/internal/models"
	"dolott_game/internal/repository"
	"dolott_game/internal/types"
	"safir/libs/idgen"
	"time"
)

type (
	GameServices interface {
		GetGameByGameId(string) (*models.Game, *types.Error)
		AddGame(*models.AddGameDTO) (*models.Game, *types.Error)
		GetNextGamesByGameType(int32, int32) ([]models.Game, *types.Error)
		DeleteGameByGameId(string) *types.Error
		CheckGameExistsById(string) (bool, *types.Error)
		GetGamesByCreatorId(int32, *models.Pagination) (*models.Games, *types.Error)
		AddResultByGameId(string, string) ([]models.DivisionResult, *types.Error)
		GetAllNextGames() ([]models.Game, *types.Error)
		GetAllPreviousGames(*models.Pagination) (*models.Games, *types.Error)
		GetAllGames(*models.Pagination) (*models.Games, *types.Error)
	}
	gameServices struct {
		repository repository.GameRepository
	}
)

func NewGameServices(repository repository.GameRepository) GameServices {
	return &gameServices{
		repository: repository,
	}
}

func (c *gameServices) GetGameByGameId(gameId string) (*models.Game, *types.Error) {
	return c.repository.GetGameByGameId(gameId)
}

func (c *gameServices) AddGame(data *models.AddGameDTO) (*models.Game, *types.Error) {
	id, rerr := idgen.NextNumericString(30)
	if rerr != nil {
		return nil, types.NewBadRequestError("failed to make the sid #4702")
	}

	data.Id = id
	data.NumMainNumbers = models.FromStringToNumberMainNumbers(data.GameTypeInt)
	data.MainNumberRange = models.FromStringToMainNumberRange(data.GameTypeInt)
	data.NumBonusNumbers = models.FromStringToNumberBonusNumbers(data.GameTypeInt)
	data.BonusNumberRange = models.FromStringToBonusNumberRange(data.GameTypeInt)
	data.GameTypeString = models.GameTypeToString(data.GameTypeInt)
	err := c.repository.AddGame(data)
	if err != nil {
		return nil, err
	}

	return &models.Game{
		Id:               id,
		Name:             data.Name,
		GameType:         models.GameTypeToString(data.GameTypeInt),
		NumMainNumbers:   data.NumMainNumbers,
		NumBonusNumbers:  data.NumBonusNumbers,
		MainNumberRange:  data.MainNumberRange,
		BonusNumberRange: data.BonusNumberRange,
		StartTime:        data.StartTime,
		EndTime:          data.EndTime,
		CreatorId:        data.CreatorId,
		Result:           nil,
		CreatedAt:        time.Now(),
	}, nil
}

func (c *gameServices) GetNextGamesByGameType(gameType int32, limit int32) ([]models.Game, *types.Error) {
	gameTypeS := models.GameTypeToString(gameType)
	return c.repository.GetNextGamesByGameType(gameTypeS, limit)
}

func (c *gameServices) GetAllNextGames() ([]models.Game, *types.Error) {
	return c.repository.GetAllNextGames()
}

func (c *gameServices) DeleteGameByGameId(gameId string) *types.Error {
	return c.repository.DeleteGameByGameId(gameId)
}

func (c *gameServices) CheckGameExistsById(gameId string) (bool, *types.Error) {
	return c.repository.CheckGameExistsById(gameId)
}

func (c *gameServices) GetGamesByCreatorId(creatorId int32, data *models.Pagination) (*models.Games, *types.Error) {
	res, err := c.repository.GetGamesByCreatorId(creatorId, data)
	if err != nil {
		return nil, err
	}
	var total *int32
	if data.Total {
		count, err := c.repository.GetGamesCountByCreatorId(creatorId)
		if err != nil {
			return nil, err
		}

		total = &count
	}

	return &models.Games{Games: res, Total: total}, nil
}

func (c *gameServices) GetAllPreviousGames(data *models.Pagination) (*models.Games, *types.Error) {
	res, err := c.repository.GetAllPreviousGames(data.Offset, data.Limit)
	if err != nil {
		return nil, err
	}
	var total *int32
	if data.Total {
		count, err := c.repository.GetAllPreviousGamesCount()
		if err != nil {
			return nil, err
		}

		total = &count
	}

	return &models.Games{Games: res, Total: total}, nil
}

func (c *gameServices) GetAllGames(data *models.Pagination) (*models.Games, *types.Error) {
	res, err := c.repository.GetAllGames(data.Offset, data.Limit)
	if err != nil {
		return nil, err
	}
	var total *int32
	if data.Total {
		count, err := c.repository.GetAllGamesCount()
		if err != nil {
			return nil, err
		}

		total = &count
	}

	return &models.Games{Games: res, Total: total}, nil
}

func (c *gameServices) AddResultByGameId(gameId string, result string) ([]models.DivisionResult, *types.Error) {
	err := c.repository.AddResultByGameId(gameId, result)
	if err != nil {
		return nil, err
	}

	return c.repository.FindUsersByResultAndGameId(gameId)
}
