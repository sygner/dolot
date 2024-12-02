package repository

import (
	"dolott_authentication/internal/models"
	"dolott_authentication/internal/types"
	"time"
)

func (c *authenticationRepository) AddLoginHistory(userId int32) *types.Error {
	query := `INSERT INTO "login_history" (user_id, created_at) VALUES ($1,NOW())`
	_, err := c.db.Exec(query, userId)
	if err != nil {
		return types.NewInternalError("internal issue, error code #1016")
	}

	return nil
}

func (c *authenticationRepository) GetLoginHistoryByUserId(data *models.Pagination, userId int32) ([]time.Time, *types.Error) {
	query := `SELECT created_at FROM "login_history" WHERE user_id = $1 OFFSET $2 LIMIT $3`

	rows, err := c.db.Query(query, userId, data.Offset, data.Limit)
	if err != nil {
		return nil, types.NewInternalError("internal issue, error code #1038")
	}

	defer rows.Close()
	loggedList := make([]time.Time, 0)
	for rows.Next() {
		var logged time.Time
		if err := rows.Scan(
			&logged); err != nil {
			return nil, types.NewInternalError("internal issue, error code #1039")
		}
		loggedList = append(loggedList, logged)
	}
	return loggedList, nil
}

func (c *authenticationRepository) GetLoginHistoryCountByUserId(userId int32) (int32, *types.Error) {
	query := `SELECT count(*) FROM "login_history" WHERE user_id = $1`
	var totalCount int32
	err := c.db.QueryRow(query, userId).Scan(&totalCount)
	if err != nil {
		return 0, types.NewInternalError("internal issue, error code #1040")
	}
	return totalCount, nil
}
