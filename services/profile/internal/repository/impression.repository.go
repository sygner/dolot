package repository

import (
	"dolott_profile/internal/types"
)

func (c *profileRepository) ChangeUserImpression(userId int32, impression int32, increment bool) *types.Error {
	var query string
	if increment {
		query = `UPDATE profiles SET impression = impression + $1 WHERE user_id = $2`
	} else {
		query = `UPDATE profiles SET impression = impression - $1 WHERE user_id = $2`
	}
	_, err := c.db.Exec(query, impression, userId)
	if err != nil {
		return types.NewInternalError("failed to increment ranks #3008")
	}
	return nil
}
