package services

import (
	"context"
	"dolott_user_gw_http/internal/models"
	"dolott_user_gw_http/internal/types"
	"dolott_user_gw_http/internal/utils"
	pb "dolott_user_gw_http/proto/api/authentication"
	"fmt"
)

type (
	AuthenticationService interface {
		Signup(*models.SignupDTO) (*models.Token, *types.Error)
		Signin(*models.SigninDTO) (*models.Token, *types.Error)
		Verify(*models.VerifyDTO) (*models.VerifyResponse, *types.Error)
		RenewToken(*models.RenewTokenDTO) (*models.Token, *types.Error)
		ForgotPassword(string) *types.Error
	}
	authenticationService struct {
		authenticationClient pb.AuthentcationServiceClient
		userClient           pb.UserServiceClient
		tokenClient          pb.TokenServiceClient
	}
)

func NewAuthenticationService(authenticationClient pb.AuthentcationServiceClient, userClient pb.UserServiceClient, tokenClient pb.TokenServiceClient) AuthenticationService {
	return &authenticationService{
		authenticationClient: authenticationClient,
		userClient:           userClient,
		tokenClient:          tokenClient,
	}
}

func (c *authenticationService) Signup(data *models.SignupDTO) (*models.Token, *types.Error) {
	println(data.IsSSO)
	res, err := c.authenticationClient.Signup(context.Background(), &pb.SignupRequest{
		Email:    data.Email,
		Password: data.Password,
		IsSso:    data.IsSSO,
		Provider: data.Provider,
		Agent:    data.Agent,
		Ip:       data.Ip,
	})

	if err != nil {
		return nil, types.ExtractGRPCErrDetails(err)
	}

	if res != nil && res.AccessToken != "" {
		createdAt, err := utils.GetTimebyTimestamp(res.CreatedAt)
		if err != nil {
			return nil, err
		}
		accessTokenExpireAt, err := utils.GetTimebyTimestamp(res.AccessTokenExpireAt)
		if err != nil {
			return nil, err
		}
		refreshTokenExpireAt, err := utils.GetTimebyTimestamp(res.RefreshTokenExpireAt)
		if err != nil {
			return nil, err
		}
		return &models.Token{
			AccessToken:          res.AccessToken,
			RefreshToken:         res.RefreshToken,
			UserId:               res.UserId,
			UserRole:             res.UserRole,
			SessionId:            res.SessionId,
			TokenStatus:          res.TokenStatus,
			Ip:                   res.Ip,
			Agent:                res.Agent,
			CreatedAt:            *createdAt,
			AccessTokenExpireAt:  *accessTokenExpireAt,
			RefreshTokenExpireAt: *refreshTokenExpireAt,
		}, nil
	}
	return nil, nil
}

func (c *authenticationService) Signin(data *models.SigninDTO) (*models.Token, *types.Error) {
	fmt.Println(data)
	res, err := c.authenticationClient.Signin(context.Background(), &pb.SigninRequest{
		Value:        data.Value,
		Password:     data.Password,
		SigninMethod: pb.SIGNIN_METHOD(data.Signin_Method),
		Provider:     data.Provider,
		IsSso:        data.IsSSO,
		Agent:        data.Agent,
		Ip:           data.Ip,
	})
	if err != nil {
		return nil, types.ExtractGRPCErrDetails(err)
	}
	if res.Token != nil {
		createdAt, err := utils.GetTimebyTimestamp(res.Token.CreatedAt)
		if err != nil {
			return nil, err
		}
		accessTokenExpireAt, err := utils.GetTimebyTimestamp(res.Token.AccessTokenExpireAt)
		if err != nil {
			return nil, err
		}
		refreshTokenExpireAt, err := utils.GetTimebyTimestamp(res.Token.RefreshTokenExpireAt)
		if err != nil {
			return nil, err
		}
		return &models.Token{
			AccessToken:          res.Token.AccessToken,
			RefreshToken:         res.Token.RefreshToken,
			UserId:               res.Token.UserId,
			UserRole:             res.Token.UserRole,
			SessionId:            res.Token.SessionId,
			TokenStatus:          res.Token.TokenStatus,
			Ip:                   res.Token.Ip,
			Agent:                res.Token.Agent,
			CreatedAt:            *createdAt,
			AccessTokenExpireAt:  *accessTokenExpireAt,
			RefreshTokenExpireAt: *refreshTokenExpireAt,
		}, nil
	}
	return nil, nil
}

func (c *authenticationService) Verify(data *models.VerifyDTO) (*models.VerifyResponse, *types.Error) {
	res, err := c.authenticationClient.Verify(context.Background(), &pb.VerifyRequest{
		Code:         data.Code,
		VerifyMethod: pb.VERIFY_METHOD(data.Verify_Method),
		NewPassword:  data.NewPassword,
		// IsSso:        data.IsSSO,
		Agent: data.Agent,
		Ip:    data.Ip,
	})
	if err != nil {
		return nil, types.ExtractGRPCErrDetails(err)
	}
	createdAt, rerr := utils.GetTimebyTimestamp(res.Token.CreatedAt)
	if rerr != nil {
		return nil, rerr
	}
	accessTokenExpireAt, rerr := utils.GetTimebyTimestamp(res.Token.AccessTokenExpireAt)
	if rerr != nil {
		return nil, rerr
	}
	refreshTokenExpireAt, rerr := utils.GetTimebyTimestamp(res.Token.RefreshTokenExpireAt)
	if rerr != nil {
		return nil, rerr
	}
	return &models.VerifyResponse{
		Token: models.Token{
			AccessToken:          res.Token.AccessToken,
			RefreshToken:         res.Token.RefreshToken,
			UserId:               res.Token.UserId,
			UserRole:             res.Token.UserRole,
			SessionId:            res.Token.SessionId,
			TokenStatus:          res.Token.TokenStatus,
			Ip:                   res.Token.Ip,
			Agent:                res.Token.Agent,
			CreatedAt:            *createdAt,
			AccessTokenExpireAt:  *accessTokenExpireAt,
			RefreshTokenExpireAt: *refreshTokenExpireAt,
		},
		Value: res.Value,
	}, nil
}

func (c *authenticationService) ForgotPassword(email string) *types.Error {
	_, err := c.userClient.ForgotPassword(context.Background(), &pb.Email{
		Email: email,
	})
	if err != nil {
		return types.ExtractGRPCErrDetails(err)
	}
	return nil
}

func (c *authenticationService) RenewToken(data *models.RenewTokenDTO) (*models.Token, *types.Error) {
	res, err := c.tokenClient.RenewToken(context.Background(), &pb.RenewTokenRequest{
		AccessToken:  data.AccessToken,
		RefreshToken: data.RefreshToken,
		Agent:        data.Agent,
		Ip:           data.Ip,
	})
	if err != nil {
		return nil, types.ExtractGRPCErrDetails(err)
	}
	createdAt, rerr := utils.GetTimebyTimestamp(res.CreatedAt)
	if rerr != nil {
		return nil, rerr
	}
	accessTokenExpireAt, rerr := utils.GetTimebyTimestamp(res.AccessTokenExpireAt)
	if rerr != nil {
		return nil, rerr
	}
	refreshTokenExpireAt, rerr := utils.GetTimebyTimestamp(res.RefreshTokenExpireAt)
	if rerr != nil {
		return nil, rerr
	}
	return &models.Token{
		AccessToken:          res.AccessToken,
		RefreshToken:         res.RefreshToken,
		UserId:               res.UserId,
		UserRole:             res.UserRole,
		SessionId:            res.SessionId,
		TokenStatus:          res.TokenStatus,
		Ip:                   res.Ip,
		Agent:                res.Agent,
		CreatedAt:            *createdAt,
		AccessTokenExpireAt:  *accessTokenExpireAt,
		RefreshTokenExpireAt: *refreshTokenExpireAt,
	}, nil
}
