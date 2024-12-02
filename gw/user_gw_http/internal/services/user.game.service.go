package services

import (
	"context"
	"dolott_user_gw_http/internal/models"
	"dolott_user_gw_http/internal/types"
	pb "dolott_user_gw_http/proto/api/game"
	"fmt"
)

type (
	UserGameService interface {
		AddUserChoice(*models.AddUserChoiceDTO) (*models.UserChoice, *types.Error)
		GetUserChoicesByUserId(int32, *models.Pagination) (*models.UserChoices, *types.Error)
		GetUserChoicesByUserIdAndTimeRange(*models.GetUserChoicesByUserIdAndTimeRangeDTO) (*models.UserChoices, *types.Error)
		GetUserChoicesByGameIdAndPagination(string, *models.Pagination) (*models.UserChoices, *types.Error)
		GetAllUserGames(int32) ([]string, *types.Error)
	}
	userGameService struct {
		userClient pb.UserServiceClient
	}
)

func NewUserGameService(userClient pb.UserServiceClient) UserGameService {
	return &userGameService{
		userClient: userClient,
	}
}

func (c *userGameService) AddUserChoice(data *models.AddUserChoiceDTO) (*models.UserChoice, *types.Error) {
	chosenMainNumbers := make([]*pb.ChosenMainNumbers, 0)
	for i := range data.ChosenMainNumbers {
		chosenMainNumbers = append(chosenMainNumbers, &pb.ChosenMainNumbers{
			ChosenMainNumbers: data.ChosenMainNumbers[i],
		})
	}

	chosenBonusNumbers := make([]*pb.ChosenBonusNumbers, 0)
	for i := range data.ChosenBonusNumbers {
		chosenBonusNumbers = append(chosenBonusNumbers, &pb.ChosenBonusNumbers{
			ChosenBonusNumbers: data.ChosenBonusNumbers[i],
		})
	}
	res, err := c.userClient.AddUserChoice(context.Background(), &pb.AddUserChoiceRequest{
		UserId:             data.UserId,
		GameId:             data.GameId,
		ChosenMainNumbers:  chosenMainNumbers,
		ChosenBonusNumbers: chosenBonusNumbers,
	})
	if err != nil {
		return nil, types.ExtractGRPCErrDetails(err)
	}
	return toUserChoiceProto(res), nil
}

func (c *userGameService) GetUserChoicesByUserId(data int32, pagination *models.Pagination) (*models.UserChoices, *types.Error) {
	res, err := c.userClient.GetUserChoicesByUserId(context.Background(), &pb.GetUserChoicesByUserIdRequest{
		UserId: data,
		Pagination: &pb.Pagination{
			Offset:   pagination.Offset,
			Limit:    pagination.Limit,
			GetTotal: pagination.Total,
		},
	})
	if err != nil {
		return nil, types.ExtractGRPCErrDetails(err)
	}

	return toUserChoisesProto(res), nil
}

func (c *userGameService) GetUserChoicesByUserIdAndTimeRange(data *models.GetUserChoicesByUserIdAndTimeRangeDTO) (*models.UserChoices, *types.Error) {
	fmt.Println(data)
	res, err := c.userClient.GetUserChoicesByUserIdAndTimeRange(context.Background(), &pb.GetUserChoicesByUserIdAndTimeRangeRequest{
		UserId:    data.UserId,
		StartTime: data.StartTime,
		EndTime:   data.EndTime,
	})
	if err != nil {
		return nil, types.ExtractGRPCErrDetails(err)
	}

	return toUserChoisesProto(res), nil
}

func (c *userGameService) GetUserChoicesByGameIdAndPagination(gameId string, pagination *models.Pagination) (*models.UserChoices, *types.Error) {
	res, err := c.userClient.GetUserChoicesByGameIdAndPagination(context.Background(), &pb.GetUserChoicesByGameIdRequest{
		GameId: gameId,
		Pagination: &pb.Pagination{
			Offset:   pagination.Offset,
			Limit:    pagination.Limit,
			GetTotal: pagination.Total,
		},
	})
	if err != nil {
		return nil, types.ExtractGRPCErrDetails(err)
	}
	return toUserChoisesProto(res), nil
}

func (c *userGameService) GetAllUserGames(userId int32) ([]string, *types.Error) {
	res, err := c.userClient.GetAllUserGames(context.Background(), &pb.UserId{UserId: userId})
	if err != nil {
		return nil, types.ExtractGRPCErrDetails(err)
	}

	return res.Ids, nil
}
