package repository

import (
	"database/sql"
	"dolott_game/internal/models"
	"dolott_game/internal/types"
	"encoding/json"
)

func (c *gameRepository) GetWinnersByGameId(gameId string) ([]models.DivisionResult, *types.Error) {
	// Query to get the winners from the 'winners' table by game_id
	query := `SELECT divisions FROM winners WHERE game_id = $1`

	var divisionsJSON []byte
	err := c.db.QueryRow(query, gameId).Scan(&divisionsJSON)
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

	// Return the division results
	return divisionResults, nil
}
