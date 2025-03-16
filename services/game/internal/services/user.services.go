package services

import (
	"dolott_game/internal/models"
	"dolott_game/internal/repository"
	"dolott_game/internal/types"
	"safir/libs/idgen"
	"time"
)

type (
	UserServices interface {
		AddUserChoice(*models.AddUserChoiceDTO, bool) (*models.UserChoice, *types.Error)
		GetUserChoicesByUserId(int32, *models.Pagination) (*models.UserChoices, *types.Error)
		GetUserChoicesByUserIdAndTimeRange(int32, time.Time, time.Time) ([]models.UserChoice, *types.Error)
		GetUserChoicesByGameIdAndPagination(string, *models.Pagination) (*models.UserChoices, *types.Error)
		GetAllUserGames(int32) ([]string, *types.Error)
	}
	userServices struct {
		repository repository.GameRepository
	}
)

func NewUserServices(repository repository.GameRepository) UserServices {
	return &userServices{
		repository: repository,
	}
}
func (c *userServices) AddUserChoice(data *models.AddUserChoiceDTO, shouldReturn bool) (*models.UserChoice, *types.Error) {
	exists, err := c.repository.CheckGameExistsByIdAndEndTime(data.GameId)
	if err != nil {
		return nil, err
	}

	if !exists {
		return nil, types.NewBadRequestError("game is expired or not found #4701")
	}

	id, rerr := idgen.NextNumericString(64)
	if rerr != nil {
		return nil, types.NewBadRequestError("failed to make the sid #4703")
	}
	data.Id = id
	err = c.repository.AddUserChoice(data)
	if err != nil {
		return nil, err
	}
	if shouldReturn {
		return &models.UserChoice{
			Id:                 id,
			UserId:             data.UserId,
			GameId:             data.GameId,
			ChosenMainNumbers:  data.ChosenMainNumbers,
			ChosenBonusNumbers: data.ChosenBonusNumbers,
			CreatedAt:          time.Now(),
		}, nil
	} else {
		return nil, nil
	}

}

func (c *userServices) GetUserChoicesByUserId(userId int32, pagination *models.Pagination) (*models.UserChoices, *types.Error) {
	res, err := c.repository.GetUserChoicesByUserId(userId, pagination)
	if err != nil {
		return nil, err
	}
	var total *int32
	if pagination.Total {
		totaly, err := c.repository.GetUserChoicesCountByUserId(userId)
		if err != nil {
			return nil, err
		}
		total = &totaly
	}

	return &models.UserChoices{UserChoices: res, Total: total}, nil

}

func (c *userServices) GetUserChoicesByUserIdAndTimeRange(userId int32, startTime time.Time, endTime time.Time) ([]models.UserChoice, *types.Error) {
	return c.repository.GetUserChoicesByUserIdAndTimeRange(userId, startTime, endTime)
}

func (c *userServices) GetUserChoicesByGameIdAndPagination(gameId string, pagination *models.Pagination) (*models.UserChoices, *types.Error) {
	res, err := c.repository.GetUserChoicesByGameIdAndPagination(gameId, pagination)
	if err != nil {
		return nil, err
	}
	var total *int32
	if pagination.Total {
		totaly, err := c.repository.GetUserChoicesCountByGameId(gameId)
		if err != nil {
			return nil, err
		}
		total = &totaly
	}

	return &models.UserChoices{UserChoices: res, Total: total}, nil
}

func (c *userServices) GetAllUserGames(userId int32) ([]string, *types.Error) {
	return c.repository.GetAllUserGames(userId)
}
