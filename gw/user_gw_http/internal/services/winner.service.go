package services

import (
	"context"
	"dolott_user_gw_http/internal/models"
	"dolott_user_gw_http/internal/types"
	"dolott_user_gw_http/internal/utils"
	pb "dolott_user_gw_http/proto/api/game"
	"fmt"
	"math"
)

type (
	WinnerService interface {
		GetWinnersByGameId(string) (*models.Winners, *types.Error)
		GetWinnersByGameIdCount(string) (*models.WinnersCount, *types.Error)
		UpdateTotalPaidUsers(string, uint64) *types.Error
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

func (c *winnerService) GetWinnersByGameId(gameId string) (*models.Winners, *types.Error) {
	res, err := c.winnerClient.GetWinnersByGameId(context.Background(), &pb.GameId{
		GameId: gameId,
	})
	if err != nil {
		return nil, types.ExtractGRPCErrDetails(err)
	}
	return toWinnerProto(res)
}

func (c *winnerService) GetWinnersByGameIdCount(gameId string) (*models.WinnersCount, *types.Error) {
	res, err := c.winnerClient.GetWinnersByGameIdCount(context.Background(), &pb.GameId{
		GameId: gameId,
	})
	if err != nil {
		return nil, types.ExtractGRPCErrDetails(err)
	}
	var divisionMap map[string]float64
	switch res.GameType {
	case 0:
		divisionMap = LottoWinnerDivisions
	case 1:
		divisionMap = OsLottoWinnerDivisions
	case 2:
		divisionMap = PowerballWinnerDivisions
	case 3:
		divisionMap = AmericanPowerballWinnerDivisions
	}

	divisionResultCount := make([]models.DivisionResultCount, len(res.Divisions.DivisionResultsCount))
	for i := range res.Divisions.DivisionResultsCount {
		divisionPercentage, exists := divisionMap[res.Divisions.DivisionResultsCount[i].Division]
		if exists {
			divisionTotalPrize := float64(res.Prize) * divisionPercentage
			divisionResultCount[i] = models.DivisionResultCount{
				MatchCount:    res.Divisions.DivisionResultsCount[i].MatchCount,
				HasBonus:      res.Divisions.DivisionResultsCount[i].HasBonus,
				Count:         uint32(res.Divisions.DivisionResultsCount[i].Count),
				Division:      res.Divisions.DivisionResultsCount[i].Division,
				DivisionPrize: math.Round(divisionTotalPrize),
			}
		}
	}

	createdAt, rerr := utils.ParseTime(res.CreatedAt, "failed to convert the created at, wrong format #1-8")
	if rerr != nil {
		return nil, rerr
	}
	return &models.WinnersCount{
		Id:           res.Id,
		GameId:       res.GameId,
		GameType:     res.GameType,
		ResultNumber: res.ResultNumber,
		Divisions:    divisionResultCount,
		Prize:        res.Prize,
		JackPot:      res.Jackpot,
		TotalPaid:    res.TotalPaid,
		CreatedAt:    createdAt,
	}, nil
}
func (c *winnerService) UpdateTotalPaidUsers(gameId string, totalPaid uint64) *types.Error {
	_, err := c.winnerClient.UpdateTotalPaid(context.Background(), &pb.UpdateTotalPaidRequest{
		GameId:    gameId,
		TotalPaid: fmt.Sprintf("%d", totalPaid),
	})
	if err != nil {
		return types.ExtractGRPCErrDetails(err)
	}
	return nil
}
