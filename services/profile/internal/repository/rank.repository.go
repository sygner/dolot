package repository

import "dolott_profile/internal/types"

func (c *profileRepository) ChangeAllRanks(increment bool) *types.Error {
	var query string
	if increment {
		query = `UPDATE profiles SET rank = rank + 1`
	} else {
		query = `UPDATE profiles SET rank = rank - 1`
	}

	_, err := c.db.Exec(query)
	if err != nil {
		return types.NewInternalError("failed to increment ranks #3006")
	}
	return nil
}

func (c *profileRepository) AdjustUserRank(userId int32, rankAmount int32, increment bool) *types.Error {
	query := `
	WITH target_user AS (
		SELECT user_id, rank 
		FROM profiles 
		WHERE user_id = $1
	), updated_user AS (
		SELECT user_id, 
		       CASE WHEN $3 THEN rank + $2 ELSE rank - $2 END AS new_rank 
		FROM target_user
	), affected_users AS (
		SELECT user_id, rank 
		FROM profiles 
		WHERE 
			(CASE 
				WHEN $3 THEN rank > (SELECT rank FROM target_user) AND rank <= (SELECT new_rank FROM updated_user)
				ELSE rank < (SELECT rank FROM target_user) AND rank >= (SELECT new_rank FROM updated_user)
			END)
		  AND user_id != (SELECT user_id FROM target_user)
	)
	UPDATE profiles 
	SET rank = CASE 
		WHEN user_id = (SELECT user_id FROM target_user) THEN 
			CASE WHEN $3 THEN rank + $2 ELSE rank - $2 END
		ELSE rank + (CASE WHEN $3 THEN -1 ELSE 1 END)
		END
	WHERE user_id = (SELECT user_id FROM target_user)
	   OR user_id IN (SELECT user_id FROM affected_users);
	`

	_, err := c.db.Exec(query, userId, rankAmount, increment)
	if err != nil {
		return types.NewInternalError("failed to adjust ranks #3501")
	}
	return nil
}
