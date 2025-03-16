package repository

import (
	"database/sql"
	"dolott_game/internal/models"
	"dolott_game/internal/types"
	"time"
)

type (
	GameRepository interface {
		GetGameByGameId(string) (*models.Game, *types.Error)
		AddGame(*models.AddGameDTO) *types.Error
		GetNextGamesByGameType(string, int32) ([]models.Game, *types.Error)
		DeleteGameByGameId(string) *types.Error
		GetGamesByCreatorId(int32, *models.Pagination) ([]models.Game, *types.Error)
		GetGamesCountByCreatorId(int32) (int32, *types.Error)
		CheckGameExistsById(string) (bool, *types.Error)
		AddResultByGameId(string, string) *types.Error
		CheckGameExistsByIdAndEndTime(string) (bool, *types.Error)
		FindUsersByResultAndGameId(string) ([]models.DivisionResult, *types.Error)
		GetUsersByResultAndGameId(string) ([]models.DivisionResult, *types.Error)
		GetUserByResultAndGameId(int32, string) ([]models.DivisionResult, *types.Error)
		GetAllNextGames() ([]models.Game, *types.Error)
		GetAllPreviousGames(int32, int32) ([]models.Game, *types.Error)
		GetAllPreviousGamesCount() (int32, *types.Error)
		GetAllGames(int32, int32) ([]models.Game, *types.Error)
		GetAllGamesCount() (int32, *types.Error)
		GetGameTypes() ([]models.GameTypeDetail, *types.Error)
		UpdateGameTypeDetail(int32, *string, int32, int32, bool) *types.Error
		GetAllUserPreviousGames(int32, int32, int32) ([]models.Game, *types.Error)
		GetAllUserPreviousGamesCount(int32) (int32, *types.Error)
		GetAllUserPreviousGamesByGameType(int32, string, int32, int32) ([]models.Game, *types.Error)
		GetAllUserPreviousGamesCountByGameType(int32, string) (int32, *types.Error)
		// FindUsersByResultAndGameId(string) ([]models.UserChoiceResult, []models.UserChoiceResult, *types.Error)
		AddUserChoice(*models.AddUserChoiceDTO) *types.Error
		GetAllUserGames(int32) ([]string, *types.Error)
		UpdateGamePrizeByGameId(string, *uint32, bool) *types.Error

		GetUserChoicesByUserId(int32, *models.Pagination) ([]models.UserChoice, *types.Error)
		GetUserChoicesByUserIdAndTimeRange(int32, time.Time, time.Time) ([]models.UserChoice, *types.Error)
		GetUserChoicesCountByUserId(int32) (int32, *types.Error)
		GetUserChoicesByGameIdAndPagination(string, *models.Pagination) ([]models.UserChoice, *types.Error)
		GetUserChoicesCountByGameId(string) (int32, *types.Error)
		GetUserChoicesEachCountByGameId(string) (int32, *types.Error)

		GetWinnersByGameId(string) (*models.Winners, *types.Error)
		UpdateTotalPaidUsers(string, string) *types.Error
		GetLastWinnersByGameType(int32) (*models.Winners, *types.Error)
	}
	gameRepository struct {
		db *sql.DB
	}
)

func NewGameRepository(db *sql.DB) GameRepository {
	return &gameRepository{
		db: db,
	}
}
