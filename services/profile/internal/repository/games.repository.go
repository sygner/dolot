package repository

import "dolott_profile/internal/types"

func (c *profileRepository) ChangeUserGamesQuantity(userId int32, increment bool) *types.Error {
	var query string

	if increment {
		query = `UPDATE profiles SET games_quantity = games_quantity + 1 WHERE user_id = $1`
	} else {
		query = `UPDATE profiles SET games_quantity = games_quantity - 1 WHERE user_id = $1`
	}

	_, err := c.db.Exec(query, userId)
	if err != nil {
		return types.NewInternalError("failed to add game quantity #3010")
	}
	return nil
}

// func (c *profileRepository) DecrementUserGamesQuantity(userId int32) *types.Error {
// 	query := `UPDATE profiles WHERE user_id = $1 SET games_quantity = games_quantity - 1`
// 	_, err := c.db.Exec(query, userId)
// 	if err != nil {
// 		return types.NewInternalError("failed to add ranks #3011")
// 	}
// 	return nil
// }

func (c *profileRepository) ChangeUserWonGames(userId int32, Increment bool) *types.Error {
	var query string
	if Increment {
		query = `UPDATE profiles SET won_games = won_games + 1 WHERE user_id = $1`
	} else {
		query = `UPDATE profilesSET won_games = won_games - 1 WHERE user_id = $1 `
	}

	_, err := c.db.Exec(query, userId)
	if err != nil {
		return types.NewInternalError("failed to add won game #3012")
	}
	return nil
}

// func (c *profileRepository) DecrementUserWonGames(userId int32) *types.Error {
// 	query := `UPDATE profiles WHERE user_id = $1 SET won_games = won_games - 1`
// 	_, err := c.db.Exec(query, userId)
// 	if err != nil {
// 		return types.NewInternalError("failed to add ranks #3013")
// 	}
// 	return nil
// }

func (c *profileRepository) ChangeUserLostGames(userId int32, increment bool) *types.Error {
	var query string
	if increment {
		query = `UPDATE profiles SET lost_games = lost_games + 1 WHERE user_id = $1 `
	} else {
		query = `UPDATE profiles SET lost_games = lost_games - 1 WHERE user_id = $1 `
	}

	_, err := c.db.Exec(query, userId)
	if err != nil {
		return types.NewInternalError("failed to add lost game #3014")
	}
	return nil
}

// func (c *profileRepository) DecrementUserLostGames(userId int32) *types.Error {
// 	query := `UPDATE profiles WHERE user_id = $1 SET lost_games = lost_games - 1`
// 	_, err := c.db.Exec(query, userId)
// 	if err != nil {
// 		return types.NewInternalError("failed to add ranks #3015")
// 	}
// 	return nil
// }
