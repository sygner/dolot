package repository

import (
	"dolott_game/internal/models"
	"sync"
)

// // GetLottoDivisionUsers categorizes users into divisions based on the number of matches between
// // their chosen numbers and the winning numbers. It returns a map where the key is the match count
// // (7, 6, 5, etc.), and the value is a list of user choices with that match count.
// func (c *gameRepository) GetLottoDivisionUsers(winningNumbers []int32, userChoices []models.UserChoiceResult, targetMatchCounts []int) map[int][]models.UserChoiceResultDetail {
// 	// Initialize a map to store division users categorized by match count
// 	divisionUsersMap := make(map[int][]models.UserChoiceResultDetail)

// 	// Convert targetMatchCounts slice to a set for faster lookup
// 	targetMatchCountSet := make(map[int]struct{})
// 	for _, matchCount := range targetMatchCounts {
// 		targetMatchCountSet[matchCount] = struct{}{}
// 	}

// 	// Iterate over each user's choices
// 	for _, choice := range userChoices {
// 		// Loop through each set of chosen numbers for the user
// 		for _, chosenArray := range choice.ChosenNumbers {
// 			matchCount := getMatchCount(winningNumbers, chosenArray)

// 			// If the match count is in the target list, categorize the user choice
// 			if _, exists := targetMatchCountSet[int(matchCount)]; exists {
// 				divisionUsersMap[int(matchCount)] = append(divisionUsersMap[int(matchCount)], models.UserChoiceResultDetail{
// 					UserId:            choice.UserId,
// 					ChosenMainNumbers: chosenArray,
// 					MatchCount:        matchCount,
// 				})
// 			}
// 		}
// 	}
// 	return divisionUsersMap
// }

// GetLottoDivisionUsers categorizes users into divisions based on the number of matches between
// their chosen numbers and the winning numbers. It returns a slice of DivisionResult structs.
func (c *gameRepository) GetLottoDivisionUsers(winningNumbers []int32, userChoices []models.UserChoiceResult, targetMatchCounts []int) []models.DivisionResult {
	// Initialize a map to store division users categorized by match count
	divisionUsersMap := make(map[int][]models.UserChoiceResultDetail)
	mu := sync.Mutex{} // Mutex for synchronizing access to the map

	// Convert targetMatchCounts slice to a set for faster lookup
	targetMatchCountSet := make(map[int]struct{})
	for _, matchCount := range targetMatchCounts {
		targetMatchCountSet[matchCount] = struct{}{}
	}

	// Create a channel to handle results from Goroutines
	resultChan := make(chan struct {
		matchCount       int
		userChoiceDetail models.UserChoiceResultDetail
	})

	// Use WaitGroup to wait for all Goroutines to finish
	var wg sync.WaitGroup

	// Iterate over each user's choices
	for _, choice := range userChoices {
		wg.Add(1) // Increment the WaitGroup counter
		go func(choice models.UserChoiceResult) {
			defer wg.Done() // Decrement the counter when the Goroutine completes

			// Loop through each set of chosen numbers for the user
			for _, chosenArray := range choice.ChosenNumbers {
				matchCount := getMatchCount(winningNumbers, chosenArray)

				// If the match count is in the target list, categorize the user choice
				if _, exists := targetMatchCountSet[int(matchCount)]; exists {
					resultChan <- struct {
						matchCount       int
						userChoiceDetail models.UserChoiceResultDetail
					}{
						matchCount: int(matchCount),
						userChoiceDetail: models.UserChoiceResultDetail{
							UserId:            choice.UserId,
							ChosenMainNumbers: chosenArray,
							MatchCount:        matchCount,
						},
					}
				}
			}
		}(choice)
	}

	// Close the result channel once all Goroutines have finished
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	// Collect results from the channel
	for result := range resultChan {
		mu.Lock() // Lock the map before modifying it
		divisionUsersMap[result.matchCount] = append(divisionUsersMap[result.matchCount], result.userChoiceDetail)
		mu.Unlock() // Unlock the map after modification
	}

	// Convert the map to a slice of DivisionResult
	var divisionResults []models.DivisionResult
	for matchCount, userChoices := range divisionUsersMap {
		divisionResults = append(divisionResults, models.DivisionResult{
			MatchCount:  int32(matchCount),
			HasBonus:    false, // Lotto doesn't require bonus matches
			UserChoices: userChoices,
		})
	}

	return divisionResults
}

// GetPowerballDivisionUsersWithBonus checks user choices against winning numbers and a bonus number,
// categorizing them into divisions based on match counts. It supports checking for exact bonus number matches.
func (c *gameRepository) GetPowerballDivisionUsersWithBonus(winningNumbers []int32, winningBonusNumber int32, userChoices []models.UserChoiceResult, divisions []models.BonusDivision) []models.DivisionResult {
	// Initialize the map to store users by match count (key)
	divisionUsersMap := make(map[int32][]models.UserChoiceResultDetail)
	mu := sync.Mutex{} // Mutex for synchronizing access to the map

	// Create a channel to handle results from Goroutines
	resultChan := make(chan struct {
		matchCount       int32
		userChoiceDetail models.UserChoiceResultDetail
	})

	// Use WaitGroup to wait for all Goroutines to finish
	var wg sync.WaitGroup

	// Iterate over each user's choices
	for _, choice := range userChoices {
		wg.Add(1) // Increment the WaitGroup counter
		go func(choice models.UserChoiceResult) {
			defer wg.Done() // Decrement the counter when the Goroutine completes

			// Loop through each set of chosen main and bonus numbers
			for i, chosenMainNumbers := range choice.ChosenNumbers {
				if i >= len(choice.ChosenBonusNumber) {
					// If the chosen bonus number array is shorter than the chosen main numbers array, skip it
					continue
				}

				chosenBonusNumber := choice.ChosenBonusNumber[i]
				matchCount := getMatchCount(winningNumbers, chosenMainNumbers)

				// Check through all predefined bonus divisions
				for _, division := range divisions {
					if division.MatchCount == matchCount {
						if division.HasBonus {
							// If an exact bonus match is required, check both main and bonus numbers
							if chosenBonusNumber == winningBonusNumber {
								resultChan <- struct {
									matchCount       int32
									userChoiceDetail models.UserChoiceResultDetail
								}{
									matchCount: matchCount,
									userChoiceDetail: models.UserChoiceResultDetail{
										UserId:            choice.UserId,
										ChosenMainNumbers: chosenMainNumbers,
										ChosenBonusNumber: chosenBonusNumber,
										MatchCount:        matchCount,
									},
								}
							}
						} else {
							// Only main numbers need to match, no check for bonus number
							resultChan <- struct {
								matchCount       int32
								userChoiceDetail models.UserChoiceResultDetail
							}{
								matchCount: matchCount,
								userChoiceDetail: models.UserChoiceResultDetail{
									UserId:            choice.UserId,
									ChosenMainNumbers: chosenMainNumbers,
									ChosenBonusNumber: chosenBonusNumber,
									MatchCount:        matchCount,
								},
							}
						}
					}
				}
			}
		}(choice)
	}

	// Close the result channel once all Goroutines have finished
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	// Collect results from the channel
	for result := range resultChan {
		mu.Lock() // Lock the map before modifying it
		divisionUsersMap[result.matchCount] = append(divisionUsersMap[result.matchCount], result.userChoiceDetail)
		mu.Unlock() // Unlock the map after modification
	}

	// Convert the map to a slice of DivisionResult
	var divisionResults []models.DivisionResult
	for matchCount, userChoices := range divisionUsersMap {
		// Determine if the division requires a bonus number check
		hasBonus := false
		for _, division := range divisions {
			if division.MatchCount == matchCount && division.HasBonus {
				hasBonus = true
				break
			}
		}
		divisionResults = append(divisionResults, models.DivisionResult{
			MatchCount:  matchCount,
			HasBonus:    hasBonus,
			UserChoices: userChoices,
		})
	}

	return divisionResults
}

func getMatchCount(winningNumbers []int32, chosenNumbers []int32) int32 {
	matchCount := int32(0)
	numberSet := make(map[int32]struct{})

	// Add winning numbers to a set for quick lookups
	for _, winNum := range winningNumbers {
		numberSet[winNum] = struct{}{}
	}

	// Count how many of the user's numbers are in the winning set
	for _, userNum := range chosenNumbers {
		if _, exists := numberSet[userNum]; exists {
			matchCount++
		}
	}
	return matchCount
}

// GetPowerballDivisionUsersWithBonus checks user choices against winning numbers and a bonus number,
// categorizing them into divisions based on match counts. It supports checking for exact bonus number matches.
// func (c *gameRepository) GetPowerballDivisionUsersWithBonus(winningNumbers []int32, winningBonusNumber int32, userChoices []models.UserChoiceResult) map[int32][]models.UserChoiceResultDetail {
// 	// Initialize the map to store users by match count (key)
// 	divisionUsersMap := make(map[int32][]models.UserChoiceResultDetail)

// 	// Iterate over each user's choices
// 	for _, choice := range userChoices {
// 		// Loop through each set of chosen main and bonus numbers
// 		for i, chosenMainNumbers := range choice.ChosenNumbers {
// 			if i >= len(choice.ChosenBonusNumber) {
// 				// If the chosen bonus number array is shorter than the chosen main numbers array, skip it
// 				continue
// 			}

// 			chosenBonusNumber := choice.ChosenBonusNumber[i]
// 			matchCount := getMatchCount(winningNumbers, chosenMainNumbers)

// 			// Check through all predefined bonus divisions
// 			for _, division := range models.BonusDivisions {
// 				if division.MatchCount == matchCount {
// 					if division.HasBonus {
// 						// If an exact bonus match is required, check both main and bonus numbers
// 						if chosenBonusNumber == winningBonusNumber {
// 							for _, item := range divisionUsersMap[matchCount] {
// 								if item.UserId != choice.UserId && item.MatchCount != choice.UserId {
// 									divisionUsersMap[matchCount] = append(divisionUsersMap[matchCount], models.UserChoiceResultDetail{
// 										UserId:            choice.UserId,
// 										ChosenMainNumbers: chosenMainNumbers,
// 										ChosenBonusNumber: chosenBonusNumber, // Use a slice to match the model's structure
// 										MatchCount:        matchCount,
// 									})
// 								}
// 							}
// 							// Both main and bonus numbers match, add to the division

// 						}
// 					} else {
// 						// Only main numbers need to match, no check for bonus number
// 						divisionUsersMap[matchCount] = append(divisionUsersMap[matchCount], models.UserChoiceResultDetail{
// 							UserId:            choice.UserId,
// 							ChosenMainNumbers: chosenMainNumbers,
// 							ChosenBonusNumber: chosenBonusNumber,
// 							MatchCount:        matchCount,
// 						})
// 					}
// 					// Do not break here, allow both HasBonus true and false for the same MatchCount
// 				}
// 			}
// 		}
// 	}

// 	return divisionUsersMap
// }

// getMatchCount returns the number of matching numbers between the winning numbers and the user's chosen numbers.
