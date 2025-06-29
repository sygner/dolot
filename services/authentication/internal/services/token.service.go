package services

import (
	"dolott_authentication/internal/models"
	"dolott_authentication/internal/repository"
	"dolott_authentication/internal/types"
	"dolott_authentication/internal/utils"
	"neo/libs/idgen"
	"time"
)

type (
	TokenService interface {
		GetTokenByAccessTokenAndAgentAndUserId(string, string, int32) (*models.Token, *types.Error)
		GetTokensByUserId(*models.Pagination, int32) (*models.Tokens, *types.Error)
		GetTokens(*models.Pagination) (*models.Tokens, *types.Error)
		DeleteTokenBySessionIdAdnUserId(int32, int32) *types.Error
		VerifyToken(string, string) (*models.Token, *types.Error)
		RenewToken(*models.RenewTokenDTO) (*models.Token, *types.Error)
	}
	tokenService struct {
		repository repository.AuthenticationRepository
	}
)

func NewTokenService(repository repository.AuthenticationRepository) TokenService {
	return &tokenService{
		repository: repository,
	}
}

func (c *tokenService) GetTokenByAccessTokenAndAgentAndUserId(accessToken string, agent string, userId int32) (*models.Token, *types.Error) {
	return c.repository.GetTokenByAccessTokenAndAgentAndUserId(accessToken, agent, userId)
}

func (c *tokenService) GetTokensByUserId(data *models.Pagination, userId int32) (*models.Tokens, *types.Error) {
	res, err := c.repository.GetTokensByUserId(data, userId)
	if err != nil {
		return nil, err
	}

	var total *int32
	if data.Total {
		count, err := c.repository.GetTokensByUserIdCount(userId)
		if err != nil {
			return nil, err
		}

		total = &count
	}

	return &models.Tokens{Tokens: res, Total: total}, nil
}

func (c *tokenService) GetTokens(data *models.Pagination) (*models.Tokens, *types.Error) {
	res, err := c.repository.GetTokens(data)
	if err != nil {
		return nil, err
	}

	var total *int32
	if data.Total {
		count, err := c.repository.GetTokensCount()
		if err != nil {
			return nil, err
		}

		total = &count
	}

	return &models.Tokens{Tokens: res, Total: total}, nil
}

func (c *tokenService) DeleteTokenBySessionIdAdnUserId(sessionId int32, userId int32) *types.Error {
	return c.repository.DeleteTokenBySessionIdAdnUserId(sessionId, userId)
}

func (c *tokenService) VerifyToken(accessToken string, agent string) (*models.Token, *types.Error) {
	res, err := c.repository.GetTokenByAccessTokenAndAgent(accessToken, agent)
	if err != nil {
		return nil, err
	}

	err = utils.ValidateTokenStatus(res.TokenStatus)
	if err != nil {
		return nil, err
	}

	err = utils.ValidateTokenExpireTime(res.AccessTokenExpireAt)
	if err != nil {
		return nil, err
	}

	return res, nil

}

func (c *tokenService) RenewToken(renewToken *models.RenewTokenDTO) (*models.Token, *types.Error) {
	res, err := c.repository.GetTokenByAccessTokenAndRefreshToken(renewToken.AccessToken, renewToken.RefreshToken)
	if err != nil {
		return nil, err
	}
	if res.TokenStatus == "Permanent_Banned" {
		return nil, types.NewBadRequestError("token banned, error code #2031")

	}

	err = utils.ValidateTokenExpireTime(res.RefreshTokenExpireAt)
	if err != nil {
		return nil, err
	}

	result, err := c.createToken(renewToken.Ip, renewToken.Agent, res.UserId, res.UserRole, res.SessionId)
	if err != nil {
		return nil, err
	}
	err = c.repository.AddToken(result)
	if err != nil {
		return nil, err
	}

	err = c.repository.AddLoginHistory(res.UserId)
	if err != nil {
		return nil, err
	}

	err = c.repository.DeleteTokenByAccessToken(renewToken.AccessToken)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *tokenService) createToken(ip string, agent string, userID int32, userRole string, sessionId int32) (*models.Token, *types.Error) {

	accessToken, err := idgen.NextAlphanumericString(40)
	if err != nil {
		return nil, types.NewInternalError("internal issue, error code #2032")
	}
	refreshToken, err := idgen.NextAlphanumericString(40)
	if err != nil {
		return nil, types.NewInternalError("internal issue, error code #2033")
	}

	return &models.Token{
		AccessToken:          accessToken,
		RefreshToken:         refreshToken,
		UserId:               userID,
		UserRole:             userRole,
		SessionId:            sessionId,
		TokenStatus:          models.LIVE.String(),
		Ip:                   ip,
		Agent:                agent,
		CreatedAt:            time.Now(),
		AccessTokenExpireAt:  time.Now().Add(time.Minute * 15),
		RefreshTokenExpireAt: time.Now().AddDate(0, 1, 0),
	}, nil
}
