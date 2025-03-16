package repository

import (
	"database/sql"
	"dolott_game/internal/models"
	"dolott_game/internal/types"
	"encoding/json"
)

func (c *gameRepository) GetWinnersByGameId(gameId string) (*models.Winners, *types.Error) {
	// Query to get the winners from the 'winners' table by game_id
	query := `SELECT id, game_id, game_type, divisions, result_number, prize, jackpot, total_paid, created_at FROM winners WHERE game_id = $1`

	var divisionsJSON []byte
	var winners models.Winners
	err := c.db.QueryRow(query, gameId).Scan(&winners.Id, &winners.GameId, &winners.GameType, &divisionsJSON, &winners.ResultNumber, &winners.Prize, &winners.JackPot, &winners.TotalPaid, &winners.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, types.NewNotFoundError("no winners found for this game, error code #4024")
		}
		return nil, types.NewInternalError("failed to query winners by gameId, error code #4025")
	}

	// Deserialize the JSONB divisions into a slice of DivisionResult structs
	var divisionResults []models.DivisionResult
	err = json.Unmarshal(divisionsJSON, &divisionResults)
	if err != nil {
		return nil, types.NewInternalError("internal issue, error code #4026")
	}

	winners.Divisions = divisionResults

	// Return the division results
	return &winners, nil
}

func (c *gameRepository) GetLastWinnersByGameType(gameType int32) (*models.Winners, *types.Error) {
	// Query to get the most recent game of the given game_type that has ended
	query := `
        SELECT id, game_id, game_type, divisions, result_number, prize, jackpot, total_paid, created_at FROM winners 
        WHERE game_type = $1 
          AND created_at < NOW() 
        ORDER BY created_at DESC 
        LIMIT 1
    `

	var divisionsJSON []byte
	var winners models.Winners
	err := c.db.QueryRow(query, gameType).Scan(&winners.Id, &winners.GameId, &winners.GameType, &divisionsJSON, &winners.ResultNumber, &winners.Prize, &winners.JackPot, &winners.TotalPaid, &winners.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, types.NewNotFoundError("no winners found for this game, error code #4024")
		}
		return nil, types.NewInternalError("failed to query winners by gameId, error code #4025")
	}

	// Deserialize the JSONB divisions into a slice of DivisionResult structs
	var divisionResults []models.DivisionResult
	err = json.Unmarshal(divisionsJSON, &divisionResults)
	if err != nil {
		return nil, types.NewInternalError("internal issue, error code #4026")
	}

	winners.Divisions = divisionResults

	// Return the division results
	return &winners, nil
}

func (c *gameRepository) UpdateTotalPaidUsers(gameId string, totalPaid string) *types.Error {
	query := `UPDATE winners SET total_paid = $1 WHERE game_id = $2`
	_, err := c.db.Exec(query, totalPaid, gameId)
	if err != nil {
		return types.NewInternalError("internal issue, error code #4049")
	}
	return nil
}
