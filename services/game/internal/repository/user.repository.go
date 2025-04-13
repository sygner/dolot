package repository

import (
	"database/sql"
	"dolott_game/internal/models"
	"dolott_game/internal/types"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/lib/pq"
)

// GetGameById retrieves a game by its ID.
func (c *gameRepository) GetGameById(gameId string) (*models.Game, *types.Error) {
	query := `SELECT id, name, game_type, num_main_numbers, num_bonus_numbers, main_number_range, bonus_number_range, start_time, end_time, creator_id, result, prize, auto_compute_prize, created_at 
			  FROM games WHERE id = $1`

	var game models.Game
	if err := c.db.QueryRow(query, gameId).Scan(
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
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, types.NewNotFoundError("no game found, error code #4013")
		}
		return nil, types.NewInternalError("internal issue, error code #4014")
	}

	return &game, nil
}

// GetUserChoicesByGameId retrieves user choices based on the game ID.
func (c *gameRepository) GetUserChoicesByGameId(gameId string) ([]models.UserChoiceResult, *types.Error) {
	query := `SELECT user_id, chosen_main_numbers, chosen_bonus_numbers, bought_price FROM user_choices WHERE game_id = $1`
	rows, err := c.db.Query(query, gameId)
	if err != nil {
		return nil, types.NewInternalError("internal issue, error code #4015")
	}
	defer rows.Close()

	var userChoices []models.UserChoiceResult
	for rows.Next() {
		var choice models.UserChoiceResult
		var chosenMainNumbers, chosenBonusNumbers []byte

		if err := rows.Scan(&choice.UserId, &chosenMainNumbers, &chosenBonusNumbers, &choice.BoughtPrice); err != nil {
			return nil, types.NewInternalError("internal issue, error code #4016")
		}

		choice.ChosenNumbers, _ = parse2DIntArray(string(chosenMainNumbers))
		if err != nil {
			return nil, types.NewInternalError("internal issue, error code #4017")
		}

		choice.ChosenBonusNumber, err = parseFlatIntArray(string(chosenBonusNumbers))
		if err != nil {
			return nil, types.NewInternalError("internal issue, error code #4018")
		}

		userChoices = append(userChoices, choice)
	}

	if err = rows.Err(); err != nil {
		return nil, types.NewInternalError("internal issue, error code #4019")
	}
	return userChoices, nil
}

func (c *gameRepository) GetUserChoicesByGameIdAndUserId(userId int32, gameId string) ([]models.UserChoiceResult, *types.Error) {
	query := `SELECT user_id, chosen_main_numbers, chosen_bonus_numbers, bought_pric FROM user_choices WHERE user_id = $1 AND game_id = $2`
	rows, err := c.db.Query(query, userId, gameId)
	if err != nil {
		return nil, types.NewInternalError("internal issue, error code #4015-1")
	}
	defer rows.Close()

	var userChoices []models.UserChoiceResult
	for rows.Next() {
		var choice models.UserChoiceResult
		var chosenMainNumbers, chosenBonusNumbers []byte

		if err := rows.Scan(&choice.UserId, &chosenMainNumbers, &chosenBonusNumbers, &choice.BoughtPrice); err != nil {
			return nil, types.NewInternalError("internal issue, error code #4016-1")
		}

		choice.ChosenNumbers, _ = parse2DIntArray(string(chosenMainNumbers))
		if err != nil {
			return nil, types.NewInternalError("internal issue, error code #4017-1")
		}

		choice.ChosenBonusNumber, err = parseFlatIntArray(string(chosenBonusNumbers))
		if err != nil {
			return nil, types.NewInternalError("internal issue, error code #4018-1")
		}

		userChoices = append(userChoices, choice)
	}

	if err = rows.Err(); err != nil {
		return nil, types.NewInternalError("internal issue, error code #4019-1")
	}
	return userChoices, nil
}

// parse2DIntArray parses a 2D int32 array from string representation.
func parse2DIntArray(arrayStr string) ([][]int32, error) {
	arrayStr = strings.Trim(arrayStr, "{}") // Remove outer braces
	if arrayStr == "" {
		return [][]int32{}, nil
	}

	var result [][]int32
	subArrayStrs := strings.Split(arrayStr, "},{")
	for _, subArrayStr := range subArrayStrs {
		subArrayStr = strings.Trim(subArrayStr, "{}")
		numStrs := strings.Split(subArrayStr, ",")
		var subArray []int32

		for _, numStr := range numStrs {
			numStr = strings.TrimSpace(numStr)
			if numStr == "" {
				continue
			}

			num, err := strconv.ParseInt(numStr, 10, 32)
			if err != nil {
				return nil, fmt.Errorf("error parsing number %s: %v", numStr, err)
			}
			subArray = append(subArray, int32(num))
		}
		result = append(result, subArray)
	}
	return result, nil
}

// parseFlatIntArray parses a flat int32 array from string representation.
func parseFlatIntArray(arrayStr string) ([]int32, error) {
	arrayStr = strings.Trim(arrayStr, "{}")
	if arrayStr == "" {
		return []int32{}, nil
	}

	numStrs := strings.Split(arrayStr, "},{")
	var result []int32

	for _, numStr := range numStrs {
		numStr = strings.Trim(numStr, "{}")
		if numStr == "" {
			continue
		}

		num, err := strconv.ParseInt(numStr, 10, 32)
		if err != nil {
			return nil, fmt.Errorf("error parsing number %s: %v", numStr, err)
		}
		result = append(result, int32(num))
	}
	return result, nil
}

// Helper function to parse a PostgreSQL multidimensional array string into a [][]int32
func parseMultiDimensionalArray(arrayStr string) ([][]int32, error) {
	// Remove the outer curly braces
	arrayStr = strings.Trim(arrayStr, "{}")

	// Split by outer commas to get individual arrays
	outerArray := strings.Split(arrayStr, "},{")

	// Prepare the result slice
	var result [][]int32

	// Iterate over each inner array
	for _, innerArray := range outerArray {
		// Remove any remaining curly braces
		innerArray = strings.Trim(innerArray, "{}")

		// Split the inner array by commas to get individual numbers
		numStrings := strings.Split(innerArray, ",")

		// Convert each number string to int32
		var nums []int32
		for _, numStr := range numStrings {
			var num int32
			if err := json.Unmarshal([]byte(numStr), &num); err != nil {
				return nil, fmt.Errorf("failed to parse number: %v", err)
			}
			nums = append(nums, num)
		}

		// Append the inner array to the result
		result = append(result, nums)
	}

	return result, nil
}

func (c *gameRepository) FindUsersByResultAndGameId(gameId string) ([]models.DivisionResult, *types.Error) {
	game, err := c.GetGameById(gameId)
	if err != nil {
		return nil, err
	}

	// Parse the result
	var winningNumbers []int32
	var bonusNumber int32
	if game.Result != nil {
		resultParts := strings.Split(*game.Result, "+")
		winningNumbers = parseNumbers(resultParts[0])

		if len(resultParts) > 1 {
			if intValue, err := strconv.Atoi(resultParts[1]); err == nil {
				bonusNumber = int32(intValue)
			} else {
				return nil, types.NewInternalError("internal issue: parsing bonus number, error code #4020")
			}
		}
	}

	userChoices, err := c.GetUserChoicesByGameId(gameId)
	if err != nil {
		return nil, err
	}

	var divisionUsers []models.DivisionResult

	// Perform division logic based on game type
	var gameType int32
	switch game.GameType {
	case models.Lotto.String():
		gameType = 0
		divisionUsers = c.GetLottoDivisionUsers(winningNumbers, userChoices, []int{6, 5, 4, 3})

	case models.Ozlotto.String():
		gameType = 1
		divisionUsers = c.GetLottoDivisionUsers(winningNumbers, userChoices, []int{7, 6, 5, 4})

	case models.Powerball.String():
		gameType = 2
		divisionUsers = c.GetPowerballDivisionUsersWithBonus(winningNumbers, bonusNumber, userChoices, models.PowerBallDivisions)

	case models.AmericanPowerball.String():
		gameType = 3
		divisionUsers = c.GetPowerballDivisionUsersWithBonus(winningNumbers, bonusNumber, userChoices, models.AmericanPowerballDivisions)

	default:
		return nil, types.NewBadRequestError("game type not found #4021")
	}

	// Convert divisionUsers to JSON format
	divisionData, rerr := json.Marshal(divisionUsers)
	if rerr != nil {
		return nil, types.NewInternalError("failed to marshal division data, error code #4022")
	}
	prize, err := c.ComputePrizeForGame(game)
	if err != nil {
		return nil, err
	}
	jackPot := false
	for _, jk := range divisionUsers {
		if jk.Division == "Division 1" {
			jackPot = true
			break
		}
	}
	// Insert the winners into the "winners" table
	query := `INSERT INTO winners (game_id, game_type, divisions, result_number, prize, jackpot) VALUES ($1, $2, $3, $4, $5, $6)`
	_, rerr = c.db.Exec(query, gameId, gameType, divisionData, *game.Result, prize.Prize, jackPot)
	if rerr != nil {
		fmt.Println(rerr)
		return nil, types.NewInternalError("failed to insert winners data, error code #4023")
	}

	return divisionUsers, nil
}

func (c *gameRepository) GetUsersByResultAndGameId(gameId string) ([]models.DivisionResult, *types.Error) {
	game, err := c.GetGameById(gameId)
	if err != nil {
		return nil, err
	}

	// Parse the result
	var winningNumbers []int32
	var bonusNumber int32
	if game.Result != nil {
		resultParts := strings.Split(*game.Result, "+")
		winningNumbers = parseNumbers(resultParts[0])

		if len(resultParts) > 1 {
			if intValue, err := strconv.Atoi(resultParts[1]); err == nil {
				bonusNumber = int32(intValue)
			} else {
				return nil, types.NewInternalError("internal issue: parsing bonus number, error code #4020")
			}
		}
	}

	userChoices, err := c.GetUserChoicesByGameId(gameId)
	if err != nil {
		return nil, err
	}

	var divisionUsers []models.DivisionResult

	// Perform division logic based on game type
	switch game.GameType {
	case models.Lotto.String():
		divisionUsers = c.GetLottoDivisionUsers(winningNumbers, userChoices, []int{6, 5, 4, 3})

	case models.Ozlotto.String():
		divisionUsers = c.GetLottoDivisionUsers(winningNumbers, userChoices, []int{7, 6, 5, 4})

	case models.Powerball.String():
		divisionUsers = c.GetPowerballDivisionUsersWithBonus(winningNumbers, bonusNumber, userChoices, models.PowerBallDivisions)

	case models.AmericanPowerball.String():
		divisionUsers = c.GetPowerballDivisionUsersWithBonus(winningNumbers, bonusNumber, userChoices, models.AmericanPowerballDivisions)

	default:
		return nil, types.NewBadRequestError("game type not found #4021")
	}

	// // Convert divisionUsers to JSON format
	// divisionData, rerr := json.Marshal(divisionUsers)
	// if rerr != nil {
	// 	return nil, types.NewInternalError("failed to marshal division data, error code #4022")
	// }

	// // Insert the winners into the "winners" table
	// query := `INSERT INTO winners (game_id, divisions, result_number) VALUES ($1, $2, $3)`
	// _, rerr = c.db.Exec(query, gameId, divisionData, *game.Result)
	// if rerr != nil {
	// 	return nil, types.NewInternalError("failed to insert winners data, error code #4023")
	// }
	return divisionUsers, nil
}

func (c *gameRepository) GetUserByResultAndGameId(userId int32, gameId string) ([]models.DivisionResult, *types.Error) {
	game, err := c.GetGameById(gameId)
	if err != nil {
		return nil, err
	}

	// Parse the result
	var winningNumbers []int32
	var bonusNumber int32
	if game.Result != nil {
		resultParts := strings.Split(*game.Result, "+")
		winningNumbers = parseNumbers(resultParts[0])

		if len(resultParts) > 1 {
			if intValue, err := strconv.Atoi(resultParts[1]); err == nil {
				bonusNumber = int32(intValue)
			} else {
				return nil, types.NewInternalError("internal issue: parsing bonus number, error code #4020")
			}
		}
	}

	userChoices, err := c.GetUserChoicesByGameIdAndUserId(userId, gameId)
	if err != nil {
		return nil, err
	}

	var divisionUsers []models.DivisionResult

	// Perform division logic based on game type
	switch game.GameType {
	case models.Lotto.String():
		divisionUsers = c.GetLottoDivisionUsers(winningNumbers, userChoices, []int{6, 5, 4, 3})

	case models.Ozlotto.String():
		divisionUsers = c.GetLottoDivisionUsers(winningNumbers, userChoices, []int{7, 6, 5, 4})

	case models.Powerball.String():
		divisionUsers = c.GetPowerballDivisionUsersWithBonus(winningNumbers, bonusNumber, userChoices, models.PowerBallDivisions)

	case models.AmericanPowerball.String():
		divisionUsers = c.GetPowerballDivisionUsersWithBonus(winningNumbers, bonusNumber, userChoices, models.AmericanPowerballDivisions)

	default:
		return nil, types.NewBadRequestError("game type not found #4021")
	}

	// Convert divisionUsers to JSON format
	// divisionData, rerr := json.Marshal(divisionUsers)
	// if rerr != nil {
	// 	return nil, types.NewInternalError("failed to marshal division data, error code #4022")
	// }

	// Insert the winners into the "winners" table
	// query := `INSERT INTO winners (game_id, divisions, result_number) VALUES ($1, $2, $3)`
	// _, rerr = c.db.Exec(query, gameId, divisionData, *game.Result)
	// if rerr != nil {
	// 	return nil, types.NewInternalError("failed to insert winners data, error code #4023")
	// }

	return divisionUsers, nil
}

// parseNumbers parses a string of comma-separated numbers into an array of int32.
func parseNumbers(numbersString string) []int32 {
	var numbers []int32
	for _, numStr := range strings.Split(strings.TrimSpace(numbersString), ",") {
		if numStr != "" {
			num, _ := strconv.Atoi(numStr) // Ignore error since it's handled in FindUsersByResultAndGameId
			numbers = append(numbers, int32(num))
		}
	}
	return numbers
}

// AddUserChoice adds a user's choice to the database.
func (c *gameRepository) AddUserChoice(data *models.AddUserChoiceDTO) *types.Error {
	query := `INSERT INTO user_choices (id, user_id, game_id, chosen_main_numbers, chosen_bonus_numbers, bought_price, created_at) VALUES ($1, $2, $3, $4, $5, $6, NOW())`
	if _, err := c.db.Exec(query, data.Id, data.UserId, data.GameId, pq.Array(data.ChosenMainNumbers), pq.Array(data.ChosenBonusNumbers), data.BoughtPrice); err != nil {
		return types.NewInternalError("failed to add the user choice, error code #4021")
	}
	return nil
}

func (c *gameRepository) GetUserChoicesByUserId(userId int32, pagination *models.Pagination) ([]models.UserChoice, *types.Error) {
	// SQL query to get user choices by userId
	query := `SELECT id, game_id, chosen_main_numbers, chosen_bonus_numbers, created_at 
	          FROM user_choices 
	          WHERE user_id = $1 OFFSET $2 LIMIT $3`

	// Prepare a slice to hold the results
	var userChoices []models.UserChoice

	// Execute the query and fetch results
	rows, err := c.db.Query(query, userId, pagination.Offset, pagination.Limit)
	if err != nil {
		return nil, types.NewInternalError("failed to query user choices by userId, error code #4026")
	}
	defer rows.Close()

	// Loop through the result rows
	for rows.Next() {
		var userChoice models.UserChoice
		var chosenMainNumbers string
		var chosenBonusNumbers sql.NullString

		// Scan the main numbers and bonus numbers as strings
		err := rows.Scan(&userChoice.Id, &userChoice.GameId, &chosenMainNumbers, &chosenBonusNumbers, &userChoice.CreatedAt)
		if err != nil {
			return nil, types.NewInternalError("failed to scan user choice row, error code #4027")
		}

		// Parse the main numbers from the string into [][]int32
		userChoice.ChosenMainNumbers, err = parseMultiDimensionalArray(chosenMainNumbers)
		if err != nil {
			return nil, types.NewInternalError("failed to parse chosen main numbers, error code #4028")
		}

		// Parse the bonus numbers if they exist
		if chosenBonusNumbers.Valid {
			userChoice.ChosenBonusNumbers, err = parseMultiDimensionalArray(chosenBonusNumbers.String)
			if err != nil {
				return nil, types.NewInternalError("failed to parse chosen bonus numbers, error code #4029")
			}
		}

		// Append to the result slice
		userChoices = append(userChoices, userChoice)
	}

	// Handle any errors encountered during iteration
	if err = rows.Err(); err != nil {
		return nil, types.NewInternalError("failed to process rows, error code #4030")
	}

	// Return the list of user choices
	return userChoices, nil
}

func (c *gameRepository) GetUserChoicesByUserIdAndTimeRange(userId int32, startTime, endTime time.Time) ([]models.UserChoice, *types.Error) {
	// SQL query to get user choices by userId and within the specified time range
	query := `SELECT id, game_id, user_id, chosen_main_numbers, chosen_bonus_numbers
	          FROM user_choices 
	          WHERE user_id = $1 AND created_at BETWEEN $2 AND $3`

	rows, err := c.db.Query(query, userId, startTime, endTime)
	if err != nil {
		return nil, types.NewInternalError("internal issue, error code #4015")
	}
	defer rows.Close()

	var userChoices []models.UserChoice
	for rows.Next() {
		var choice models.UserChoice
		var chosenMainNumbers, chosenBonusNumbers []byte

		if err := rows.Scan(&choice.Id, &choice.GameId, &choice.UserId, &chosenMainNumbers, &chosenBonusNumbers); err != nil {
			return nil, types.NewInternalError("internal issue, error code #4016")
		}

		cmn, err := parse2DIntArray(string(chosenMainNumbers))
		if err != nil {
			return nil, types.NewInternalError("internal issue, error code #4017")
		}

		cbn, err := parse2DIntArray(string(chosenBonusNumbers))
		if err != nil {
			return nil, types.NewInternalError("internal issue, error code #4018")
		}

		choice.ChosenMainNumbers = cmn
		choice.ChosenBonusNumbers = cbn

		userChoices = append(userChoices, choice)
	}

	if err = rows.Err(); err != nil {
		return nil, types.NewInternalError("internal issue, error code #4019")
	}
	return userChoices, nil
}

func (c *gameRepository) GetUserChoicesCountByUserId(userId int32) (int32, *types.Error) {
	query := `SELECT count(*) FROM "user_choices" WHERE user_id = $1`
	var totalCount int32
	err := c.db.QueryRow(query, userId).Scan(&totalCount)
	if err != nil {
		return 0, types.NewInternalError("internal issue, error code #4032")
	}
	return totalCount, nil
}

func (c *gameRepository) GetUserChoicesByGameIdAndPagination(gameId string, pagination *models.Pagination) ([]models.UserChoice, *types.Error) {
	// SQL query to get user choices by userId
	query := `SELECT id, game_id, user_id, chosen_main_numbers, chosen_bonus_numbers, created_at 
	          FROM user_choices 
	          WHERE game_id = $1 OFFSET $2 LIMIT $3`

	// Prepare a slice to hold the results
	var userChoices []models.UserChoice

	// Execute the query and fetch results
	rows, err := c.db.Query(query, gameId, pagination.Offset, pagination.Limit)
	if err != nil {
		return nil, types.NewInternalError("failed to query user choices by game id, error code #4031")
	}
	defer rows.Close()

	// Loop through the result rows
	for rows.Next() {
		var userChoice models.UserChoice
		var chosenMainNumbers string
		var chosenBonusNumbers sql.NullString

		// Scan the main numbers and bonus numbers as strings
		err := rows.Scan(&userChoice.Id, &userChoice.GameId, &userChoice.UserId, &chosenMainNumbers, &chosenBonusNumbers, &userChoice.CreatedAt)
		if err != nil {
			return nil, types.NewInternalError("failed to scan user choice row, error code #4032")
		}

		// Parse the main numbers from the string into [][]int32
		userChoice.ChosenMainNumbers, err = parseMultiDimensionalArray(chosenMainNumbers)
		if err != nil {
			return nil, types.NewInternalError("failed to parse chosen main numbers, error code #4033")
		}

		// Parse the bonus numbers if they exist
		if chosenBonusNumbers.Valid {
			userChoice.ChosenBonusNumbers, err = parseMultiDimensionalArray(chosenBonusNumbers.String)
			if err != nil {
				return nil, types.NewInternalError("failed to parse chosen bonus numbers, error code #4034")
			}
		}

		// Append to the result slice
		userChoices = append(userChoices, userChoice)
	}
	return userChoices, nil
}

func (c *gameRepository) GetAllUserGames(userId int32) ([]string, *types.Error) {
	// SQL query to get user choices by userId
	query := `SELECT game_id FROM user_choices 
	          WHERE user_id = $1 ORDER BY created_at DESC`

	// Prepare a slice to hold the results
	gameIds := make([]string, 0)

	// Execute the query and fetch results
	rows, err := c.db.Query(query, userId)
	if err != nil {
		return nil, types.NewInternalError("failed to query user choices by game id, error code #4043")
	}
	defer rows.Close()

	// Loop through the result rows
	for rows.Next() {
		var gameId string
		// Scan the main numbers and bonus numbers as strings
		err := rows.Scan(&gameId)
		if err != nil {
			return nil, types.NewInternalError("failed to scan user choice row, error code #4044")
		}

		// Append to the result slice
		gameIds = append(gameIds, gameId)
	}
	return gameIds, nil
}

func (c *gameRepository) GetUserChoicesCountByGameId(gameId string) (int32, *types.Error) {
	query := `SELECT count(*) FROM "user_choices" WHERE game_id = $1`
	var totalCount int32
	err := c.db.QueryRow(query, gameId).Scan(&totalCount)
	if err != nil {
		fmt.Println(err)
		return 0, types.NewInternalError("internal issue, error code #4081")
	}
	return totalCount, nil
}

func (c *gameRepository) GetUserChoicesEachCountByGameId(gameId string) (int32, *types.Error) {
	query := `
		SELECT SUM(array_length(chosen_main_numbers, 1)) 
		FROM "user_choices" 
		WHERE game_id = $1`

	var totalCount int32
	err := c.db.QueryRow(query, gameId).Scan(&totalCount)
	if err != nil {
		fmt.Println(err)
		return 0, types.NewInternalError("internal issue, error code #4081")
	}
	return totalCount, nil
}

func (c *gameRepository) ComputePrizeForGame(game *models.Game) (*models.Game, *types.Error) {
	if game == nil {
		return nil, types.NewInternalError("invalid game data, error code #4704")
	}
	// If auto_compute_prize is true, calculate prize from user choices
	if game.AutoCompute {
		userChoicesCount, err := c.GetUserChoicesEachCountByGameId(game.Id)
		if err != nil {
			return nil, err
		}
		*game.Prize = uint32(userChoicesCount) // Set the computed prize
	}

	return game, nil
}
