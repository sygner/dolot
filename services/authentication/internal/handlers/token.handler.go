package handlers

import (
	"context"
	"dolott_authentication/internal/models"
	"dolott_authentication/internal/services"
	pb "dolott_authentication/proto/api"
)

type TokenHandler struct {
	pb.UnimplementedTokenServiceServer
	tokenService services.TokenService
}

func NewTokenHandler(tokenService services.TokenService) *TokenHandler {
	return &TokenHandler{
		tokenService: tokenService,
	}
}

func (c *TokenHandler) GetTokenByATAU(ctx context.Context, request *pb.GetTokenByATAURequest) (*pb.Token, error) {
	res, err := c.tokenService.GetTokenByAccessTokenAndAgentAndUserId(request.AccessToken, request.Agent, request.UserId)
	if err != nil {
		return nil, err.ErrorToGRPCStatus()
	}

	return tokenToProto(res), nil
}

func (c *TokenHandler) GetTokensByUserId(ctx context.Context, request *pb.GetTokensByUserIdRequest) (*pb.Tokens, error) {

	data := &models.Pagination{
		Offset: request.Pagination.Offset,
		Limit:  request.Pagination.Limit,
		Total:  request.Pagination.GetTotal,
	}

	res, err := c.tokenService.GetTokensByUserId(data, request.UserId)
	if err != nil {
		return nil, err.ErrorToGRPCStatus()
	}

	return tokenToProtoList(res), nil
}

func (c *TokenHandler) GetTokens(ctx context.Context, request *pb.Pagination) (*pb.Tokens, error) {

	data := &models.Pagination{
		Offset: request.Offset,
		Limit:  request.Limit,
		Total:  request.GetTotal,
	}

	res, err := c.tokenService.GetTokens(data)
	if err != nil {
		return nil, err.ErrorToGRPCStatus()
	}

	return tokenToProtoList(res), nil
}

func (c *TokenHandler) VerifyToken(ctx context.Context, request *pb.VerifyTokenRequest) (*pb.Token, error) {
	res, err := c.tokenService.VerifyToken(request.AccessToken, request.Agent)
	if err != nil {
		return nil, err.ErrorToGRPCStatus()
	}

	return tokenToProto(res), nil
}

func (c *TokenHandler) RenewToken(ctx context.Context, request *pb.RenewTokenRequest) (*pb.Token, error) {
	data := &models.RenewTokenDTO{
		AccessToken:  request.AccessToken,
		RefreshToken: request.RefreshToken,
		Agent:        request.Agent,
		Ip:           request.Ip,
	}
	res, err := c.tokenService.RenewToken(data)
	if err != nil {
		return nil, err.ErrorToGRPCStatus()
	}
	return tokenToProto(res), nil
}

func tokenToProtoList(tokenData *models.Tokens) *pb.Tokens {
	tokens := make([]*pb.Token, 0, len(tokenData.Tokens))
	for _, token := range tokenData.Tokens {
		tokens = append(tokens, tokenToProto(&token))
	}

	return &pb.Tokens{Tokens: tokens, Total: tokenData.Total}
}

func tokenToProto(up *models.Token) *pb.Token {
	return &pb.Token{
		AccessToken:          up.AccessToken,
		RefreshToken:         up.RefreshToken,
		TokenStatus:          up.TokenStatus,
		Ip:                   up.Ip,
		Agent:                up.Agent,
		SessionId:            up.SessionId,
		UserId:               up.UserId,
		UserRole:             up.UserRole,
		CreatedAt:            up.CreatedAt.Unix(),
		AccessTokenExpireAt:  up.AccessTokenExpireAt.Unix(),
		RefreshTokenExpireAt: up.RefreshTokenExpireAt.Unix(),
	}
}
