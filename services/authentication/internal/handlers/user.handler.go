package handlers

import (
	"context"
	"dolott_authentication/internal/models"
	"dolott_authentication/internal/services"
	pb "dolott_authentication/proto/api"
)

type UserHandler struct {
	pb.UnimplementedUserServiceServer
	userService         services.UserService
	loginHistoryService services.LoginHistoryService
}

func NewUserHandler(userService services.UserService, loginHistoryService services.LoginHistoryService) *UserHandler {
	return &UserHandler{
		userService:         userService,
		loginHistoryService: loginHistoryService,
	}
}

func (c *UserHandler) GetUsers(ctx context.Context, request *pb.Pagination) (*pb.Users, error) {
	data := models.Pagination{
		Offset: request.Offset,
		Limit:  request.Limit,
		Total:  request.GetTotal,
	}

	res, err := c.userService.GetUsers(&data)
	if err != nil {
		return nil, err.ErrorToGRPCStatus()
	}

	return userToProtoList(res), nil
}

func (c *UserHandler) GetUserByUserId(ctx context.Context, request *pb.UserId) (*pb.User, error) {
	res, err := c.userService.GetUserByUserId(request.UserId)
	if err != nil {
		return nil, err.ErrorToGRPCStatus()
	}
	return userToProto(res), nil
}

func (c *UserHandler) GetUserByEmail(ctx context.Context, request *pb.Email) (*pb.User, error) {
	res, err := c.userService.GetUserByEmail(request.Email)
	if err != nil {
		return nil, err.ErrorToGRPCStatus()
	}
	return userToProto(res), nil
}
func (c *UserHandler) GetUserByAccountUsername(ctx context.Context, request *pb.AccountUsername) (*pb.User, error) {
	res, err := c.userService.GetUserByAccountUsername(request.AccountUsername)
	if err != nil {
		return nil, err.ErrorToGRPCStatus()
	}
	return userToProto(res), nil
}

func (c *UserHandler) ChangeUserStatus(ctx context.Context, request *pb.ChangeUserStatusRequest) (*pb.Empty, error) {
	err := c.userService.ChangeUserStatus(request.UserId, request.Status)
	if err != nil {
		return nil, err.ErrorToGRPCStatus()
	}

	return &pb.Empty{}, nil
}

func (c *UserHandler) GetLoginHistoryByUserId(ctx context.Context, request *pb.GetLoginHistoryRequest) (*pb.LoginHistory, error) {
	res, err := c.loginHistoryService.GetLoginHistoryByUserId(&models.Pagination{
		Offset: request.Pagination.Offset,
		Limit:  request.Pagination.Limit,
		Total:  request.Pagination.GetTotal,
	}, request.UserId)
	if err != nil {
		return nil, err.ErrorToGRPCStatus()
	}
	logins := make([]int64, 0)
	for i := range res.LoginHistories {
		logins = append(logins, res.LoginHistories[i].Unix())
	}
	return &pb.LoginHistory{LoggedInAt: logins, Total: res.Total}, nil
}

func (c *UserHandler) ResetPassword(ctx context.Context, request *pb.ResetPasswordRequest) (*pb.Empty, error) {
	err := c.userService.ResetPassword(&models.ResetPasswordDTO{
		UserId:          request.UserId,
		CurrentPassword: request.CurrentPassword,
		NewPassword:     request.NewPassword,
	})

	if err != nil {
		return nil, err.ErrorToGRPCStatus()
	}
	return &pb.Empty{}, nil
}

func (c *UserHandler) ForgotPassword(ctx context.Context, request *pb.Email) (*pb.Empty, error) {
	err := c.userService.ForgotPassword(request.Email)

	if err != nil {
		return nil, err.ErrorToGRPCStatus()
	}
	return &pb.Empty{}, nil
}
func userToProtoList(userData *models.Users) *pb.Users {
	users := make([]*pb.User, 0, len(userData.Users))
	for _, user := range userData.Users {
		users = append(users, userToProto(&user))
	}

	return &pb.Users{Users: users, Total: userData.Total}
}

func userToProto(up *models.User) *pb.User {
	return &pb.User{
		UserId:          up.UserId,
		PhoneNumber:     up.PhoneNumber,
		Email:           up.Email,
		UserRole:        up.UserRole,
		UserStatus:      up.UserStatus,
		AccountUsername: up.AccountUsername,
		Provider:        up.Provider,
		IsSso:           up.IsSSO,
		CreatedAt:       up.CreatedAt.Unix(),
	}
}
