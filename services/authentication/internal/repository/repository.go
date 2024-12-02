package repository

import (
	"database/sql"
	"dolott_authentication/internal/models"
	"dolott_authentication/internal/types"
	"time"
)

type (
	AuthenticationRepository interface {
		UserExistsByEmail(string) (bool, *types.Error)
		UserExistsByPhoneNumber(string) (bool, *types.Error)
		UserExistsByUserId(int32) (bool, *types.Error)
		UserExistsByAccountUsername(string) (bool, *types.Error)
		GetUserByUserId(int32) (*models.User, *types.Error)
		GetUserByEmail(string) (*models.User, *types.Error)
		GetUserByAccountUsername(string) (*models.User, *types.Error)
		GetRoleByUserId(int32) (string, *types.Error)
		AddUser(*models.UserDTO) (*models.User, *types.Error)

		GetPasswordByUserId(int32) (*models.UserPassword, *types.Error)
		AddPassword(*models.UserPassword) *types.Error
		UpdatePassword(int32, string) *types.Error

		AddToken(*models.Token) *types.Error
		DeleteUserTokens(int32) *types.Error
		DeleteTokenByAccessToken(string) *types.Error
		GetTokenByAccessToken(string) (*models.Token, *types.Error)
		GetTokenByAccessTokenAndAgent(string, string) (*models.Token, *types.Error)
		GetTokenByAccessTokenAndAgentAndUserId(string, string, int32) (*models.Token, *types.Error)
		GetTokensByUserId(*models.Pagination, int32) ([]models.Token, *types.Error)
		GetTokens(*models.Pagination) ([]models.Token, *types.Error)
		DeleteTokenBySessionIdAdnUserId(int32, int32) *types.Error
		GetTokenByAccessTokenAndRefreshToken(string, string) (*models.Token, *types.Error)
		GetTokensCount() (int32, *types.Error)
		GetTokensByUserIdCount(int32) (int32, *types.Error)

		AddLoginHistory(int32) *types.Error
		GetLoginHistoryByUserId(*models.Pagination, int32) ([]time.Time, *types.Error)
		GetLoginHistoryCountByUserId(int32) (int32, *types.Error)

		GetUsers(*models.Pagination) ([]models.User, *types.Error)
		GetUserCount() (int32, *types.Error)
		ChangeUserStatus(int32, string) *types.Error
	}
	authenticationRepository struct {
		db *sql.DB
	}
)

func NewAuthenticationRepository(db *sql.DB) AuthenticationRepository {
	return &authenticationRepository{
		db: db,
	}
}
