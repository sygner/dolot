package services

import (
	"context"
	"dolott_user_gw_http/internal/constants"
	"dolott_user_gw_http/internal/models"
	"dolott_user_gw_http/internal/types"
	pb "dolott_user_gw_http/proto/api/game"
	profile_pb "dolott_user_gw_http/proto/api/profile"
	ticket_pb "dolott_user_gw_http/proto/api/ticket"
	wallet_pb "dolott_user_gw_http/proto/api/wallet"
	"log"
	"strings"

	"fmt"
	"math"
)

type (
	GameService interface {
		GetGameByGameId(string) (*models.Game, *types.Error)
		AddGame(*models.AddGameDTO) (*models.Game, *types.Error)
		GetNextGamesByGameType(int32, int32) (*models.Games, *types.Error)
		DeleteGameByGameId(string) *types.Error
		CheckGameExistsGameId(string) *types.Error
		GetGamesByCreatorId(int32, *models.Pagination) (*models.Games, *types.Error)
		AddResultByGameId(string, string) ([]models.DivisionResult, *types.Error)
		GetAllNextGames() (*models.Games, *types.Error)
		GetAllPreviousGames(*models.Pagination) (*models.Games, *types.Error)
		GetAllGames(*models.Pagination) (*models.Games, *types.Error)
		GetAllGameTypes() ([]models.GameTypeDetail, *types.Error)
		UpdateGameTypeDetail(int32, *string, int32, int32, bool) ([]models.GameTypeDetail, *types.Error)
		GetAllUserPreviousGames(int32, *models.Pagination) (*models.Games, *types.Error)
		GetAllUserPreviousGamesByGameType(int32, string, *models.Pagination) (*models.Games, *types.Error)
		GetAllUserChoiceDivisionsByGameId(int32, string) ([]models.DivisionResult, *types.Error)
		GetAllUsersChoiceDivisionsByGameId(string) ([]models.DivisionResult, *types.Error)
		UpdateGamePrizeByGameId(string, *uint32, bool) *types.Error
		GetUserGamesByTimesAndGameTypes(int32, *string, string, string) ([]models.GameAndUserChoice, *types.Error)
		TransactionForWinner([]models.DivisionResult, string)
	}
	gameService struct {
		gameClient    pb.GameServiceClient
		winnerClient  pb.WinnerServiceClient
		ticketClient  ticket_pb.TicketServiceClient
		walletClient  wallet_pb.WalletServiceClient
		profileClient profile_pb.ProfileServiceClient
		serverAddress string
	}
)

func NewGameService(gameClient pb.GameServiceClient, ticketClient ticket_pb.TicketServiceClient, walletClient wallet_pb.WalletServiceClient, winnerClient pb.WinnerServiceClient, profileClient profile_pb.ProfileServiceClient, serverAddress string) GameService {
	return &gameService{
		gameClient:    gameClient,
		winnerClient:  winnerClient,
		ticketClient:  ticketClient,
		walletClient:  walletClient,
		profileClient: profileClient,
		serverAddress: serverAddress,
	}
}

func (c *gameService) GetGameByGameId(gameId string) (*models.Game, *types.Error) {
	res, err := c.gameClient.GetGameByGameId(context.Background(), &pb.GameId{GameId: gameId})
	if err != nil {
		return nil, types.ExtractGRPCErrDetails(err)
	}
	return toGameProto(res)
}

func (c *gameService) AddGame(data *models.AddGameDTO) (*models.Game, *types.Error) {
	res, err := c.gameClient.AddGame(context.Background(), &pb.AddGameRequest{
		Name:        data.Name,
		GameType:    pb.GameType(data.GameTypeInt),
		StartTime:   data.StartTime,
		EndTime:     data.EndTime,
		Prize:       data.Prize,
		AutoCompute: data.AutoCompute,
		CreatorId:   data.CreatorId,
	})
	if err != nil {
		return nil, types.ExtractGRPCErrDetails(err)
	}
	return toGameProto(res)
}

func (c *gameService) GetNextGamesByGameType(gameType int32, limit int32) (*models.Games, *types.Error) {
	res, err := c.gameClient.GetNextGamesByGameType(context.Background(), &pb.GameTypeRequest{GameType: pb.GameType(gameType), Limit: limit})
	if err != nil {
		return nil, types.ExtractGRPCErrDetails(err)
	}
	return toGamesProto(res)
}

func (c *gameService) GetAllNextGames() (*models.Games, *types.Error) {
	res, err := c.gameClient.GetAllNextGames(context.Background(), &pb.Empty{})
	if err != nil {
		return nil, types.ExtractGRPCErrDetails(err)
	}
	return toGamesProto(res)
}

func (c *gameService) GetAllPreviousGames(pagination *models.Pagination) (*models.Games, *types.Error) {
	res, err := c.gameClient.GetAllPreviousGames(context.Background(), &pb.Pagination{
		Offset:   pagination.Offset,
		Limit:    pagination.Limit,
		GetTotal: pagination.Total,
	})
	if err != nil {
		return nil, types.ExtractGRPCErrDetails(err)
	}
	return toGamesProto(res)
}

func (c *gameService) GetAllGames(pagination *models.Pagination) (*models.Games, *types.Error) {
	res, err := c.gameClient.GetAllGames(context.Background(), &pb.Pagination{
		Offset:   pagination.Offset,
		Limit:    pagination.Limit,
		GetTotal: pagination.Total,
	})
	if err != nil {
		return nil, types.ExtractGRPCErrDetails(err)
	}
	return toGamesProto(res)
}

func (c *gameService) GetAllGameTypes() ([]models.GameTypeDetail, *types.Error) {
	res, err := c.gameClient.GetAllGameTypes(context.Background(), &pb.Empty{})
	if err != nil {
		return nil, types.ExtractGRPCErrDetails(err)
	}

	gameTypeDetails := make([]models.GameTypeDetail, 0)

	for _, gameType := range res.GameTypes {
		gameTypeDetails = append(gameTypeDetails, models.GameTypeDetail{
			Id:          gameType.Id,
			Name:        gameType.Name,
			Description: gameType.Description,
			TypeName:    gameType.TypeName,
			DayName:     gameType.DayName,
			PicturePath: c.serverAddress + "/api/game/dl?path=" + "svg/static/" + gameType.TypeName + ".svg",
			PrizeReward: gameType.PrizeReward,
			TokenBurn:   gameType.TokenBurn,
		})
	}
	return gameTypeDetails, nil
}

func (c *gameService) DeleteGameByGameId(gameId string) *types.Error {
	_, err := c.gameClient.DeleteGameByGameId(context.Background(), &pb.GameId{GameId: gameId})
	if err != nil {
		return types.ExtractGRPCErrDetails(err)
	}
	return nil
}

func (c *gameService) CheckGameExistsGameId(gameId string) *types.Error {
	_, err := c.gameClient.CheckGameExistsByGameId(context.Background(), &pb.GameId{GameId: gameId})
	if err != nil {
		return types.ExtractGRPCErrDetails(err)
	}
	return nil
}

func (c *gameService) GetGamesByCreatorId(userId int32, pagination *models.Pagination) (*models.Games, *types.Error) {
	res, err := c.gameClient.GetGamesByCreatorId(context.Background(), &pb.GetGamesByCreatorIdRequest{
		CreatorId: userId,
		Pagination: &pb.Pagination{
			Offset:   pagination.Offset,
			Limit:    pagination.Limit,
			GetTotal: pagination.Total}})
	if err != nil {
		return nil, types.ExtractGRPCErrDetails(err)
	}

	return toGamesProto(res)
}

func (c *gameService) AddResultByGameId(gameId string, result string) ([]models.DivisionResult, *types.Error) {
	res, err := c.gameClient.AddResultByGameId(context.Background(), &pb.AddResultByGameIdRequest{
		GameId: gameId,
		Result: result,
	})
	if err != nil {
		return nil, types.ExtractGRPCErrDetails(err)
	}

	// Enqueue safely in a goroutine, but handle errors
	go func() {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Recovered from panic in job enqueue: %v", r)
			}
		}()

		JOB_QUEUE.Enqueue(Job{
			Priority: Normal,
			Type:     LotteryJob,
			GameID:   gameId,
			Results:  toDivisionResultsProto(res.DivisionResults),
		})
	}()

	return toDivisionResultsProto(res.DivisionResults), nil
}

func (c *gameService) UpdateGameTypeDetail(gameType int32, dayName *string, prizeReward int32, tokenBurn int32, autoCompute bool) ([]models.GameTypeDetail, *types.Error) {
	res, err := c.gameClient.ChangeGameDetailCalculation(context.Background(), &pb.ChangeGameDetailCalculationRequest{
		GameType:    pb.GameType(gameType),
		DayName:     dayName,
		PrizeReward: prizeReward,
		TokenBurn:   tokenBurn,
		AutoCompute: autoCompute,
	})
	if err != nil {
		return nil, types.ExtractGRPCErrDetails(err)
	}
	gameTypeDetails := make([]models.GameTypeDetail, 0)

	for _, gameType := range res.GameTypes {
		gameTypeDetails = append(gameTypeDetails, models.GameTypeDetail{
			Id:          gameType.Id,
			Name:        gameType.Name,
			Description: gameType.Description,
			TypeName:    gameType.TypeName,
			DayName:     gameType.DayName,
			PicturePath: c.serverAddress + "/api/game/dl?path=" + "svg/static/" + gameType.TypeName + ".svg",
			PrizeReward: gameType.PrizeReward,
			TokenBurn:   gameType.TokenBurn,
		})
	}
	return gameTypeDetails, nil
}

func (c *gameService) GetAllUserPreviousGames(userId int32, data *models.Pagination) (*models.Games, *types.Error) {
	res, err := c.gameClient.GetAllUserPreviousGames(context.Background(), &pb.GetAllUserPreviousGamesRequest{
		UserId: userId,
		Pagination: &pb.Pagination{
			Offset:   data.Offset,
			Limit:    data.Limit,
			GetTotal: data.Total,
		},
	})

	if err != nil {
		return nil, types.ExtractGRPCErrDetails(err)
	}

	return toGamesProto(res)
}

func (c *gameService) GetAllUserPreviousGamesByGameType(userId int32, gameType string, data *models.Pagination) (*models.Games, *types.Error) {
	res, err := c.gameClient.GetAllUserPreviousGamesByGameType(context.Background(), &pb.GetAllUserPreviousGamesByGameTypeRequest{
		UserId:   userId,
		GameType: gameType,
		Pagination: &pb.Pagination{
			Offset:   data.Offset,
			Limit:    data.Limit,
			GetTotal: data.Total,
		},
	})

	if err != nil {
		return nil, types.ExtractGRPCErrDetails(err)
	}

	return toGamesProto(res)
}

func (c *gameService) GetAllUserChoiceDivisionsByGameId(userId int32, gameId string) ([]models.DivisionResult, *types.Error) {
	res, err := c.gameClient.GetAllUserChoiceDivisionsByGameId(context.Background(), &pb.GetAllUserChoiceDivisionsByGameIdRequest{
		UserId: userId,
		GameId: gameId,
	})

	if err != nil {
		return nil, types.ExtractGRPCErrDetails(err)
	}

	return toDivisionResultsProto(res.DivisionResults), nil
}

func (c *gameService) GetAllUsersChoiceDivisionsByGameId(gameId string) ([]models.DivisionResult, *types.Error) {
	res, err := c.gameClient.GetAllUsersChoiceDivisionsByGameId(context.Background(), &pb.GameId{
		GameId: gameId,
	})

	if err != nil {
		return nil, types.ExtractGRPCErrDetails(err)
	}

	return toDivisionResultsProto(res.DivisionResults), nil
}

func (c *gameService) UpdateGamePrizeByGameId(gameId string, prize *uint32, autoCompute bool) *types.Error {
	_, err := c.gameClient.UpdateGamePrizeByGameId(context.Background(), &pb.UpdateGamePrizeByGameIdRequest{
		GameId:      gameId,
		Prize:       prize,
		AutoCompute: autoCompute,
	})

	if err != nil {
		return types.ExtractGRPCErrDetails(err)
	}

	return nil
}

func (c *gameService) GetUserGamesByTimesAndGameTypes(userId int32, gameType *string, startTime, endTime string) ([]models.GameAndUserChoice, *types.Error) {
	res, err := c.gameClient.GetUserGamesByTimesAndGameTypes(context.Background(), &pb.GetUserGamesByTimesAndGameTypesRequest{
		UserId:    userId,
		GameType:  gameType,
		StartTime: startTime,
		EndTime:   endTime,
	})
	if err != nil {
		return nil, types.ExtractGRPCErrDetails(err)
	}

	games := make([]models.GameAndUserChoice, 0, len(res.Games))

	for _, game := range res.Games {
		gm, err := toGameProto(game.Game)
		if err != nil {
			return nil, err
		}

		gameResult := models.GameAndUserChoice{
			Game:       *gm,
			Won:        false,
			TicketUsed: game.TicketUsed,
		}

		if len(game.DivisionDetails) > 0 {
			divisionDetails := make([]models.DivisionDetail, len(game.DivisionDetails))
			for i, dv := range game.DivisionDetails {
				divisionDetails[i] = models.DivisionDetail{
					Division:      dv.Division,
					UserCount:     dv.UserCount,
					DivisionPrize: dv.DivisionPrize,
				}
			}
			gameResult.DivisionDetails = &divisionDetails
		}

		if len(game.DivisionResults.DivisionResults) > 0 {
			gameResult.Won = true
			divisionResults := make([]models.DivisionResult, len(game.DivisionResults.DivisionResults))
			for i, c := range game.DivisionResults.DivisionResults {
				userChoices := make([]models.UserChoiceResultDetail, len(c.UserChoice))
				for j, uc := range c.UserChoice {
					userChoices[j] = models.UserChoiceResultDetail{
						UserId:              uc.UserId,
						ChosenMainNumbers:   uc.ChosenMainNumber,
						ChosenBonusNumber:   uc.ChosenBonusNumber,
						MainAndBonusNumbers: formatChosenNumbers(int32(game.Game.GameType), uc.ChosenMainNumber, uc.ChosenBonusNumber),
						BoughtPrice:         uc.BoughtPrice,
					}
				}
				divisionResults[i] = models.DivisionResult{
					MatchCount:  c.MatchCount,
					HasBonus:    c.HasBonus,
					Division:    c.Division,
					UserChoices: userChoices,
				}
			}
			gameResult.DivisionResult = &divisionResults
		}

		if len(game.UserChoices) > 0 {
			userChoices := make([]models.UserChoiceResult, len(game.UserChoices))
			for i, uc := range game.UserChoices {
				formattedChoices := make([]string, len(uc.ChosenMainNumbers))
				for j, c := range uc.ChosenMainNumbers {
					formattedChoices[j] = formatChosenNumbers(int32(game.Game.GameType), c.ChosenMainNumbers, getBonusNumber(uc.ChosenBonusNumber, j))
				}
				userChoices[i] = models.UserChoiceResult{
					UserId:              uc.UserId,
					MainAndBonusNumbers: formattedChoices,
					BoughtPrice:         uc.BoughtPrice,
				}
			}
			gameResult.UserChoice = &userChoices
		}

		games = append(games, gameResult)
	}

	return games, nil
}

func getBonusNumber(bonusNumbers []int32, index int) int32 {
	if len(bonusNumbers) > index {
		return bonusNumbers[index]
	}
	return 0
}

func formatChosenNumbers(gameType int32, mainNumbers []int32, bonusNumber int32) string {
	mainStr := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(mainNumbers)), ","), "[]")
	if gameType == 2 || gameType == 3 {
		return fmt.Sprintf("%s+%d", mainStr, bonusNumber)
	}
	return mainStr
}

func (c *gameService) TransactionForWinner(data []models.DivisionResult, gameId string) {
	// wlt, err := c.walletClient.GetWalletsByUserId(context.Background(), &wallet_pb.UserId{
	// 	UserId: 0,
	// })
	// if err != nil || len(wlt.Wallets) == 0 {
	// 	fmt.Println("Wallet Not Found")
	// 	return
	// }

	luncPrice, rerr := constants.GetLUNCPriceCoinPaprika()
	if rerr != nil {
		return
	}

	// Fetch game details
	gameRes, err := c.gameClient.GetGameByGameId(context.Background(), &pb.GameId{GameId: gameId})
	if err != nil || gameRes.Prize == nil {
		fmt.Println("Error fetching game data or prize is nil")
		return
	}

	gameData, rerr := toGameProto(gameRes)
	if rerr != nil {
		return
	}

	jackpot := float64(*gameData.Prize)

	// Determine division payout structure
	var divisionMap map[string]float64
	gameType := 0
	switch gameData.GameType {
	case "LOTTO":
		divisionMap = LottoWinnerDivisions
		gameType = 0
	case "OZLOTTO":
		divisionMap = OsLottoWinnerDivisions
		gameType = 1
	case "POWERBALL":
		divisionMap = PowerballWinnerDivisions
		gameType = 2
	case "AMERICAN_POWERBALL":
		divisionMap = AmericanPowerballWinnerDivisions
		gameType = 3

	}

	lastWinnerGameRes, err := c.winnerClient.GetLastWinnersByGameType(context.Background(), &pb.GameTypeRequest{
		GameType: pb.GameType(gameType),
		Limit:    0,
	})
	if err != nil {
		return
	}
	if !lastWinnerGameRes.Jackpot {
		if gameType != 3 {
			jackpot += float64(lastWinnerGameRes.Prize)
		}
	}

	// Map to store total prizes per user
	userPrizeMap := make(map[int32]float64)
	userIds := make([]int32, 0)
	divisionUpdateWinners := make([]*pb.DivisionUpdate, 0)

	// Process each division result
	for _, dr := range data {
		updateUserWonPrizeWinners := make([]*pb.UserPrizeUpdate, 0)
		divisionPercentage, exists := divisionMap[dr.Division]
		if !exists || len(dr.UserChoices) == 0 {
			continue
		}

		// Use a temporary jackpot value that can be modified only for Division 1
		currentJackpot := jackpot

		// Only add the last winner's prize if it's Division 1 and the last game had no jackpot
		if dr.Division == "Division 1" {
			if !lastWinnerGameRes.Jackpot {
				if gameType != 3 {
					currentJackpot += float64(lastWinnerGameRes.Prize)
				}
			}
		}

		// Calculate total division prize using the modified or original jackpot
		divisionTotalPrize := currentJackpot * divisionPercentage
		winnerCount := float64(len(dr.UserChoices)) // Total winners in this division

		// Ensure divisionTotalPrize is split among all winners
		if winnerCount > 0 {
			prizePerWinner := divisionTotalPrize / winnerCount

			for _, uc := range dr.UserChoices {
				if _, exists := userPrizeMap[uc.UserId]; !exists {
					userIds = append(userIds, uc.UserId)
				}
				updateUserWonPrizeWinners = append(updateUserWonPrizeWinners, &pb.UserPrizeUpdate{
					UserId:   uc.UserId,
					WonPrize: float32(prizePerWinner * luncPrice),
				})
				fmt.Println("Share of Winning UserId:", uc.UserId, "Prize:", userPrizeMap[uc.UserId], prizePerWinner, "Division:", dr.Division, "Map of Division:", divisionMap[dr.Division], "JackPot:", currentJackpot, "Percentage:", divisionTotalPrize)
				userPrizeMap[uc.UserId] += prizePerWinner // Add user's share of winnings
			}
			divisionUpdateWinners = append(divisionUpdateWinners, &pb.DivisionUpdate{
				DivisionName: dr.Division,
				Users:        updateUserWonPrizeWinners,
			})
		}
	}

	fmt.Println("Final User Prize Map:", userPrizeMap, "Full Game Prize:", *gameData.Prize)

	// Fetch wallets for users
	res, err := c.walletClient.GetWalletsByUserIdsAndCoinId(context.Background(), &wallet_pb.GetWalletsByUserIdsAndCoinIdRequest{
		UserIds: userIds,
		CoinId:  1,
	})
	if err != nil {
		fmt.Println("Error fetching wallets:", err)
		return
	}

	walletMap := make(map[int32]*wallet_pb.Wallet)
	for _, wallet := range res.Wallets {
		walletMap[wallet.UserId] = wallet
	}
	paidPrize := float64(0)
	// Process payments
	for userId, totalPrize := range userPrizeMap {
		wallet, found := walletMap[userId]
		if !found {
			fmt.Printf("No wallet found for user %d\n", userId)
			continue
		}
		c.profileClient.ChangeImpressionAndDCoin(context.Background(), &profile_pb.ChangeImpressionAndDCreditRequest{
			UserId:     userId,
			Impression: int32(math.Round(totalPrize)),
			DCoin:      int32(math.Round(totalPrize)),
		})
		fmt.Printf("Paying User %d Prize: %.2f\n", userId, totalPrize)
		trx, err := c.walletClient.AddTransaction(context.Background(), &wallet_pb.AddTransactionRequest{
			FromWalletUserId: constants.MAIN_LUNC_USER_WALLET_ID,
			FromWalletId:     constants.MAIN_LUNC_WALLET_ID,
			ToWalletAddress:  wallet.Address,
			Amount:           math.Round(totalPrize), // Round amount for precision
			CoinId:           1,
		})
		if err != nil {
			fmt.Printf("Transaction failed for user %d: %v\n", userId, err)
		} else {
			paidPrize += float64(totalPrize)
			fmt.Println("Transaction Successful:", trx)
		}
	}
	c.winnerClient.UpdateTotalPaid(context.Background(), &pb.UpdateTotalPaidRequest{
		GameId:    gameId,
		TotalPaid: fmt.Sprintf("%.2f", paidPrize),
	})
	fmt.Println(" ********** \n ", divisionUpdateWinners, "\n ******")
	if _, err := c.gameClient.UpdateUserGameDivisionPrize(context.Background(), &pb.UpdateUserGameDivisionPrizeRequest{
		DivisionUpdates: divisionUpdateWinners,
		GameId:          gameId,
	}); err != nil {
		fmt.Println(err)
	}
	fmt.Println("All transactions processed successfully.")
}

// // Process each division result
// for _, dr := range data {
// 	divisionPercentage, exists := divisionMap[dr.Division]
// 	if !exists || len(dr.UserChoices) == 0 {
// 		continue
// 	}
// 	if dr.Division == "Division 1" {
// 		if !lastWinnerGameRes.Jackpot {
// 			if gameType != 3 {
// 				jackpot += float64(lastWinnerGameRes.Prize)
// 			}
// 		}
// 	}

// 	// Calculate total division prize
// 	divisionTotalPrize := jackpot * divisionPercentage
// 	winnerCount := float64(len(dr.UserChoices)) // Total winners in this division

// 	// Ensure divisionTotalPrize is split among all winners
// 	if winnerCount > 0 {
// 		prizePerWinner := divisionTotalPrize / winnerCount
// 		for _, uc := range dr.UserChoices {
// 			if _, exists := userPrizeMap[uc.UserId]; !exists {
// 				userIds = append(userIds, uc.UserId)
// 			}
// 			fmt.Println("Share of Winning UserId:", uc.UserId, "Prize:", userPrizeMap[uc.UserId], prizePerWinner, "Division:", dr.Division, "Map of Division:", divisionMap[dr.Division], "JackPot:", jackpot, "Percentage:", divisionTotalPrize)
// 			userPrizeMap[uc.UserId] += prizePerWinner // Add user's share of winnings
// 		}
// 	}
// }
