package handlers

import (
	"context"
	"dolott_game/internal/services"
	pb "dolott_game/proto/api"
)

type WinnerHandler struct {
	pb.UnimplementedWinnerServiceServer
	winnerService services.WinnerServices
}

func NewWinenrHandler(winnerService services.WinnerServices) *WinnerHandler {
	return &WinnerHandler{
		winnerService: winnerService,
	}
}

func (c *WinnerHandler) GetWinnersByGameId(ctx context.Context, request *pb.GameId) (*pb.DivisionResults, error) {
	res, err := c.winnerService.GetWinnersByGameId(request.GameId)
	if err != nil {
		return nil, err.ErrorToGRPCStatus()
	}

	return toDivisionResultsProto(res), nil
}
