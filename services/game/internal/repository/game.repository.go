package repository

import (
	"database/sql"
	"dolott_game/internal/models"
	"dolott_game/internal/types"
	"fmt"
)

func (c *gameRepository) GetGameByGameId(gameId string) (*models.Game, *types.Error) {
	query := `SELECT id, name, game_type, num_main_numbers, num_bonus_numbers, main_number_range, bonus_number_range, start_time, end_time, creator_id, result, prize, auto_compute_prize, created_at FROM games WHERE id = $1`
	var data models.Game
	row := c.db.QueryRow(query, gameId).Scan(&data.Id, &data.Name, &data.GameType, &data.NumMainNumbers, &data.NumBonusNumbers, &data.MainNumberRange, &data.BonusNumberRange, &data.StartTime, &data.EndTime, &data.CreatorId, &data.Result, &data.Prize, &data.AutoCompute, &data.CreatedAt)
	if row != nil {
		fmt.Println(row)
		if row == sql.ErrNoRows {
			return nil, types.NewNotFoundError("game not found, error code #4001")
		}
		return nil, types.NewInternalError("failed to fetch data, error code #4002")
	}
	return &data, nil
}

func (c *gameRepository) AddGame(data *models.AddGameDTO) *types.Error {
	query := `INSERT INTO games(id, name, game_type, num_main_numbers, num_bonus_numbers, main_number_range, bonus_number_range, start_time, end_time, creator_id, result, prize, auto_compute_prize, created_at) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,NOW())`
	_, err := c.db.Exec(query, data.Id, data.Name, data.GameTypeString, data.NumMainNumbers, data.NumBonusNumbers, data.MainNumberRange, data.BonusNumberRange, data.StartTime, data.EndTime, data.CreatorId, data.Result, data.Prize, data.AutoCompute)
	if err != nil {
		fmt.Println(err)
		return types.NewInternalError("failed to add game, error code #4003")
	}
	return nil
}

func (c *gameRepository) GetNextGamesByGameType(gameType string, limit int32) ([]models.Game, *types.Error) {
	query := `SELECT id, name, game_type, num_main_numbers, num_bonus_numbers, main_number_range, bonus_number_range, start_time, end_time, creator_id, result, prize, auto_compute_prize, created_at FROM games WHERE end_time > NOW() AND game_type = $1 LIMIT $2`

	rows, err := c.db.Query(query, gameType, limit)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, types.NewNotFoundError("game not found, error code #4005")
		}
		return nil, types.NewInternalError("failed to fetch the games, error code #4004")
	}
	defer rows.Close()

	games := make([]models.Game, 0)
	for rows.Next() {
		var game models.Game
		rows.Scan(
			&game.Id,
			&game.Name,
			&game.GameType,
			&game.NumMainNumbers,
			&game.NumBonusNumbers,
			&game.MainNumberRange,
			&game.BonusNumberRange,
			&game.StartTime,
			&game.EndTime,
			&game.CreatorId,
			&game.Result,
			&game.Prize,
			&game.AutoCompute,
			&game.CreatedAt)
		games = append(games, game)
	}
	return games, nil
}

func (c *gameRepository) GetAllNextGames() ([]models.Game, *types.Error) {
	query := `SELECT id, name, game_type, num_main_numbers, num_bonus_numbers, main_number_range, bonus_number_range, start_time, end_time, creator_id, result, prize, auto_compute_prize, created_at FROM games WHERE end_time > NOW()`

	rows, err := c.db.Query(query)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, types.NewNotFoundError("game not found, error code #4041")
		}
		return nil, types.NewInternalError("failed to fetch the games, error code #4042")
	}
	defer rows.Close()

	games := make([]models.Game, 0)
	for rows.Next() {
		var game models.Game
		rows.Scan(
			&game.Id,
			&game.Name,
			&game.GameType,
			&game.NumMainNumbers,
			&game.NumBonusNumbers,
			&game.MainNumberRange,
			&game.BonusNumberRange,
			&game.StartTime,
			&game.EndTime,
			&game.CreatorId,
			&game.Result,
			&game.Prize,
			&game.AutoCompute,
			&game.CreatedAt)
		games = append(games, game)
	}
	return games, nil
}

func (c *gameRepository) GetAllPreviousGames(offset, limit int32) ([]models.Game, *types.Error) {
	query := `SELECT id, name, game_type, num_main_numbers, num_bonus_numbers, main_number_range, bonus_number_range, start_time, end_time, creator_id, result, prize, auto_compute_prize, created_at FROM games WHERE end_time < NOW() LIMIT $1 OFFSET $2`

	rows, err := c.db.Query(query, limit, offset)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, types.NewNotFoundError("game not found, error code #4035")
		}
		return nil, types.NewInternalError("failed to fetch the games, error code #4036")
	}
	defer rows.Close()

	games := make([]models.Game, 0)
	for rows.Next() {
		var game models.Game
		rows.Scan(
			&game.Id,
			&game.Name,
			&game.GameType,
			&game.NumMainNumbers,
			&game.NumBonusNumbers,
			&game.MainNumberRange,
			&game.BonusNumberRange,
			&game.StartTime,
			&game.EndTime,
			&game.CreatorId,
			&game.Result,
			&game.Prize,
			&game.AutoCompute,
			&game.CreatedAt)
		games = append(games, game)
	}
	return games, nil
}

func (c *gameRepository) GetAllPreviousGamesCount() (int32, *types.Error) {
	query := `SELECT count(*) FROM "games" WHERE end_time < NOW()`
	var totalCount int32
	err := c.db.QueryRow(query).Scan(&totalCount)
	if err != nil {
		return 0, types.NewInternalError("internal issue, error code #4037")
	}
	return totalCount, nil
}

func (c *gameRepository) GetAllGames(offset, limit int32) ([]models.Game, *types.Error) {
	query := `SELECT id, name, game_type, num_main_numbers, num_bonus_numbers, main_number_range, bonus_number_range, start_time, end_time, creator_id, result, prize, auto_compute_prize, created_at FROM games ORDER BY created_at DESC LIMIT $1 OFFSET $2`

	rows, err := c.db.Query(query, limit, offset)
	if err != nil {
		fmt.Println(err)
		if err == sql.ErrNoRows {
			return nil, types.NewNotFoundError("game not found, error code #4038")
		}
		return nil, types.NewInternalError("failed to fetch the games, error code #4039")
	}
	defer rows.Close()

	games := make([]models.Game, 0)
	for rows.Next() {
		var game models.Game
		rows.Scan(
			&game.Id,
			&game.Name,
			&game.GameType,
			&game.NumMainNumbers,
			&game.NumBonusNumbers,
			&game.MainNumberRange,
			&game.BonusNumberRange,
			&game.StartTime,
			&game.EndTime,
			&game.CreatorId,
			&game.Result,
			&game.Prize,
			&game.AutoCompute,
			&game.CreatedAt)
		games = append(games, game)
	}
	return games, nil
}

func (c *gameRepository) GetAllGamesCount() (int32, *types.Error) {
	query := `SELECT count(*) FROM "games"`
	var totalCount int32
	err := c.db.QueryRow(query).Scan(&totalCount)
	if err != nil {
		fmt.Println(err)
		return 0, types.NewInternalError("internal issue, error code #4040")
	}
	return totalCount, nil
}

func (c *gameRepository) DeleteGameByGameId(gameId string) *types.Error {
	query := `DELETE FROM games WHERE id = $1`
	_, err := c.db.Exec(query, gameId)
	if err != nil {
		return types.NewInternalError("failed to delete the game, error code #4006")
	}
	return nil
}

func (c *gameRepository) GetGamesByCreatorId(creatorId int32, data *models.Pagination) ([]models.Game, *types.Error) {
	query := `SELECT id, name, game_type, num_main_numbers, num_bonus_numbers, main_number_range, bonus_number_range, start_time, end_time, creator_id, result, prize, auto_compute_prize, created_at FROM games WHERE creator_id = $1 LIMIT $2 OFFSET $3`

	rows, err := c.db.Query(query, creatorId, data.Limit, data.Offset)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, types.NewNotFoundError("no games found for creator, error code #4008")
		}
		return nil, types.NewInternalError("failed to fetch games by creator, error code #4009")
	}
	defer rows.Close()

	games := make([]models.Game, 0)
	for rows.Next() {
		var game models.Game
		rows.Scan(
			&game.Id,
			&game.Name,
			&game.GameType,
			&game.NumMainNumbers,
			&game.NumBonusNumbers,
			&game.MainNumberRange,
			&game.BonusNumberRange,
			&game.StartTime,
			&game.EndTime,
			&game.CreatorId,
			&game.Result,
			&game.Prize,
			&game.AutoCompute,
			&game.CreatedAt)
		games = append(games, game)
	}
	return games, nil
}

func (c *gameRepository) GetGameTypes() ([]models.GameTypeDetail, *types.Error) {
	query := `SELECT id, name, description, type_name, day_name, prize_reward, token_burn, auto_compute FROM game_types`

	rows, err := c.db.Query(query)
	if err != nil {
		fmt.Println(err)
		if err == sql.ErrNoRows {
			return nil, types.NewNotFoundError("no game types found, error code #4013")
		}
		return nil, types.NewInternalError("failed to fetch game types, error code #4014")
	}
	defer rows.Close()

	games := make([]models.GameTypeDetail, 0)
	for rows.Next() {
		var game models.GameTypeDetail
		rows.Scan(
			&game.Id,
			&game.Name,
			&game.Description,
			&game.TypeName,
			&game.DayName,
			&game.PrizeReward,
			&game.TokenBurn,
			&game.AutoCompute,
		)
		games = append(games, game)
	}
	return games, nil
}

func (c *gameRepository) GetGamesCountByCreatorId(creatorId int32) (int32, *types.Error) {
	query := `SELECT count(*) FROM "games" WHERE creator_id = $1`
	var totalCount int32
	err := c.db.QueryRow(query, creatorId).Scan(&totalCount)
	if err != nil {
		return 0, types.NewInternalError("internal issue, error code #4011")
	}
	return totalCount, nil
}

func (c *gameRepository) CheckGameExistsById(gameId string) (bool, *types.Error) {
	query := `SELECT COUNT(1) FROM games WHERE id = $1`
	var count int
	err := c.db.QueryRow(query, gameId).Scan(&count)
	if err != nil {
		return false, types.NewInternalError("failed to check if game exists, error code #4010")
	}
	return count > 0, nil
}
func (c *gameRepository) CheckGameExistsByIdAndEndTime(gameId string) (bool, *types.Error) {
	query := `SELECT COUNT(1) FROM games WHERE id = $1 AND end_time > NOW()`
	var count int
	err := c.db.QueryRow(query, gameId).Scan(&count)
	if err != nil {
		return false, types.NewInternalError("failed to check if game exists, error code #4019")
	}
	return count > 0, nil
}

func (c *gameRepository) AddResultByGameId(gameId string, result string) *types.Error {
	query := `UPDATE games SET result = $1, end_time = NOW() WHERE id = $2 `
	_, err := c.db.Exec(query, result, gameId)
	if err != nil {
		return types.NewInternalError("failed to add result to the game, error code #4012")
	}
	return nil
}

func (c *gameRepository) UpdateGamePrizeByGameId(gameId string, prize *uint32, autoCompute bool) *types.Error {
	query := `UPDATE games SET prize = $1, auto_compute_prize = $2 WHERE id = $3`
	_, err := c.db.Exec(query, prize, autoCompute, gameId)
	if err != nil {
		return types.NewInternalError("failed to add result to the game, error code #4050")
	}
	return nil

}

func (c *gameRepository) UpdateGameTypeDetail(gameType int32, dayName *string, prizeReward int32, tokenBurn int32, autoCompute bool) *types.Error {
	query := `
		UPDATE game_types 
		SET 
			auto_compute = $1, 
			prize_reward = $2, 
			token_burn = $3, 
			day_name = COALESCE($4, day_name)
		WHERE id = $5`

	_, err := c.db.Exec(query, autoCompute, prizeReward, tokenBurn, dayName, gameType)
	if err != nil {
		fmt.Println(err)
		return types.NewInternalError("failed to update game detail, error code #4045")
	}
	return nil
}

func (c *gameRepository) GetAllUserPreviousGames(userId int32, offset int32, limit int32) ([]models.Game, *types.Error) {
	query := `
        SELECT DISTINCT g.id,
                        g.name,
                        g.game_type,
                        g.num_main_numbers,
                        g.num_bonus_numbers,
                        g.main_number_range,
                        g.bonus_number_range,
                        g.start_time,
                        g.end_time,
                        g.creator_id,
                        g.result,
						g.prize, 
						g.auto_compute_prize,
                        g.created_at
        FROM games g
        JOIN user_choices uc ON uc.game_id = g.id
        WHERE uc.user_id = $1
          AND g.end_time < NOW()
        ORDER BY g.created_at DESC
        LIMIT $2 OFFSET $3
    `

	rows, err := c.db.Query(query, userId, limit, offset)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, types.NewNotFoundError("game not found, error code #4046")
		}
		return nil, types.NewInternalError("failed to fetch the games, error code #4047")
	}
	defer rows.Close()

	games := make([]models.Game, 0)
	for rows.Next() {
		var game models.Game
		scanErr := rows.Scan(
			&game.Id,
			&game.Name,
			&game.GameType,
			&game.NumMainNumbers,
			&game.NumBonusNumbers,
			&game.MainNumberRange,
			&game.BonusNumberRange,
			&game.StartTime,
			&game.EndTime,
			&game.CreatorId,
			&game.Result,
			&game.Prize,
			&game.AutoCompute,
			&game.CreatedAt,
		)
		if scanErr != nil {
			return nil, types.NewInternalError("failed to scan game row, error code #4047-A")
		}
		games = append(games, game)
	}

	return games, nil
}

func (c *gameRepository) GetAllUserPreviousGamesByGameType(userId int32, gameType string, offset int32, limit int32) ([]models.Game, *types.Error) {
	query := `
        SELECT DISTINCT g.id,
                        g.name,
                        g.game_type,
                        g.num_main_numbers,
                        g.num_bonus_numbers,
                        g.main_number_range,
                        g.bonus_number_range,
                        g.start_time,
                        g.end_time,
                        g.creator_id,
                        g.result,
						g.prize, 
						g.auto_compute_prize,
                        g.created_at
        FROM games g
        JOIN user_choices uc ON uc.game_id = g.id
        WHERE uc.user_id = $1
          AND g.end_time < NOW()
		  AND g.game_type = $2
        ORDER BY g.created_at DESC
        LIMIT $3 OFFSET $4
    `

	rows, err := c.db.Query(query, userId, gameType, limit, offset)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, types.NewNotFoundError("game not found, error code #4046")
		}
		return nil, types.NewInternalError("failed to fetch the games, error code #4047")
	}
	defer rows.Close()

	games := make([]models.Game, 0)
	for rows.Next() {
		var game models.Game
		scanErr := rows.Scan(
			&game.Id,
			&game.Name,
			&game.GameType,
			&game.NumMainNumbers,
			&game.NumBonusNumbers,
			&game.MainNumberRange,
			&game.BonusNumberRange,
			&game.StartTime,
			&game.EndTime,
			&game.CreatorId,
			&game.Result,
			&game.Prize,
			&game.AutoCompute,
			&game.CreatedAt,
		)
		if scanErr != nil {
			return nil, types.NewInternalError("failed to scan game row, error code #4047-A")
		}
		games = append(games, game)
	}

	return games, nil
}

func (c *gameRepository) GetAllUserPreviousGamesCount(userId int32) (int32, *types.Error) {
	query := `
SELECT COUNT(DISTINCT g.id)
FROM games g
JOIN user_choices uc ON uc.game_id = g.id
WHERE uc.user_id = $1
  AND g.end_time < NOW()

    `
	var totalCount int32
	err := c.db.QueryRow(query, userId).Scan(&totalCount)
	if err != nil {
		return 0, types.NewInternalError("internal issue, error code #4048")
	}
	return totalCount, nil
}

func (c *gameRepository) GetAllUserPreviousGamesCountByGameType(userId int32, gameType string) (int32, *types.Error) {
	query := `
SELECT COUNT(DISTINCT g.id)
FROM games g
JOIN user_choices uc ON uc.game_id = g.id
WHERE uc.user_id = $1
  AND g.end_time < NOW()
  AND g.game_type = $2

    `
	var totalCount int32
	err := c.db.QueryRow(query, userId, gameType).Scan(&totalCount)
	if err != nil {
		return 0, types.NewInternalError("internal issue, error code #4048")
	}
	return totalCount, nil
}
