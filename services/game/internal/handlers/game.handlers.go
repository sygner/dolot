package handlers

import (
	"context"
	"dolott_game/internal/models"
	"dolott_game/internal/services"
	"dolott_game/internal/types"
	pb "dolott_game/proto/api"
	"strconv"
	"time"
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

func toGamesProto(res *models.Games) *pb.Games {
	games := make([]*pb.Game, 0)
	for _, game := range res.Games {
		games = append(games, toGameProto(&game))
	}
	return &pb.Games{Games: games, Total: res.Total}
}

func toGameProto(res *models.Game) *pb.Game {
	return &pb.Game{
		NumMainNumbers:   res.NumMainNumbers,
		NumBonusNumbers:  res.NumBonusNumbers,
		MainNumberRange:  res.MainNumberRange,
		BonusNumberRange: res.BonusNumberRange,
		CreatorId:        res.CreatorId,
		Id:               res.Id,
		Name:             res.Name,
		Result:           res.Result,
		GameType:         pb.GameType(models.FromString(res.GameType)),
		StartTime:        strconv.FormatInt(res.StartTime.Unix(), 10),
		EndTime:          strconv.FormatInt(res.EndTime.Unix(), 10),
		CreatedAt:        strconv.FormatInt(res.CreatedAt.Unix(), 10),
	}
}

func toUserChoiceResultDetailsProto(res []models.UserChoiceResultDetail) []*pb.UserChoiceResultDetail {
	choices := make([]*pb.UserChoiceResultDetail, 0)
	for _, division := range res {
		choices = append(choices, toUserChoiceResultDetailProto(&division))
	}
	return choices
}
func toUserChoiceResultDetailProto(res *models.UserChoiceResultDetail) *pb.UserChoiceResultDetail {
	return &pb.UserChoiceResultDetail{
		UserId:            res.UserId,
		ChosenMainNumber:  res.ChosenMainNumbers,
		ChosenBonusNumber: res.ChosenBonusNumber,
		MatchCount:        res.MatchCount,
	}
}

func toDivisionResultsProto(res []models.DivisionResult) *pb.DivisionResults {
	divisions := make([]*pb.DivisionResult, 0)
	for _, division := range res {
		divisions = append(divisions, toDivisionResultProto(&division))
	}
	return &pb.DivisionResults{DivisionResults: divisions}
}
func toDivisionResultProto(res *models.DivisionResult) *pb.DivisionResult {
	return &pb.DivisionResult{
		HasBonus:   res.HasBonus,
		UserChoice: toUserChoiceResultDetailsProto(res.UserChoices),
		MatchCount: res.MatchCount,
	}
}

func toUserChoisesProto(res models.UserChoices) *pb.UserChoices {
	users := make([]*pb.UserChoice, 0)
	for _, user := range res.UserChoices {
		users = append(users, toUserChoiceProto(&user))
	}
	return &pb.UserChoices{UserChoices: users, Total: res.Total}
}

func toUserChoiceProto(res *models.UserChoice) *pb.UserChoice {
	outMainNumbers := make([]*pb.ChosenMainNumbers, 0)
	for _, c := range res.ChosenMainNumbers {
		outMainNumbers = append(outMainNumbers, &pb.ChosenMainNumbers{
			ChosenMainNumbers: c,
		})
	}
	outBonusNumbers := make([]*pb.ChosenBonusNumbers, 0)
	for _, c := range res.ChosenBonusNumbers {
		outBonusNumbers = append(outBonusNumbers, &pb.ChosenBonusNumbers{
			ChosenBonusNumbers: c,
		})
	}
	return &pb.UserChoice{
		Id:                 res.Id,
		UserId:             res.UserId,
		GameId:             res.GameId,
		ChosenMainNumbers:  outMainNumbers,
		ChosenBonusNumbers: outBonusNumbers,
		CreatedAt:          res.CreatedAt.Format("2006-01-02 15:04:05"),
	}
}
func parseTime(timeStr string, errMsg string) (time.Time, error) {
	timeInt64, err := strconv.ParseInt(timeStr, 10, 64)
	if err != nil {
		rerr := types.NewInternalError(errMsg)
		return time.Time{}, rerr.ErrorToGRPCStatus()
	}
	return time.Unix(timeInt64, 0), nil
}
