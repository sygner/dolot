package models

import (
	"time"
)

type GameTypeDetail struct {
	Id          int32  `db:"id" json:"id"`
	Name        string `db:"name" json:"name"`
	Description string `db:"description" json:"description"`
	TypeName    string `db:"type_name" json:"type_name"`
	DayName     string `db:"day_name" json:"day_name"`
	PrizeReward int32  `db:"prize_reward" json:"prize_reward"`
	TokenBurn   int32  `db:"token_burn" json:"token_burn"`
	AutoCompute bool   `db:"auto_compute" json:"auto_compute"`
}

type Game struct {
	Id               string    `db:"id" json:"id"`
	Name             string    `db:"name" json:"name"`
	GameType         string    `db:"game_type" json:"game_type"`
	NumMainNumbers   int32     `db:"num_main_numbers" json:"num_main_numbers"`
	NumBonusNumbers  *int32    `db:"num_bonus_numbers" json:"num_bonus_numbers"`
	MainNumberRange  int32     `db:"main_number_range" json:"main_number_range"`
	BonusNumberRange *int32    `db:"bonus_number_range" json:"bonus_number_range"`
	StartTime        time.Time `db:"start_time" json:"start_time"`
	EndTime          time.Time `db:"end_time" json:"end_time"`
	CreatorId        int32     `db:"creator_id" json:"creator_id"`
	Result           *string   `db:"result" json:"result"`
	Prize            *uint32   `db:"prize" json:"prize"`
	AutoCompute      bool      `db:"auto_compute" json:"auto_compute"`
	CreatedAt        time.Time `db:"created_at" json:"created_at"`
}

type DivisionDetail struct {
	Division      string  `json:"division"`
	DivisionPrize float32 `json:"won_prize"`
	UserCount     uint32  `json:"user_count"`
}
type GameAndUserChoice struct {
	Game            Game               `json:"game"`
	DivisionResult  []DivisionResult   `json:"division_result"`
	UserChoice      []UserChoiceResult `json:"user_choices"`
	TicketUsed      uint32             `json:"ticket_used"`
	DivisionDetails []DivisionDetail   `json:"division_details"`
}
type GamesAndUserChoices struct {
	Games []GameAndUserChoice `json:"games"`
}

type UserPrizeUpdate struct {
	UserId   int32   `json:"user_id"`
	WonPrize float32 `json:"won_prize"`
}

type DivisionUpdate struct {
	DivisionName string            `json:"division_name"`
	Users        []UserPrizeUpdate `json:"users"`
}

// GameType is an enum type for lottery game types
type GameType int

const (
	Lotto GameType = iota
	Ozlotto
	Powerball
	AmericanPowerball
)

// String converts the GameType enum to a string
func (g GameType) String() string {
	return [...]string{"lotto", "ozlotto", "powerball", "american_powerball"}[g]
}

func GameTypeToString(a int32) string {
	return [...]string{"lotto", "ozlotto", "powerball", "american_powerball"}[a]
}

// FromString converts a string to the GameType enum
func FromString(s string) GameType {
	switch s {
	case "lotto":
		return Lotto
	case "ozlotto":
		return Ozlotto
	case "powerball":
		return Powerball
	case "american_powerball":
		return AmericanPowerball
	default:
		return Lotto
	}
}

func FromStringToMainNumberRange(s int32) int32 {
	switch s {
	case 0:
		return 49
	case 1:
		return 45
	case 2:
		return 35
	case 3:
		return 69
	default:
		return 49
	}
}

func FromStringToNumberMainNumbers(s int32) int32 {
	switch s {
	case 0:
		return 6
	case 1:
		return 7
	case 2:
		return 7
	case 3:
		return 5
	default:
		return 6
	}
}
func FromStringToNumberBonusNumbers(s int32) *int32 {
	var bonusNumbers *int32
	switch s {
	case 0: // Lotto (no bonus number)
		bonusNumbers = nil
	case 1: // Ozlotto (no bonus number)
		bonusNumbers = nil
	case 2: // Powerball
		bonusNumbersValue := int32(1)
		bonusNumbers = &bonusNumbersValue
	case 3: // American Powerball
		bonusNumbersValue := int32(1)
		bonusNumbers = &bonusNumbersValue
	default: // Default to no bonus number
		bonusNumbers = nil
	}
	return bonusNumbers
}

func FromStringToBonusNumberRange(s int32) *int32 {
	var bonusNumberRange *int32
	switch s {
	case 0: // Lotto (no bonus number range)
		bonusNumberRange = nil
	case 1: // Ozlotto (no bonus number range)
		bonusNumberRange = nil
	case 2: // Powerball
		bonusNumberRangeValue := int32(20) // Powerball bonus number range is 1-26
		bonusNumberRange = &bonusNumberRangeValue
	case 3: // American Powerball
		bonusNumberRangeValue := int32(26) // American Powerball bonus number range is also 1-26
		bonusNumberRange = &bonusNumberRangeValue
	default: // Default to no bonus number range
		bonusNumberRange = nil
	}
	return bonusNumberRange
}

type Games struct {
	Games []Game `json:"game"`
	Total *int32 `json:"total"`
}

type PreviousGame struct {
	Game           Game             `json:"game"`
	DivisionResult []DivisionResult `json:"division_result"`
}

type PreviousGames struct {
	PreviousGames []PreviousGame `json:"games"`
	Total         *int32         `json:"total"`
}
