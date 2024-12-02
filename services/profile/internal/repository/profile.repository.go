package repository

import (
	"database/sql"
	"dolott_profile/internal/models"
	"dolott_profile/internal/types"
)

func (c *profileRepository) GetProfileByUsername(username string) (*models.Profile, *types.Error) {
	query := `SELECT user_id, sid, username, score, impression, rank, games_quantity, won_games, lost_games, created_at FROM profiles WHERE username = $1`

	var profile models.Profile
	err := c.db.QueryRow(query, username).Scan(
		&profile.UserId, &profile.Sid, &profile.Username, &profile.Score,
		&profile.Impression, &profile.Rank, &profile.GamesQuantity, &profile.WonGames, &profile.LostGames, &profile.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, types.NewNotFoundError("profile not found #3002")
		}
		return nil, types.NewInternalError("failed to fetch the data #3001")
	}
	return &profile, nil
}

func (c *profileRepository) GetProfileBySid(sid string) (*models.Profile, *types.Error) {
	query := `SELECT user_id, sid, username, score, impression, rank, games_quantity, won_games, lost_games, created_at FROM profiles WHERE sid = $1`

	var profile models.Profile
	err := c.db.QueryRow(query, sid).Scan(
		&profile.UserId, &profile.Sid, &profile.Username, &profile.Score,
		&profile.Impression, &profile.Rank, &profile.GamesQuantity, &profile.WonGames, &profile.LostGames, &profile.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, types.NewNotFoundError("profile not found #3019")
		}
		return nil, types.NewInternalError("failed to fetch the data #3020")
	}
	return &profile, nil
}

func (c *profileRepository) GetProfileByUserId(userId int32) (*models.Profile, *types.Error) {
	query := `SELECT user_id, sid, username, score, impression, rank, games_quantity, won_games, lost_games, created_at FROM profiles WHERE user_id = $1`

	var profile models.Profile
	err := c.db.QueryRow(query, userId).Scan(
		&profile.UserId, &profile.Sid, &profile.Username, &profile.Score,
		&profile.Impression, &profile.Rank, &profile.GamesQuantity, &profile.WonGames, &profile.LostGames, &profile.CreatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, types.NewNotFoundError("profile not found #3003")
		}
		return nil, types.NewInternalError("failed to fetch the data #3004")
	}
	return &profile, nil
}

func (c *profileRepository) AddProfile(data *models.AddProfileDTO) *types.Error {
	query := `INSERT INTO profiles (user_id, sid, username, score, impression, rank, games_quantity, won_games, lost_games, created_at) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,NOW())`
	_, err := c.db.Exec(query, data.UserId, data.Sid, data.Username, data.Score, data.Impression, data.Rank, data.GamesQuantity, data.WonGames, data.LostGames)
	if err != nil {
		return types.NewInternalError("failed to store the data #3005")
	}
	return nil
}

func (c *profileRepository) UpdateProfile(userId int32, username string) *types.Error {
	query := `UPDATE profiles SET username = $1 WHERE user_id = $2`
	_, err := c.db.Exec(query, username, userId)
	if err != nil {
		return types.NewInternalError("failed to store the data #3016")
	}
	return nil
}

func (c *profileRepository) CheckUsernameExists(username string) *types.Error {
	query := `SELECT 1 FROM profiles WHERE username = $1`

	var exists int32
	err := c.db.QueryRow(query, username).Scan(&exists)
	if err != nil {
		if err == sql.ErrNoRows {
			return types.NewNotFoundError("profile not found #3017")
		}
		return types.NewInternalError("failed to fetch the data #3018")
	}
	return nil
}
