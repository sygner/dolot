package services

import (
	"dolott_authentication/internal/models"
	"dolott_authentication/internal/repository"
	"dolott_authentication/internal/types"
)

type (
	LoginHistoryService interface {
		GetLoginHistoryByUserId(*models.Pagination, int32) (*models.LoginHistories, *types.Error)
	}
	loginHistoryService struct {
		repository repository.AuthenticationRepository
	}
)

func NewLoginHistoryService(repository repository.AuthenticationRepository) LoginHistoryService {
	return &loginHistoryService{
		repository: repository,
	}
}

func (c *loginHistoryService) GetLoginHistoryByUserId(data *models.Pagination, userId int32) (*models.LoginHistories, *types.Error) {
	res, err := c.repository.GetLoginHistoryByUserId(data, userId)
	if err != nil {
		return nil, err
	}

	var total *int32
	if data.Total {
		count, err := c.repository.GetLoginHistoryCountByUserId(userId)
		if err != nil {
			return nil, err
		}

		total = &count
	}

	return &models.LoginHistories{LoginHistories: res, Total: total}, nil
}
