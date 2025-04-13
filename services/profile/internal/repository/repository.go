package repository

import (
	"database/sql"
	"dolott_profile/internal/models"
	"dolott_profile/internal/types"
)

type (
	ProfileRepository interface {
		GetProfileByUsername(string) (*models.Profile, *types.Error)
		GetProfileBySid(string) (*models.Profile, *types.Error)
		GetProfileByUserId(int32) (*models.Profile, *types.Error)
		AddProfile(*models.AddProfileDTO) *types.Error
		GetHighestRank() (int32, *types.Error)
		ChangeAllRanks(bool) *types.Error

		// IncrementAllRanks() *types.Error
		// DecrementAllRanks() *types.Error

		ChangeUserScore(int32, float32, bool) *types.Error

		ChangeUserGamesQuantity(int32, bool) *types.Error

		ChangeUserWonGames(int32, bool) *types.Error

		ChangeUserLostGames(int32, bool) *types.Error

		AdjustUserRank(int32, int32, bool) *types.Error

		GetAllUserRanking(int32) (*models.Ranking, *types.Error)

		SearchUsername(string) ([]models.Profile, *types.Error)

		GetUserLeaderBoard(int32) ([]models.Profile, *types.Error)

		ChangeUserImpression(int32, int32, bool) *types.Error

		CheckUsernameExists(string) *types.Error

		UpdateProfile(int32, string) *types.Error

		ChangeImpressionAndDCoin(int32, int32, int32) *types.Error
		// IncrementUserScore(int32, float32) *types.Error
		// DecrementUserScore(int32, float32) *types.Error

		// IncrementUserGamesQuantity(int32) *types.Error
		// DecrementUserGamesQuantity(int32) *types.Error

		// IncrementUserWonGames(int32) *types.Error
		// DecrementUserWonGames(int32) *types.Error

		// IncrementUserLostGames(int32) *types.Error
		// DecrementUserLostGames(int32) *types.Error
	}
	profileRepository struct {
		db *sql.DB
	}
)

func NewProfileRepository(db *sql.DB) ProfileRepository {
	return &profileRepository{
		db: db,
	}
}
