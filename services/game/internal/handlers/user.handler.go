package handlers

import (
	"context"
	"dolott_game/internal/models"
	"dolott_game/internal/services"
	"dolott_game/internal/types"
	"fmt"
	"strconv"
	"time"

	pb "dolott_game/proto/api"
)

type UserHandler struct {
	userServices services.UserServices
	pb.UnimplementedUserServiceServer
}

func NewUserHandler(userServices services.UserServices) *UserHandler {
	return &UserHandler{
		userServices: userServices,
	}
}

func (c *UserHandler) AddUserChoice(ctx context.Context, request *pb.AddUserChoiceRequest) (*pb.UserChoice, error) {
	outMainNumbers := make([][]int32, 0)
	for _, c := range request.ChosenMainNumbers {
		outMainNumbers = append(outMainNumbers, c.ChosenMainNumbers)
	}

	outBonusNumbers := make([][]int32, 0)
	for _, c := range request.ChosenBonusNumbers {
		outBonusNumbers = append(outBonusNumbers, c.ChosenBonusNumbers)
	}

	data := &models.AddUserChoiceDTO{
		UserId:             request.UserId,
		GameId:             request.GameId,
		ChosenMainNumbers:  outMainNumbers,
		ChosenBonusNumbers: outBonusNumbers,
	}
	res, err := c.userServices.AddUserChoice(data, request.ShouldReturn)
	if err != nil {
		return nil, err.ErrorToGRPCStatus()
	}
	if request.ShouldReturn {
		return toUserChoiceProto(res), nil
	} else {
		return nil, nil
	}
}

func (c *UserHandler) GetUserChoicesByUserId(ctx context.Context, request *pb.GetUserChoicesByUserIdRequest) (*pb.UserChoices, error) {
	data := models.Pagination{
		Offset: request.Pagination.GetOffset(),
		Limit:  request.Pagination.GetLimit(),
		Total:  request.Pagination.GetGetTotal(),
	}
	res, err := c.userServices.GetUserChoicesByUserId(request.UserId, &data)
	if err != nil {
		return nil, err.ErrorToGRPCStatus()
	}

	return toUserChoisesProto(*res), nil
}

// Range function to fetch user choices by userId and time range
func (c *UserHandler) GetUserChoicesByUserIdAndTimeRange(ctx context.Context, request *pb.GetUserChoicesByUserIdAndTimeRangeRequest) (*pb.UserChoices, error) {
	// Helper function to parse time from both RFC3339 and Unix timestamp formats
	parseTime := func(timeStr string) (time.Time, error) {
		// Try parsing as RFC3339 format first
		parsedTime, err := time.Parse(time.RFC3339, timeStr)
		if err == nil {
			return parsedTime, nil
		}

		// If RFC3339 fails, attempt to parse as Unix timestamp
		unixTimestamp, unixErr := strconv.ParseInt(timeStr, 10, 64)
		if unixErr != nil {
			return time.Time{}, fmt.Errorf("invalid time format: %v", err)
		}

		// Convert Unix timestamp to time.Time
		return time.Unix(unixTimestamp, 0), nil
	}

	// Parse start time
	startTime, err := parseTime(request.StartTime)
	if err != nil {
		return nil, types.NewBadRequestError(fmt.Sprintf("invalid start time format: %v", err)).ErrorToGRPCStatus()
	}

	// Parse end time
	endTime, err := parseTime(request.EndTime)
	if err != nil {
		return nil, types.NewBadRequestError(fmt.Sprintf("invalid end time format: %v", err)).ErrorToGRPCStatus()
	}

	// Fetch user choices by userId and time range
	res, rerr := c.userServices.GetUserChoicesByUserIdAndTimeRange(request.UserId, startTime, endTime)
	if rerr != nil {
		return nil, rerr.ErrorToGRPCStatus()
	}

	// Convert the result to the protobuf response format
	return toUserChoisesProto(models.UserChoices{UserChoices: res, Total: nil}), nil
}

func (c *UserHandler) GetAllUserGames(ctx context.Context, request *pb.UserId) (*pb.GameIds, error) {
	res, err := c.userServices.GetAllUserGames(request.UserId)
	if err != nil {
		return nil, err.ErrorToGRPCStatus()
	}
	return &pb.GameIds{Ids: res}, nil
}

func (c *UserHandler) GetUserChoicesByGameIdAndPagination(ctx context.Context, request *pb.GetUserChoicesByGameIdRequest) (*pb.UserChoices, error) {
	data := models.Pagination{
		Offset: request.Pagination.Offset,
		Limit:  request.Pagination.Limit,
		Total:  request.Pagination.GetTotal,
	}
	res, err := c.userServices.GetUserChoicesByGameIdAndPagination(request.GameId, &data)
	if err != nil {
		return nil, err.ErrorToGRPCStatus()
	}

	return toUserChoisesProto(*res), nil
}
