package repository

import (
	"dolott_profile/internal/models"
	"dolott_profile/internal/types"
	"fmt"
)

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
           CASE 
               WHEN $3 THEN rank + $2 
               ELSE GREATEST(rank - $2, 1) -- Ensures rank never goes below 1
           END AS new_rank 
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
        CASE 
            WHEN $3 THEN rank + $2 
            ELSE GREATEST(rank - $2, 1) -- Ensures rank never goes below 1
        END
    ELSE rank + (CASE WHEN $3 THEN -1 ELSE 1 END)
    END
WHERE user_id = (SELECT user_id FROM target_user)
   OR user_id IN (SELECT user_id FROM affected_users);

	`

	_, err := c.db.Exec(query, userId, rankAmount, increment)
	if err != nil {
		fmt.Println(err)
		return types.NewInternalError("failed to adjust ranks #3501")
	}

	rerr := c.AddRankLog(userId)
	if rerr != nil {
		return rerr
	}

	return nil
}

func (c *profileRepository) GetHighestRank() (int32, *types.Error) {
	var result int32
	query := "SELECT COALESCE(MAX(rank), 0) FROM profiles;"

	err := c.db.QueryRow(query).Scan(&result)
	if err != nil {
		return result, types.NewInternalError("failed to get the highest rank #3006")
	}

	return result, nil
}

func (c *profileRepository) GetAllUserRanking(userId int32) (*models.Ranking, *types.Error) {
	query := `
	WITH latest_ranks AS (
    SELECT rank FROM rank_logs 
    WHERE user_id = $1
    ORDER BY created_at DESC
    LIMIT 1
),
season_ranks AS (
    SELECT rank FROM rank_logs 
    WHERE user_id = $1 
    AND created_at >= date_trunc('month', NOW()) - INTERVAL '3 months'
    ORDER BY created_at DESC
    LIMIT 1
),
month_ranks AS (
    SELECT rank FROM rank_logs 
    WHERE user_id = $1 
    AND created_at >= date_trunc('month', NOW()) 
    ORDER BY created_at DESC
    LIMIT 1
),
profile_rank AS (
    SELECT rank FROM profiles WHERE user_id = $1
),
season_user_changes AS (
    SELECT COUNT(DISTINCT user_id) AS user_count
    FROM rank_logs 
    WHERE created_at >= date_trunc('month', NOW()) - INTERVAL '3 months'
    AND rank IS NOT NULL
),
month_user_changes AS (
    SELECT COUNT(DISTINCT user_id) AS user_count
    FROM rank_logs 
    WHERE created_at >= date_trunc('month', NOW()) 
    AND rank IS NOT NULL
),
all_user_ranks AS (
    SELECT COUNT(DISTINCT user_id) AS user_count
    FROM profiles  
)
SELECT 
    COALESCE((SELECT rank FROM latest_ranks), 1) AS total_rank,
    COALESCE((SELECT rank FROM profile_rank), 1) AS individual_rank,
    COALESCE((SELECT rank FROM season_ranks), 1) AS season_rank,
    COALESCE((SELECT rank FROM month_ranks), 1) AS month_rank,
    COALESCE((SELECT user_count FROM season_user_changes), 0) AS season_rank_changes_count,
	COALESCE((SELECT user_count FROM month_user_changes), 0) AS month_rank_changes_count,
	COALESCE((SELECT user_count FROM all_user_ranks), 0) AS all_rank_changes_count
	;

	`
	ranking := &models.Ranking{} // Initialize the struct

	err := c.db.QueryRow(query, userId).Scan(
		&ranking.TotalRanking,
		&ranking.IndividualRanking,
		&ranking.SeasonRanking,
		&ranking.MonthRanking,
		&ranking.SeasonRankChangesCount,
		&ranking.MonthRankChangesCount,
		&ranking.AllRankChangesCount,
	)
	if err != nil {
		return nil, types.NewInternalError("failed to get the user ranking #3023")
	}
	return ranking, nil
}

func (c *profileRepository) GetUserLeaderBoard(userId int32) ([]models.Profile, *types.Error) {
	query := `
WITH ranked_users AS (
    SELECT user_id, sid, username, score, impression, d_credit, rank, games_quantity, won_games, lost_games, created_at
    FROM profiles
    ORDER BY rank ASC
),
user_rank_details AS (
    SELECT user_id, sid, username, score, impression, d_credit, rank, games_quantity, won_games, lost_games, created_at
    FROM ranked_users
    WHERE user_id = $1  -- Current user
),
highest_ranks AS (
    SELECT user_id, sid, username, score, impression, d_credit, rank, games_quantity, won_games, lost_games, created_at
    FROM ranked_users
    ORDER BY rank ASC
    LIMIT 3
),
above_user_ranks AS (
    SELECT user_id, sid, username, score, impression, d_credit, rank, games_quantity, won_games, lost_games, created_at
    FROM ranked_users
    WHERE rank < (SELECT rank FROM user_rank_details)
    ORDER BY rank ASC
    LIMIT 3
),
below_user_ranks AS (
    SELECT user_id, sid, username, score, impression, d_credit, rank, games_quantity, won_games, lost_games, created_at
    FROM ranked_users
    WHERE rank > (SELECT rank FROM user_rank_details)
    ORDER BY rank ASC
    LIMIT 3
)
-- Use UNION (not UNION ALL) to avoid duplicates, then apply DISTINCT
SELECT DISTINCT user_id, sid, username, score, impression, d_credit, rank, games_quantity, won_games, lost_games, created_at
FROM (
    -- Current User Rank
    SELECT user_id, sid, username, score, impression, d_credit, rank, games_quantity, won_games, lost_games, created_at
    FROM user_rank_details
    UNION
    -- Top 3 Ranks excluding the current user
    SELECT user_id, sid, username, score, impression, d_credit, rank, games_quantity, won_games, lost_games, created_at
    FROM highest_ranks
    WHERE user_id != (SELECT user_id FROM user_rank_details)
    UNION
    -- Users Above the current user
    SELECT user_id, sid, username, score, impression, d_credit, rank, games_quantity, won_games, lost_games, created_at
    FROM above_user_ranks
    UNION
    -- Users Below the current user
    SELECT user_id, sid, username, score, impression, d_credit, rank, games_quantity, won_games, lost_games, created_at
    FROM below_user_ranks
) AS final_result
ORDER BY rank ASC;
`
	rows, err := c.db.Query(query, userId)
	if err != nil {
		return nil, types.NewInternalError("failed to fetch the data #3023")
	}
	defer rows.Close()

	profiles := make([]models.Profile, 0)
	for rows.Next() {
		var profile models.Profile
		err := rows.Scan(&profile.UserId, &profile.Sid, &profile.Username, &profile.Score, &profile.Impression, &profile.DCoin, &profile.Rank, &profile.GamesQuantity, &profile.WonGames, &profile.LostGames, &profile.CreatedAt)
		if err != nil {
			return nil, types.NewInternalError("failed to fetch the data #3024")
		}
		profiles = append(profiles, profile)
	}
	return profiles, nil
}
