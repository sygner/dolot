package repository

import (
	"dolott_profile/internal/types"
)

func (c *profileRepository) AddRankLog(userId int32) *types.Error {
	query := `
		INSERT INTO rank_logs (user_id, rank, created_at)
		SELECT user_id, rank, NOW() FROM profiles WHERE user_id = $1
	`
	result, err := c.db.Exec(query, userId)
	if err != nil {
		return types.NewInternalError("failed to add rank log #3021")
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return types.NewNotFoundError("user not found in profiles #3022")
	}

	return nil
}
