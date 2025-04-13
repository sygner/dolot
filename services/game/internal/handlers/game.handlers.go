package handlers

import (
	"context"
	"dolott_game/internal/models"
	"dolott_game/internal/services"
	"dolott_game/internal/types"
	pb "dolott_game/proto/api"
	"fmt"
	"strings"
)

type GameHandler struct {
	pb.UnimplementedGameServiceServer
	gameService services.GameServices
}

func NewGameHandler(gameService services.GameServices) *GameHandler {
	return &GameHandler{
		gameService: gameService,
	}
}

func (c *GameHandler) GetGameByGameId(ctx context.Context, request *pb.GameId) (*pb.Game, error) {
	res, err := c.gameService.GetGameByGameId(request.GameId)
	if err != nil {
		return nil, err.ErrorToGRPCStatus()
	}

	return toGameProto(res), nil
}

func (c *GameHandler) AddGame(ctx context.Context, request *pb.AddGameRequest) (*pb.Game, error) {
	// Helper function to parse time and handle errors

	startTime, err := parseTime(request.StartTime, "failed to convert the start time, wrong format #4501")
	if err != nil {
		return nil, err
	}

	endTime, err := parseTime(request.EndTime, "failed to convert the end time, wrong format #4502")
	if err != nil {
		return nil, err
	}

	data := models.AddGameDTO{
		Name:        request.Name,
		StartTime:   startTime,
		EndTime:     endTime,
		GameTypeInt: int32(request.GameType),
		CreatorId:   request.CreatorId,
		Prize:       request.Prize,
		AutoCompute: request.AutoCompute,
	}

	res, rerr := c.gameService.AddGame(&data)
	if rerr != nil {
		return nil, rerr.ErrorToGRPCStatus()
	}

	// Create the response in a more concise way
	return toGameProto(res), nil
}

func (c *GameHandler) GetNextGamesByGameType(ctx context.Context, request *pb.GameTypeRequest) (*pb.Games, error) {
	res, err := c.gameService.GetNextGamesByGameType(int32(request.GameType), request.Limit)
	if err != nil {
		return nil, err.ErrorToGRPCStatus()
	}
	return toGamesProto(&models.Games{Games: res, Total: nil}), nil
}

func (c *GameHandler) GetAllNextGames(ctx context.Context, request *pb.Empty) (*pb.Games, error) {
	res, err := c.gameService.GetAllNextGames()
	if err != nil {
		return nil, err.ErrorToGRPCStatus()
	}
	return toGamesProto(&models.Games{Games: res, Total: nil}), nil
}

func (c *GameHandler) GetAllPreviousGames(ctx context.Context, request *pb.Pagination) (*pb.Games, error) {
	data := models.Pagination{
		Offset: request.Offset,
		Limit:  request.Limit,
		Total:  request.GetTotal,
	}
	res, err := c.gameService.GetAllPreviousGames(&data)
	if err != nil {
		return nil, err.ErrorToGRPCStatus()
	}
	return toGamesProto(res), nil
}

func (c *GameHandler) GetAllGames(ctx context.Context, request *pb.Pagination) (*pb.Games, error) {
	data := models.Pagination{
		Offset: request.Offset,
		Limit:  request.Limit,
		Total:  request.GetTotal,
	}
	res, err := c.gameService.GetAllGames(&data)
	if err != nil {
		return nil, err.ErrorToGRPCStatus()
	}
	return toGamesProto(res), nil
}

func (c *GameHandler) DeleteGameByGameId(ctx context.Context, request *pb.GameId) (*pb.Empty, error) {
	err := c.gameService.DeleteGameByGameId(request.GameId)
	if err != nil {
		return nil, err.ErrorToGRPCStatus()
	}
	return &pb.Empty{}, nil
}

func (c *GameHandler) CheckGameExistsByGameId(ctx context.Context, request *pb.GameId) (*pb.Empty, error) {
	res, err := c.gameService.CheckGameExistsById(request.GameId)
	if err != nil {
		return nil, err.ErrorToGRPCStatus()
	}
	if res {
		return &pb.Empty{}, nil
	}
	return nil, types.NewNotFoundError("no game found #4503").ErrorToGRPCStatus()
}

func (c *GameHandler) GetGamesByCreatorId(ctx context.Context, request *pb.GetGamesByCreatorIdRequest) (*pb.Games, error) {
	data := models.Pagination{
		Offset: request.Pagination.Offset,
		Limit:  request.Pagination.Limit,
		Total:  request.Pagination.GetTotal,
	}

	res, err := c.gameService.GetGamesByCreatorId(request.CreatorId, &data)
	if err != nil {
		return nil, err.ErrorToGRPCStatus()
	}

	return toGamesProto(res), nil
}

func (c *GameHandler) AddResultByGameId(ctx context.Context, request *pb.AddResultByGameIdRequest) (*pb.DivisionResults, error) {
	exres, err := c.gameService.AddResultByGameId(request.GameId, request.Result)
	if err != nil {
		return nil, err.ErrorToGRPCStatus()
	}
	return toDivisionResultsProto(exres), nil
}

func (c *GameHandler) GetAllGameTypes(ctx context.Context, request *pb.Empty) (*pb.GameTypes, error) {
	exres, err := c.gameService.GetGameTypes()
	if err != nil {
		return nil, err.ErrorToGRPCStatus()
	}
	gameTypesDetails := make([]*pb.GameTypeDetails, 0)

	for _, gameType := range exres {
		gameTypesDetails = append(gameTypesDetails, &pb.GameTypeDetails{
			Id:          gameType.Id,
			Name:        gameType.Name,
			Description: gameType.Description,
			TypeName:    gameType.TypeName,
			DayName:     gameType.DayName,
			PrizeReward: gameType.PrizeReward,
			TokenBurn:   gameType.TokenBurn,
		})
	}
	return &pb.GameTypes{GameTypes: gameTypesDetails}, nil
}

func (c *GameHandler) ChangeGameDetailCalculation(ctx context.Context, request *pb.ChangeGameDetailCalculationRequest) (*pb.GameTypes, error) {
	var gameType int32
	switch request.GameType {
	case 0:
		gameType = 1
	case 1:
		gameType = 2
	case 2:
		gameType = 3
	case 3:
		gameType = 4
	}
	fmt.Println(gameType, request.PrizeReward, request.TokenBurn, request.AutoCompute)
	err := c.gameService.UpdateGameTypeDetail(gameType, request.DayName, request.PrizeReward, request.TokenBurn, request.AutoCompute)
	if err != nil {
		return nil, err.ErrorToGRPCStatus()
	}
	exres, err := c.gameService.GetGameTypes()
	if err != nil {
		return nil, err.ErrorToGRPCStatus()
	}
	gameTypesDetails := make([]*pb.GameTypeDetails, 0)

	for _, gameType := range exres {
		gameTypesDetails = append(gameTypesDetails, &pb.GameTypeDetails{
			Id:          gameType.Id,
			Name:        gameType.Name,
			Description: gameType.Description,
			TypeName:    gameType.TypeName,
			DayName:     gameType.DayName,
			PrizeReward: gameType.PrizeReward,
			TokenBurn:   gameType.TokenBurn,
		})
	}
	return &pb.GameTypes{GameTypes: gameTypesDetails}, nil
}

func (c *GameHandler) GetAllUserPreviousGames(ctx context.Context, request *pb.GetAllUserPreviousGamesRequest) (*pb.Games, error) {
	data := models.Pagination{
		Offset: request.Pagination.Offset,
		Limit:  request.Pagination.Limit,
		Total:  request.Pagination.GetTotal,
	}

	res, err := c.gameService.GetAllUserPreviousGames(request.UserId, &data)
	if err != nil {
		return nil, err.ErrorToGRPCStatus()
	}

	return toGamesProto(res), nil
}

func (c *GameHandler) GetAllUserPreviousGamesByGameType(ctx context.Context, request *pb.GetAllUserPreviousGamesByGameTypeRequest) (*pb.Games, error) {
	data := models.Pagination{
		Offset: request.Pagination.Offset,
		Limit:  request.Pagination.Limit,
		Total:  request.Pagination.GetTotal,
	}

	res, err := c.gameService.GetAllUserPreviousGamesByGameType(request.UserId, strings.ToLower(request.GameType), &data)
	if err != nil {
		return nil, err.ErrorToGRPCStatus()
	}

	return toGamesProto(res), nil
}

func (c *GameHandler) GetAllUserChoiceDivisionsByGameId(ctx context.Context, request *pb.GetAllUserChoiceDivisionsByGameIdRequest) (*pb.DivisionResults, error) {
	res, err := c.gameService.GetAllUserChoiceDivisionsByGameId(request.UserId, request.GameId)
	if err != nil {
		return nil, err.ErrorToGRPCStatus()
	}

	return toDivisionResultsProto(res), nil
}

func (c *GameHandler) GetAllUsersChoiceDivisionsByGameId(ctx context.Context, request *pb.GameId) (*pb.DivisionResults, error) {
	res, err := c.gameService.GetAllUsersChoiceDivisionsByGameId(request.GameId)
	if err != nil {
		return nil, err.ErrorToGRPCStatus()
	}

	return toDivisionResultsProto(res.Divisions), nil
}

func (c *GameHandler) UpdateGamePrizeByGameId(ctx context.Context, request *pb.UpdateGamePrizeByGameIdRequest) (*pb.Empty, error) {
	err := c.gameService.UpdateGamePrizeByGameId(request.GameId, request.Prize, request.AutoCompute)
	if err != nil {
		return nil, err.ErrorToGRPCStatus()
	}
	return &pb.Empty{}, nil
}

func (c *GameHandler) GetUserGamesByTimesAndGameTypes(ctx context.Context, request *pb.GetUserGamesByTimesAndGameTypesRequest) (*pb.GamesAndUserChoices, error) {
	startTime, err := parseTime(request.StartTime, "failed to convert the start time, wrong format #4501")
	if err != nil {
		return nil, err
	}

	endTime, err := parseTime(request.EndTime, "failed to convert the end time, wrong format #4502")
	if err != nil {
		return nil, err
	}

	res, rerr := c.gameService.GetUserGamesByTimesAndGameType(request.UserId, startTime, endTime, request.GameType)
	if rerr != nil {
		return nil, rerr.ErrorToGRPCStatus()
	}

	gamesAndUserChoices := make([]*pb.GameAndUserChoice, 0)

	for _, gameAndUserChoice := range res {

		divisionDetails := make([]*pb.DivisionDetail, 0)
		for _, dv := range gameAndUserChoice.DivisionDetails {
			divisionDetails = append(divisionDetails, &pb.DivisionDetail{
				Division:      dv.Division,
				UserCount:     dv.UserCount,
				DivisionPrize: dv.DivisionPrize,
			})
		}
		userChoicesResult := make([]*pb.UserChoiceResultFiltered, 0)

		for _, ucr := range gameAndUserChoice.UserChoice {

			outMainNumbers := make([]*pb.ChosenMainNumbers, 0)
			for _, c := range ucr.ChosenNumbers {
				outMainNumbers = append(outMainNumbers, &pb.ChosenMainNumbers{
					ChosenMainNumbers: c,
				})
			}

			userChoicesResult = append(userChoicesResult, &pb.UserChoiceResultFiltered{
				UserId:            ucr.UserId,
				ChosenMainNumbers: outMainNumbers,
				ChosenBonusNumber: ucr.ChosenBonusNumber,
				BoughtPrice:       float32(ucr.BoughtPrice),
			})
		}

		gamesAndUserChoices = append(gamesAndUserChoices, &pb.GameAndUserChoice{
			Game:            toGameProto(&gameAndUserChoice.Game),
			DivisionResults: toDivisionResultsProto(gameAndUserChoice.DivisionResult),
			UserChoices:     userChoicesResult,
			TicketUsed:      gameAndUserChoice.TicketUsed,
			DivisionDetails: divisionDetails,
		})
	}
	return &pb.GamesAndUserChoices{Games: gamesAndUserChoices}, nil
}

func (c *GameHandler) UpdateUserGameDivisionPrize(ctx context.Context, request *pb.UpdateUserGameDivisionPrizeRequest) (*pb.Empty, error) {
	divisionUpdates := make([]models.DivisionUpdate, 0)
	for _, dv := range request.DivisionUpdates {
		users := make([]models.UserPrizeUpdate, 0)
		for _, us := range dv.Users {
			users = append(users, models.UserPrizeUpdate{
				UserId:   us.UserId,
				WonPrize: us.WonPrize,
			})
		}
		divisionUpdates = append(divisionUpdates, models.DivisionUpdate{
			DivisionName: dv.DivisionName,
			Users:        users,
		})
	}
	err := c.gameService.UpdateWonPrizeForUsers(request.GameId, divisionUpdates)
	if err != nil {
		return nil, err.ErrorToGRPCStatus()
	}

	return &pb.Empty{}, nil
}
