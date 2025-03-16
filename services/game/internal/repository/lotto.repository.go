package repository

import (
	"dolott_game/internal/models"
	"fmt"
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

//-----------------------------------------------
// GetLottoDivisionUsers categorizes users into divisions based on the number of matches between
// their chosen numbers and the winning numbers. It returns a slice of DivisionResult structs.
// func (c *gameRepository) GetLottoDivisionUsers(winningNumbers []int32, userChoices []models.UserChoiceResult, targetMatchCounts []int) []models.DivisionResult {
// 	// Initialize a map to store division users categorized by match count
// 	divisionUsersMap := make(map[int][]models.UserChoiceResultDetail)
// 	mu := sync.Mutex{} // Mutex for synchronizing access to the map

// 	// Convert targetMatchCounts slice to a set for faster lookup
// 	targetMatchCountSet := make(map[int]struct{})
// 	for _, matchCount := range targetMatchCounts {
// 		targetMatchCountSet[matchCount] = struct{}{}
// 	}

// 	// Create a channel to handle results from Goroutines
// 	resultChan := make(chan struct {
// 		matchCount       int
// 		userChoiceDetail models.UserChoiceResultDetail
// 	})

// 	// Use WaitGroup to wait for all Goroutines to finish
// 	var wg sync.WaitGroup

// 	// Iterate over each user's choices
// 	for _, choice := range userChoices {
// 		wg.Add(1) // Increment the WaitGroup counter
// 		go func(choice models.UserChoiceResult) {
// 			defer wg.Done() // Decrement the counter when the Goroutine completes

// 			// Loop through each set of chosen numbers for the user
// 			for _, chosenArray := range choice.ChosenNumbers {
// 				matchCount := getMatchCount(winningNumbers, chosenArray)

// 				// If the match count is in the target list, categorize the user choice
// 				if _, exists := targetMatchCountSet[int(matchCount)]; exists {
// 					resultChan <- struct {
// 						matchCount       int
// 						userChoiceDetail models.UserChoiceResultDetail
// 					}{
// 						matchCount: int(matchCount),
// 						userChoiceDetail: models.UserChoiceResultDetail{
// 							UserId:            choice.UserId,
// 							ChosenMainNumbers: chosenArray,
// 							MatchCount:        matchCount,
// 						},
// 					}
// 				}
// 			}
// 		}(choice)
// 	}

// 	// Close the result channel once all Goroutines have finished
// 	go func() {
// 		wg.Wait()
// 		close(resultChan)
// 	}()

// 	// Collect results from the channel
// 	for result := range resultChan {
// 		mu.Lock() // Lock the map before modifying it
// 		divisionUsersMap[result.matchCount] = append(divisionUsersMap[result.matchCount], result.userChoiceDetail)
// 		mu.Unlock() // Unlock the map after modification
// 	}

// 	// Convert the map to a slice of DivisionResult
// 	var divisionResults []models.DivisionResult
// 	for matchCount, userChoices := range divisionUsersMap {
// 		divisionResults = append(divisionResults, models.DivisionResult{
// 			MatchCount:  int32(matchCount),
// 			HasBonus:    false, // Lotto doesn't require bonus matches
// 			UserChoices: userChoices,
// 		})
// 	}

// 	return divisionResults
// }

// // GetPowerballDivisionUsersWithBonus checks user choices against winning numbers and a bonus number,
// // categorizing them into divisions based on match counts. It supports checking for exact bonus number matches.
// func (c *gameRepository) GetPowerballDivisionUsersWithBonus(winningNumbers []int32, winningBonusNumber int32, userChoices []models.UserChoiceResult, divisions []models.BonusDivision) []models.DivisionResult {
// 	// Initialize the map to store users by match count (key)
// 	divisionUsersMap := make(map[int32][]models.UserChoiceResultDetail)
// 	mu := sync.Mutex{} // Mutex for synchronizing access to the map

// 	// Create a channel to handle results from Goroutines
// 	resultChan := make(chan struct {
// 		matchCount       int32
// 		userChoiceDetail models.UserChoiceResultDetail
// 	})

// 	// Use WaitGroup to wait for all Goroutines to finish
// 	var wg sync.WaitGroup

// 	// Iterate over each user's choices
// 	for _, choice := range userChoices {
// 		wg.Add(1) // Increment the WaitGroup counter
// 		go func(choice models.UserChoiceResult) {
// 			defer wg.Done() // Decrement the counter when the Goroutine completes

// 			// Loop through each set of chosen main and bonus numbers
// 			for i, chosenMainNumbers := range choice.ChosenNumbers {
// 				if i >= len(choice.ChosenBonusNumber) {
// 					// If the chosen bonus number array is shorter than the chosen main numbers array, skip it
// 					continue
// 				}

// 				chosenBonusNumber := choice.ChosenBonusNumber[i]
// 				matchCount := getMatchCount(winningNumbers, chosenMainNumbers)

// 				// Check through all predefined bonus divisions
// 				for _, division := range divisions {
// 					if division.MatchCount == matchCount {
// 						if division.HasBonus {
// 							// If an exact bonus match is required, check both main and bonus numbers
// 							if chosenBonusNumber == winningBonusNumber {
// 								resultChan <- struct {
// 									matchCount       int32
// 									userChoiceDetail models.UserChoiceResultDetail
// 								}{
// 									matchCount: matchCount,
// 									userChoiceDetail: models.UserChoiceResultDetail{
// 										UserId:            choice.UserId,
// 										ChosenMainNumbers: chosenMainNumbers,
// 										ChosenBonusNumber: chosenBonusNumber,
// 										MatchCount:        matchCount,
// 									},
// 								}
// 							}
// 						} else {
// 							// Only main numbers need to match, no check for bonus number
// 							resultChan <- struct {
// 								matchCount       int32
// 								userChoiceDetail models.UserChoiceResultDetail
// 							}{
// 								matchCount: matchCount,
// 								userChoiceDetail: models.UserChoiceResultDetail{
// 									UserId:            choice.UserId,
// 									ChosenMainNumbers: chosenMainNumbers,
// 									ChosenBonusNumber: chosenBonusNumber,
// 									MatchCount:        matchCount,
// 								},
// 							}
// 						}
// 					}
// 				}
// 			}
// 		}(choice)
// 	}

// 	// Close the result channel once all Goroutines have finished
// 	go func() {
// 		wg.Wait()
// 		close(resultChan)
// 	}()

// 	// Collect results from the channel
// 	for result := range resultChan {
// 		mu.Lock() // Lock the map before modifying it
// 		divisionUsersMap[result.matchCount] = append(divisionUsersMap[result.matchCount], result.userChoiceDetail)
// 		mu.Unlock() // Unlock the map after modification
// 	}

// 	// Convert the map to a slice of DivisionResult
// 	var divisionResults []models.DivisionResult
// 	for matchCount, userChoices := range divisionUsersMap {
// 		// Determine if the division requires a bonus number check
// 		hasBonus := false
// 		for _, division := range divisions {
// 			if division.MatchCount == matchCount && division.HasBonus {
// 				hasBonus = true
// 				break
// 			}
// 		}
// 		divisionResults = append(divisionResults, models.DivisionResult{
// 			MatchCount:  matchCount,
// 			HasBonus:    hasBonus,
// 			UserChoices: userChoices,
// 		})
// 	}

// 	return divisionResults
// }

// // getMatchCount counts how many times the user’s numbers appear in the winningNumbers
// // without letting repeated user picks exceed the frequency in winningNumbers.
// func getMatchCount(winningNumbers []int32, chosenNumbers []int32) int32 {
// 	// Build a frequency map for the winning numbers
// 	freqMap := make(map[int32]int)
// 	for _, winNum := range winningNumbers {
// 		freqMap[winNum]++
// 	}

// 	var matchCount int32
// 	// For each chosen number, match it against freqMap if available
// 	for _, userNum := range chosenNumbers {
// 		if freqMap[userNum] > 0 {
// 			matchCount++
// 			freqMap[userNum]-- // Decrement so we don’t double-count
// 		}
// 	}

// 	return matchCount
// }
// ---------------------------------------

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

// GetLottoDivisionUsers categorizes users into divisions based on the number of matches between
// their chosen numbers and the winning numbers. It returns a slice of DivisionResult structs.
func (c *gameRepository) GetLottoDivisionUsers(
	winningNumbers []int32,
	userChoices []models.UserChoiceResult,
	targetMatchCounts []int,
) []models.DivisionResult {
	// Initialize a map: matchCount -> []UserChoiceResultDetail
	divisionUsersMap := make(map[int][]models.UserChoiceResultDetail)
	mu := sync.Mutex{}

	// For quick membership checks
	targetMatchCountSet := make(map[int]struct{})
	for _, matchCount := range targetMatchCounts {
		targetMatchCountSet[matchCount] = struct{}{}
	}

	// Channel to receive (matchCount, userChoiceDetail)
	resultChan := make(chan struct {
		matchCount       int
		userChoiceDetail models.UserChoiceResultDetail
	})

	var wg sync.WaitGroup

	// Concurrency: process each user’s choices
	for _, choice := range userChoices {
		wg.Add(1)
		go func(choice models.UserChoiceResult) {
			defer wg.Done()
			for _, chosenArray := range choice.ChosenNumbers {
				matchCount := getMatchCount(winningNumbers, chosenArray)

				if _, exists := targetMatchCountSet[int(matchCount)]; exists {
					// Send to channel for accumulation
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

	// Close channel when goroutines complete
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	// Accumulate results
	for res := range resultChan {
		mu.Lock()
		divisionUsersMap[res.matchCount] = append(
			divisionUsersMap[res.matchCount],
			res.userChoiceDetail,
		)
		mu.Unlock()
	}

	// Build the final divisions in the order of targetMatchCounts
	var divisionResults []models.DivisionResult
	for i, mc := range targetMatchCounts {
		if userDetailList, ok := divisionUsersMap[mc]; ok {
			divisionResults = append(divisionResults, models.DivisionResult{
				Division:    fmt.Sprintf("Division %d", i+1),
				MatchCount:  int32(mc),
				HasBonus:    false, // Lotto does not require a bonus
				UserChoices: userDetailList,
			})
		}
	}

	return divisionResults
}

// GetPowerballDivisionUsersWithBonus checks user choices against winning numbers and a bonus number,
// categorizing them into divisions based on match counts. It supports checking for exact bonus number matches.
func (c *gameRepository) GetPowerballDivisionUsersWithBonus(
	winningNumbers []int32,
	winningBonusNumber int32,
	userChoices []models.UserChoiceResult,
	divisions []models.BonusDivision,
) []models.DivisionResult {
	// matchCount -> []UserChoiceResultDetail
	divisionUsersMap := make(map[int32][]models.UserChoiceResultDetail)
	mu := sync.Mutex{}

	resultChan := make(chan struct {
		matchCount       int32
		userChoiceDetail models.UserChoiceResultDetail
	})

	var wg sync.WaitGroup
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
										HasBonus:          true,
									},
								}
								break // Break after sending the result to move on to the next set of chosen main numbers
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
										HasBonus:          false,
									},
								}
								break // Break after sending the result to move on to the next set of chosen main numbers
							}
							break
						}
						// Break from divisions loop if the correct division was found
						break
					}
				}
			}
		}(choice)
	}
	// Close channel after all goroutines finish
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	// Collect
	for res := range resultChan {
		mu.Lock()
		divisionUsersMap[res.matchCount] = append(
			divisionUsersMap[res.matchCount],
			res.userChoiceDetail,
		)
		mu.Unlock()
	}

	// Build results in the order of "divisions" based on the matchCount
	var divisionResults []models.DivisionResult

	// Step 1: Iterate through each division and process user choices
	for i, div := range divisions {
		if userList, ok := divisionUsersMap[div.MatchCount]; ok && len(userList) > 0 {

			// Iterate over the userList and check for matching divisions based on the HasBonus field
			for c := range userList {
				if div.HasBonus == userList[c].HasBonus {
					// Step 2: Check if the division already exists in the divisionResults
					existingDivisionIndex := -1
					for j, result := range divisionResults {
						if result.Division == fmt.Sprintf("Division %d", i+1) {
							// Found the existing division, now merge user choices
							existingDivisionIndex = j
							break
						}
					}

					if existingDivisionIndex != -1 {
						// If division exists, merge user choice into the existing division
						divisionResults[existingDivisionIndex].UserChoices = append(divisionResults[existingDivisionIndex].UserChoices, userList[c])
					} else {
						// If division does not exist, create a new division
						divisionResults = append(divisionResults, models.DivisionResult{
							Division:    fmt.Sprintf("Division %d", i+1),
							MatchCount:  div.MatchCount,
							HasBonus:    div.HasBonus,
							UserChoices: []models.UserChoiceResultDetail{userList[c]},
						})
					}
				}
			}
		}
	}

	// Step 3: Return the division results
	fmt.Println("Final Division Results:", divisionResults)
	return divisionResults
}

// getMatchCount counts how many times the user’s numbers appear in the winningNumbers
// without letting repeated user picks exceed the frequency in winningNumbers.
func getMatchCount(winningNumbers []int32, chosenNumbers []int32) int32 {
	// Build a frequency map for the winning numbers
	freqMap := make(map[int32]int)
	for _, winNum := range winningNumbers {
		freqMap[winNum]++
	}

	var matchCount int32
	// For each chosen number, match it against freqMap if available
	for _, userNum := range chosenNumbers {
		if freqMap[userNum] > 0 {
			matchCount++
			freqMap[userNum]-- // Decrement so we don’t double-count
		}
	}

	return matchCount
}
