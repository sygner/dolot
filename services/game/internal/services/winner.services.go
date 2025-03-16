package services

import (
	"dolott_game/internal/models"
	"dolott_game/internal/repository"
	"dolott_game/internal/types"
)

type (
	WinnerServices interface {
		GetWinnersByGameId(string) (*models.Winners, *types.Error)
		UpdateTotalPaidUsers(string, string) *types.Error
		GetLastWinnersByGameType(int32) (*models.Winners, *types.Error)
	}
	winnerService struct {
		repository repository.GameRepository
	}
)

func NewWinnerServices(repository repository.GameRepository) WinnerServices {
	return &winnerService{
		repository: repository,
	}
}

func (c *winnerService) GetWinnersByGameId(gameId string) (*models.Winners, *types.Error) {
	return c.repository.GetWinnersByGameId(gameId)
}

func (c *winnerService) UpdateTotalPaidUsers(gameId string, totalPaid string) *types.Error {
	return c.repository.UpdateTotalPaidUsers(gameId, totalPaid)
}

func (c *winnerService) GetLastWinnersByGameType(gameType int32) (*models.Winners, *types.Error) {
	return c.repository.GetLastWinnersByGameType(gameType)
}
