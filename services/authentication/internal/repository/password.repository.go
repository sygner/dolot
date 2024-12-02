package repository

import (
	"database/sql"
	"dolott_authentication/internal/models"
	"dolott_authentication/internal/types"
	"fmt"
)

func (c *authenticationRepository) GetPasswordByUserId(user_id int32) (*models.UserPassword, *types.Error) {
	query := `SELECT password FROM "passwords" WHERE user_id = $1`
	var password string
	err := c.db.QueryRow(query, user_id).Scan(&password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, types.NewNotFoundError("user not found, error code #1012")
		}
		fmt.Println(err)
		return nil, types.NewInternalError("internal issue, error code #1013")
	}

	return &models.UserPassword{UserId: user_id, Password: password}, nil
}

func (c *authenticationRepository) AddPassword(data *models.UserPassword) *types.Error {
	query := `INSERT INTO "passwords" (user_id, password) VALUES ($1,$2)`
	_, err := c.db.Exec(query, data.UserId, data.Password)
	if err != nil {
		return types.NewInternalError("internal issue, error code #1014")
	}
	return nil
}

func (c *authenticationRepository) UpdatePassword(userId int32, newPassword string) *types.Error {
	query := `UPDATE passwords SET password = $1 WHERE user_id = $2`
	_, err := c.db.Exec(query, newPassword, userId)
	if err != nil {
		fmt.Println(err)
		return types.NewInternalError("internal issue, error code #1041")
	}
	return nil
}
