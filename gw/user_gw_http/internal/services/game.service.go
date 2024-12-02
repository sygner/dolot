package services

import (
	"context"
	"dolott_user_gw_http/internal/models"
	"dolott_user_gw_http/internal/types"
	pb "dolott_user_gw_http/proto/api/game"
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
	}
	gameService struct {
		gameClient pb.GameServiceClient
	}
)

func NewGameService(gameClient pb.GameServiceClient) GameService {
	return &gameService{
		gameClient: gameClient,
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
		Name:      data.Name,
		GameType:  pb.GameType(data.GameTypeInt),
		StartTime: data.StartTime,
		EndTime:   data.EndTime,
		CreatorId: data.CreatorId,
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

	return toDivisionResultsProto(res.DivisionResults), nil
}
