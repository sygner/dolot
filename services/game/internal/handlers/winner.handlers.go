package handlers

import (
	"context"
	"dolott_game/internal/services"
	pb "dolott_game/proto/api"
	"strconv"
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

func (c *WinnerHandler) GetWinnersByGameId(ctx context.Context, request *pb.GameId) (*pb.Winner, error) {
	res, err := c.winnerService.GetWinnersByGameId(request.GameId)
	if err != nil {
		return nil, err.ErrorToGRPCStatus()
	}

	return toWinnerResultProto(res), nil
}

func (c *WinnerHandler) GetWinnersByGameIdCount(ctx context.Context, request *pb.GameId) (*pb.WinnerCount, error) {
	res, err := c.winnerService.GetWinnersByGameId(request.GameId)
	if err != nil {
		return nil, err.ErrorToGRPCStatus()
	}

	divisions := make([]*pb.DivisionResultCount, 0)
	for _, division := range res.Divisions {
		divisions = append(divisions, &pb.DivisionResultCount{
			Division:   division.Division,
			MatchCount: division.MatchCount,
			HasBonus:   division.HasBonus,
			Count:      int32(len(division.UserChoices)),
		})
	}
	return &pb.WinnerCount{
		Id:           res.Id,
		GameId:       res.GameId,
		GameType:     res.GameType,
		ResultNumber: res.ResultNumber,
		TotalPaid:    res.TotalPaid,
		Prize:        uint32(res.Prize),
		Jackpot:      res.JackPot,
		Divisions:    &pb.DivisionResultsCount{DivisionResultsCount: divisions},
		CreatedAt:    strconv.FormatInt(res.CreatedAt.Unix(), 10),
	}, nil
}

func (c *WinnerHandler) UpdateTotalPaid(ctx context.Context, request *pb.UpdateTotalPaidRequest) (*pb.Empty, error) {
	err := c.winnerService.UpdateTotalPaidUsers(request.GameId, request.TotalPaid)
	if err != nil {
		return nil, err.ErrorToGRPCStatus()
	}
	return &pb.Empty{}, nil
}

func (c *WinnerHandler) GetLastWinnersByGameType(ctx context.Context, request *pb.GameTypeRequest) (*pb.Winner, error) {
	res, err := c.winnerService.GetLastWinnersByGameType(int32(request.GameType))
	if err != nil {
		return nil, err.ErrorToGRPCStatus()
	}
	return toWinnerResultProto(res), nil
}
