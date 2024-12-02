package handlers

import (
	"context"
	"dolott_authentication/internal/models"
	"dolott_authentication/internal/services"
	"dolott_authentication/internal/types"
	pb "dolott_authentication/proto/api"
	"strings"

	"google.golang.org/protobuf/proto"
)

type AuthenticationHandler struct {
	pb.UnimplementedAuthentcationServiceServer
	authenticationService services.AuthenticationService
}

func NewAuthenticationHandler(authenticationService services.AuthenticationService) *AuthenticationHandler {
	return &AuthenticationHandler{
		authenticationService: authenticationService,
	}
}

func (c *AuthenticationHandler) Signup(ctx context.Context, request *pb.SignupRequest) (*pb.Token, error) {
	var provider *string
	if request.Provider != nil {
		p := strings.ToLower(*request.Provider)
		provider = &p
	} else {
		p := "local"
		provider = &p
	}
	data := &models.UserDTO{
		Email:      request.Email,
		UserRole:   models.USER.String(),
		UserStatus: models.ONGOING.String(),
		Password:   request.Password,
		Agent:      request.Agent,
		Provider:   provider,
		IsSSO:      request.IsSso,
		Ip:         request.Ip,
	}
	if !data.IsSSO {
		if data.Password == nil {
			return nil, types.NewBadRequestError("enter password, error code #8002").ErrorToGRPCStatus()
		}
	} else {
		if data.Provider == nil {
			return nil, types.NewBadRequestError("enter provider, error code #8003").ErrorToGRPCStatus()
		}
	}
	res, err := c.authenticationService.Signup(data)
	if err != nil {
		return nil, err.ErrorToGRPCStatus()
	}
	if res != nil {
		return &pb.Token{
			AccessToken:          res.AccessToken,
			RefreshToken:         res.RefreshToken,
			UserId:               res.UserId,
			UserRole:             res.UserRole,
			SessionId:            res.SessionId,
			TokenStatus:          res.TokenStatus,
			Ip:                   res.Ip,
			Agent:                res.Agent,
			CreatedAt:            res.CreatedAt.Unix(),
			AccessTokenExpireAt:  res.AccessTokenExpireAt.Unix(),
			RefreshTokenExpireAt: res.RefreshTokenExpireAt.Unix(),
		}, nil
	}

	return nil, nil

}

func (c *AuthenticationHandler) Verify(ctx context.Context, request *pb.VerifyRequest) (*pb.VerifyResponse, error) {
	res, err := c.authenticationService.Verify(&models.VerifyDTO{
		Code:         request.Code,
		Agent:        request.Agent,
		Ip:           request.Ip,
		VerifyMethod: int32(request.VerifyMethod),
		NewPassword:  request.NewPassword,
	})
	if err != nil {
		return nil, err.ErrorToGRPCStatus()
	}

	return &pb.VerifyResponse{
		Token: tokenToProto(&res.Token),
		Value: res.Value,
	}, nil

}

func (c *AuthenticationHandler) Signin(ctx context.Context, request *pb.SigninRequest) (*pb.SigninResponse, error) {
	var provider *string
	if request.Provider != nil {
		p := strings.ToLower(*request.Provider)
		provider = &p
	}
	loginDTO := models.LoginDTO{
		Value:       request.Value,
		Agent:       request.Agent,
		Ip:          request.Ip,
		IsSSO:       request.IsSso,
		Provider:    provider,
		LoginMethod: int32(request.SigninMethod),
	}

	if request.SigninMethod == 0 {
		if request.Password == nil {
			return nil, types.NewBadRequestError("password is empty, error code #8001").ErrorToGRPCStatus()
		}
		loginDTO.Password = *request.Password
	}

	res, err := c.authenticationService.Signin(&loginDTO)
	if err != nil {
		return nil, err.ErrorToGRPCStatus()
	}

	if request.SigninMethod == 0 {
		return &pb.SigninResponse{
			Token: &pb.Token{
				AccessToken:          res.AccessToken,
				RefreshToken:         res.RefreshToken,
				UserId:               res.UserId,
				UserRole:             res.UserRole,
				SessionId:            res.SessionId,
				TokenStatus:          res.TokenStatus,
				Ip:                   res.Ip,
				Agent:                res.Agent,
				CreatedAt:            res.CreatedAt.Unix(),
				AccessTokenExpireAt:  res.AccessTokenExpireAt.Unix(),
				RefreshTokenExpireAt: res.RefreshTokenExpireAt.Unix(),
			},
		}, nil
	} else {
		if res != nil {
			return &pb.SigninResponse{
				Token: &pb.Token{
					AccessToken:          res.AccessToken,
					RefreshToken:         res.RefreshToken,
					UserId:               res.UserId,
					UserRole:             res.UserRole,
					SessionId:            res.SessionId,
					TokenStatus:          res.TokenStatus,
					Ip:                   res.Ip,
					Agent:                res.Agent,
					CreatedAt:            res.CreatedAt.Unix(),
					AccessTokenExpireAt:  res.AccessTokenExpireAt.Unix(),
					RefreshTokenExpireAt: res.RefreshTokenExpireAt.Unix(),
				},
			}, nil
		}
		return &pb.SigninResponse{
			Msg: proto.String("email sent"),
		}, nil
	}

}
