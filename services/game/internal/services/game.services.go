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
		GetGameTypes() ([]models.GameTypeDetail, *types.Error)
		UpdateGameTypeDetail(int32, *string, int32, int32, bool) *types.Error
		GetAllUserPreviousGames(int32, *models.Pagination) (*models.Games, *types.Error)
		GetAllUserChoiceDivisionsByGameId(int32, string) ([]models.DivisionResult, *types.Error)
		GetAllUsersChoiceDivisionsByGameId(string) (*models.Winners, *types.Error)
		UpdateGamePrizeByGameId(string, *uint32, bool) *types.Error
		GetAllUserPreviousGamesByGameType(int32, string, *models.Pagination) (*models.Games, *types.Error)
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

func (c *gameServices) ComputePrizeForGame(game *models.Game) (*models.Game, *types.Error) {
	if game == nil {
		return nil, types.NewInternalError("invalid game data, error code #4704")
	}
	// If auto_compute_prize is true, calculate prize from user choices
	if game.AutoCompute {
		userChoicesCount, err := c.repository.GetUserChoicesEachCountByGameId(game.Id)
		if err != nil {
			return nil, err
		}
		*game.Prize = uint32(userChoicesCount) // Set the computed prize
	}

	return game, nil
}

func (c *gameServices) GetGameByGameId(gameId string) (*models.Game, *types.Error) {
	res, err := c.repository.GetGameByGameId(gameId)
	if err != nil {
		return nil, err
	}
	return c.ComputePrizeForGame(res)
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
		Prize:            data.Prize,
		AutoCompute:      data.AutoCompute,
		CreatedAt:        time.Now(),
	}, nil
}

func (c *gameServices) GetNextGamesByGameType(gameType int32, limit int32) ([]models.Game, *types.Error) {
	gameTypeS := models.GameTypeToString(gameType)
	res, err := c.repository.GetNextGamesByGameType(gameTypeS, limit)
	if err != nil {
		return nil, err
	}
	for i, r := range res {
		game, err := c.ComputePrizeForGame(&r)
		if err != nil {
			return nil, err
		}
		res[i].Prize = game.Prize
	}
	return res, nil
}

func (c *gameServices) GetAllNextGames() ([]models.Game, *types.Error) {
	res, err := c.repository.GetAllNextGames()
	if err != nil {
		return nil, err
	}
	for i, r := range res {
		game, err := c.ComputePrizeForGame(&r)
		if err != nil {
			return nil, err
		}
		res[i].Prize = game.Prize
	}
	return res, nil
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
	for i, r := range res {
		game, err := c.ComputePrizeForGame(&r)
		if err != nil {
			return nil, err
		}
		res[i].Prize = game.Prize
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
	for i, r := range res {
		game, err := c.ComputePrizeForGame(&r)
		if err != nil {
			return nil, err
		}
		res[i].Prize = game.Prize
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
	for i, r := range res {
		game, err := c.ComputePrizeForGame(&r)
		if err != nil {
			return nil, err
		}
		res[i].Prize = game.Prize
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

func (c *gameServices) GetGameTypes() ([]models.GameTypeDetail, *types.Error) {
	gameTypes, err := c.repository.GetGameTypes()
	if err != nil {
		return nil, err
	}

	for i, gameType := range gameTypes {
		if gameType.AutoCompute {
			// Get upcoming games of this game type
			upcomingGames, err := c.repository.GetNextGamesByGameType(gameType.TypeName, 1000)
			if err != nil {
				return nil, err
			}

			if len(upcomingGames) != 0 {
				continue
			}

			// Calculate total user choices across all upcoming games
			var totalUserChoices int32
			for _, game := range upcomingGames {
				userChoiceCount, err := c.repository.GetUserChoicesCountByGameId(game.Id)
				if err != nil {
					return nil, err
				}
				totalUserChoices += userChoiceCount
			}

			// Get the last game's winners of this game type
			lastWin, err := c.repository.GetLastWinnersByGameType(gameType.Id - 1)
			if err != nil {
				if err.Code != 404 {
					return nil, err
				}
			}

			lastJackpot := int32(0)
			if lastWin != nil && lastWin.Divisions != nil {
				division := models.SearchDivision(lastWin.Divisions, "Division 1")

				// If Division 1 has no winners, roll over the jackpot
				if division == nil {
					lastJackpot = int32(len(lastWin.Divisions)) // Roll over the last jackpot amount
				}
			}

			var averageUserChoices int32 = 0

			if len(upcomingGames) > 0 {
				averageUserChoices = (totalUserChoices / int32(len(upcomingGames))) + lastJackpot
			} else {
				averageUserChoices = totalUserChoices + lastJackpot // If no upcoming games, just use last jackpot
			}

			// Update the game type prize and token burn
			gameTypes[i].PrizeReward = averageUserChoices
			gameTypes[i].TokenBurn = totalUserChoices
		}
	}

	return gameTypes, nil
}

func (c *gameServices) UpdateGameTypeDetail(gameType int32, dayName *string, prizeReward int32, tokenBurn int32, autoCompute bool) *types.Error {
	return c.repository.UpdateGameTypeDetail(gameType, dayName, prizeReward, tokenBurn, autoCompute)
}

func (c *gameServices) GetAllUserPreviousGames(userId int32, data *models.Pagination) (*models.Games, *types.Error) {
	res, err := c.repository.GetAllUserPreviousGames(userId, data.Offset, data.Limit)
	if err != nil {
		return nil, err
	}
	var total *int32
	if data.Total {
		count, err := c.repository.GetAllUserPreviousGamesCount(userId)
		if err != nil {
			return nil, err
		}

		total = &count
	}
	for i, r := range res {
		game, err := c.ComputePrizeForGame(&r)
		if err != nil {
			return nil, err
		}
		res[i].Prize = game.Prize
	}
	return &models.Games{Games: res, Total: total}, nil
}

func (c *gameServices) GetAllUserPreviousGamesByGameType(userId int32, gameType string, data *models.Pagination) (*models.Games, *types.Error) {
	res, err := c.repository.GetAllUserPreviousGamesByGameType(userId, gameType, data.Offset, data.Limit)
	if err != nil {
		return nil, err
	}
	var total *int32
	if data.Total {
		count, err := c.repository.GetAllUserPreviousGamesCountByGameType(userId, gameType)
		if err != nil {
			return nil, err
		}

		total = &count
	}
	for i, r := range res {
		game, err := c.ComputePrizeForGame(&r)
		if err != nil {
			return nil, err
		}
		res[i].Prize = game.Prize
	}
	return &models.Games{Games: res, Total: total}, nil
}

func (c *gameServices) GetAllUserChoiceDivisionsByGameId(userId int32, gameId string) ([]models.DivisionResult, *types.Error) {
	return c.repository.GetUserByResultAndGameId(userId, gameId)
	// res, err := c.repository.GetWinnersByGameId(gameId)
	// if err != nil {
	// 	return nil, err
	// }
	// resultDivisions := make([]models.DivisionResult, 0)
	// for _, result := range res {
	// 	for _, choices := range result.UserChoices {
	// 		if choices.UserId == userId {
	// 			resultDivisions = append(resultDivisions, result)
	// 		}

	// 	}

	// }
	// return res, nil
}

func (c *gameServices) GetAllUsersChoiceDivisionsByGameId(gameId string) (*models.Winners, *types.Error) {
	return c.repository.GetWinnersByGameId(gameId)
}

func (c *gameServices) UpdateGamePrizeByGameId(gameId string, prize *uint32, autoCompute bool) *types.Error {
	return c.repository.UpdateGamePrizeByGameId(gameId, prize, autoCompute)
}
