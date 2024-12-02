package services

import (
	"context"
	"dolott_user_gw_http/internal/models"
	"dolott_user_gw_http/internal/types"
	"dolott_user_gw_http/internal/utils"
	pb "dolott_user_gw_http/proto/api/authentication"
)

type (
	UserService interface {
		GetUserByUserId(int32) (*models.User, *types.Error)
		GetUserByEmail(string) (*models.User, *types.Error)
		GetLoginHistoryByUserId(*models.Pagination, int32) (*models.LoginHistories, *types.Error)
		ResetPassword(*models.ResetPasswordDTO) *types.Error
		GetUserByAccountUsername(string) (*models.User, *types.Error)
	}
	userService struct {
		userClient pb.UserServiceClient
	}
)

func NewUserService(userClient pb.UserServiceClient) UserService {
	return &userService{
		userClient: userClient,
	}
}

func (c *userService) GetUserByUserId(userId int32) (*models.User, *types.Error) {
	res, err := c.userClient.GetUserByUserId(context.Background(), &pb.UserId{
		UserId: userId,
	})

	if err != nil {
		return nil, types.ExtractGRPCErrDetails(err)
	}
	createdAt, rerr := utils.GetTimebyTimestamp(res.CreatedAt)
	if rerr != nil {
		return nil, rerr
	}
	return &models.User{
		UserId:          res.UserId,
		PhoneNumber:     res.PhoneNumber,
		Email:           res.Email,
		AccountUsername: res.AccountUsername,
		UserRole:        res.UserRole,
		UserStatus:      res.UserStatus,
		CreatedAt:       *createdAt,
	}, nil
}

func (c *userService) GetUserByEmail(email string) (*models.User, *types.Error) {
	res, err := c.userClient.GetUserByEmail(context.Background(), &pb.Email{
		Email: email,
	})

	if err != nil {
		return nil, types.ExtractGRPCErrDetails(err)
	}
	createdAt, rerr := utils.GetTimebyTimestamp(res.CreatedAt)
	if rerr != nil {
		return nil, rerr
	}
	return &models.User{
		UserId:          res.UserId,
		PhoneNumber:     res.PhoneNumber,
		AccountUsername: res.AccountUsername,
		Email:           res.Email,
		UserRole:        res.UserRole,
		UserStatus:      res.UserStatus,
		CreatedAt:       *createdAt,
	}, nil
}

func (c *userService) GetUserByAccountUsername(accountUsername string) (*models.User, *types.Error) {
	res, err := c.userClient.GetUserByAccountUsername(context.Background(), &pb.AccountUsername{
		AccountUsername: accountUsername,
	})

	if err != nil {
		return nil, types.ExtractGRPCErrDetails(err)
	}
	createdAt, rerr := utils.GetTimebyTimestamp(res.CreatedAt)
	if rerr != nil {
		return nil, rerr
	}
	return &models.User{
		UserId:          res.UserId,
		PhoneNumber:     res.PhoneNumber,
		AccountUsername: res.AccountUsername,
		Email:           res.Email,
		UserRole:        res.UserRole,
		UserStatus:      res.UserStatus,
		CreatedAt:       *createdAt,
	}, nil
}

func (c *userService) GetLoginHistoryByUserId(data *models.Pagination, userId int32) (*models.LoginHistories, *types.Error) {
	res, err := c.userClient.GetLoginHistoryByUserId(context.Background(), &pb.GetLoginHistoryRequest{
		Pagination: &pb.Pagination{
			Offset:   data.Offset,
			Limit:    data.Limit,
			GetTotal: data.Total,
		},
		UserId: userId,
	})
	if err != nil {
		return nil, types.ExtractGRPCErrDetails(err)
	}

	return &models.LoginHistories{LoginHistories: res.LoggedInAt, Total: res.Total}, nil
}

func (c *userService) ResetPassword(data *models.ResetPasswordDTO) *types.Error {
	_, err := c.userClient.ResetPassword(context.Background(), &pb.ResetPasswordRequest{
		UserId:          data.UserId,
		CurrentPassword: data.CurrentPassword,
		NewPassword:     data.NewPassword,
	})
	if err != nil {
		return types.ExtractGRPCErrDetails(err)
	}
	return nil
}
