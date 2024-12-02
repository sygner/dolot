package services

import (
	"dolott_game/internal/models"
	"dolott_game/internal/repository"
	"dolott_game/internal/types"
)

type (
	WinnerServices interface {
		GetWinnersByGameId(string) ([]models.DivisionResult, *types.Error)
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

func (c *winnerService) GetWinnersByGameId(gameId string) ([]models.DivisionResult, *types.Error) {
	return c.repository.GetWinnersByGameId(gameId)
}
