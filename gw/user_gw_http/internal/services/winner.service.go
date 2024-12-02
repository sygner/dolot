package services

import (
	"context"
	"dolott_user_gw_http/internal/models"
	"dolott_user_gw_http/internal/types"
	pb "dolott_user_gw_http/proto/api/game"
)

type (
	WinnerService interface {
		GetWinnersByGameId(string) ([]models.DivisionResult, *types.Error)
	}
	winnerService struct {
		winnerClient pb.WinnerServiceClient
	}
)

func NewWinnerService(winnerClient pb.WinnerServiceClient) WinnerService {
	return &winnerService{
		winnerClient: winnerClient,
	}
}

func (c *winnerService) GetWinnersByGameId(gameId string) ([]models.DivisionResult, *types.Error) {
	res, err := c.winnerClient.GetWinnersByGameId(context.Background(), &pb.GameId{
		GameId: gameId,
	})
	if err != nil {
		return nil, types.ExtractGRPCErrDetails(err)
	}
	return toDivisionResultsProto(res.DivisionResults), nil
}
