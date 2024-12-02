package models

import "time"

type UserChoiceResult struct {
	UserId            int32     `json:"user_id"`
	ChosenNumbers     [][]int32 `json:"chosen_numbers"` // Assuming it's a 2D array
	ChosenBonusNumber []int32   `json:"chosen_bonus_number"`
	MatchCounts       []int32   `json:"match_counts"` // New field for match counts
}

type UserChoice struct {
	Id                 string    `json:"id"`
	UserId             int32     `json:"user_id"`
	GameId             string    `json:"game_id"`
	ChosenMainNumbers  [][]int32 `json:"chosen_main_numbers"`
	ChosenBonusNumbers [][]int32 `json:"chosen_bonus_numbers"`
	CreatedAt          time.Time `json:"created_at"`
}

type UserChoiceResultDetail struct {
	UserId            int32   `json:"user_id"`
	ChosenMainNumbers []int32 `json:"chosen_main_numbers"`
	ChosenBonusNumber int32   `json:"chosen_bonus_number"`
	MatchCount        int32   `json:"match_count"`
}

type UserChoices struct {
	UserChoices []UserChoice `json:"user_choice"`
	Total       *int32       `json:"total"`
}

// DivisionResult represents the result of categorizing user choices into divisions based on match counts.
type DivisionResult struct {
	MatchCount  int32                    `json:"match_count"`  // The number of matching numbers
	HasBonus    bool                     `json:"has_bonus"`    // Whether the division requires a matching bonus number
	UserChoices []UserChoiceResultDetail `json:"user_choices"` // A list of user choices that belong to this division
}

type BonusDivision struct {
	MatchCount int32 `json:"match_count"`
	HasBonus   bool  `json:"has_bonus"`
}

var (
	PowerBallDivisions = []BonusDivision{
		{MatchCount: 7, HasBonus: true},
		{MatchCount: 7, HasBonus: false},
		{MatchCount: 6, HasBonus: true},
		{MatchCount: 6, HasBonus: false},
		{MatchCount: 5, HasBonus: true},
		{MatchCount: 5, HasBonus: false},
		{MatchCount: 4, HasBonus: true},
		{MatchCount: 3, HasBonus: true},
		{MatchCount: 2, HasBonus: true},
		{MatchCount: 2, HasBonus: false},
	}

	AmericanPowerballDivisions = []BonusDivision{
		{MatchCount: 5, HasBonus: true},
		{MatchCount: 5, HasBonus: false},
		{MatchCount: 4, HasBonus: true},
		{MatchCount: 4, HasBonus: false},
		{MatchCount: 3, HasBonus: true},
		{MatchCount: 3, HasBonus: false},
		{MatchCount: 2, HasBonus: true},
		{MatchCount: 1, HasBonus: true},
		{MatchCount: 0, HasBonus: true},
	}
)
