package repository

import (
	"dolott_profile/internal/types"
	"fmt"
)

func (c *profileRepository) ChangeUserScore(userId int32, score float32, increment bool) *types.Error {
	var query string
	if increment {
		query = `UPDATE profiles SET score = score + $1 WHERE user_id = $2`
	} else {
		query = `UPDATE profiles SET score = score - $1 WHERE user_id = $2`
	}
	fmt.Println(query)
	_, err := c.db.Exec(query, score, userId)
	if err != nil {
		fmt.Println(err)
		return types.NewInternalError("failed to increment ranks #3008")
	}
	return nil
}

// func (c *profileRepository) DecrementUserScore(userId int32, score float32) *types.Error {
// 	query := `UPDATE profiles WHERE user_id = $1 SET rank = score - $2`

// 	_, err := c.db.Exec(query, userId, score)
// 	if err != nil {
// 		return types.NewInternalError("failed to decrement ranks #3009")
// 	}
// 	return nil
// }
