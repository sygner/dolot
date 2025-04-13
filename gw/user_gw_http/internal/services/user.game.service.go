package services

import (
	"context"
	"dolott_user_gw_http/internal/constants"
	"dolott_user_gw_http/internal/models"
	"dolott_user_gw_http/internal/types"
	pb "dolott_user_gw_http/proto/api/game"
	ticket_pb "dolott_user_gw_http/proto/api/ticket"
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
		userClient   pb.UserServiceClient
		ticketClient ticket_pb.TicketServiceClient
	}
)

func NewUserGameService(userClient pb.UserServiceClient, ticketClient ticket_pb.TicketServiceClient) UserGameService {
	return &userGameService{
		userClient:   userClient,
		ticketClient: ticketClient,
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
	_, err := c.ticketClient.UseTickets(context.Background(), &ticket_pb.UseTicketsRequest{
		UserId:            data.UserId,
		GameId:            data.GameId,
		TotalUsingTickets: int32(len(chosenMainNumbers)),
	})
	if err != nil {
		return nil, types.ExtractGRPCErrDetails(err)
	}
	luncPrice, rerr := constants.GetLUNCPriceCoinPaprika()
	if rerr != nil {
		return nil, rerr
	}
	res, err := c.userClient.AddUserChoice(context.Background(), &pb.AddUserChoiceRequest{
		UserId:             data.UserId,
		GameId:             data.GameId,
		ChosenMainNumbers:  chosenMainNumbers,
		ChosenBonusNumbers: chosenBonusNumbers,
		BoughtPrice:        float32(luncPrice) * float32(len(chosenMainNumbers)),
		ShouldReturn:       data.ShouldReturn,
	})
	if err != nil {
		tickets := make([]*ticket_pb.AddTicketRequest, 0)
		for i := 0; i < int(len(chosenMainNumbers)); i++ {
			tickets = append(tickets, &ticket_pb.AddTicketRequest{UserId: data.UserId, TicketType: "purchased"})
		}
		_, rerr := c.ticketClient.AddTickets(context.Background(), &ticket_pb.AddTicketsRequest{
			Tickets: tickets,
		})
		fmt.Println(rerr)
		return nil, types.ExtractGRPCErrDetails(err)
	}
	// check id is not nil
	if res != nil && res.Id != "" {
		return toUserChoiceProto(res), nil
	} else {
		return nil, nil
	}
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
