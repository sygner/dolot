package repository

import (
	"database/sql"
	"dolott_authentication/internal/models"
	"dolott_authentication/internal/types"
	"fmt"
)

func (c *authenticationRepository) AddToken(data *models.Token) *types.Error {
	query := `INSERT INTO "tokens" (access_token, refresh_token, user_id, user_role, session_id, token_status, ip, agent, created_at, access_token_expire_at, refresh_token_expire_at) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)`

	_, err := c.db.Exec(query, data.AccessToken, data.RefreshToken, data.UserId, data.UserRole, data.SessionId, data.TokenStatus, data.Ip, data.Agent, data.CreatedAt, data.AccessTokenExpireAt, data.RefreshTokenExpireAt)
	if err != nil {
		fmt.Println(err)
		return types.NewInternalError("internal issue, error code #1015")
	}

	return nil
}

func (c *authenticationRepository) DeleteUserTokens(userId int32) *types.Error {
	query := `DELETE FROM "tokens" WHERE user_id = $1`
	_, err := c.db.Exec(query, userId)
	if err != nil {
		return types.NewInternalError("internal issue, error code #1021")
	}
	return nil
}

func (c *authenticationRepository) DeleteTokenByAccessToken(accessToken string) *types.Error {
	query := `DELETE FROM "tokens" WHERE access_token = $1`
	_, err := c.db.Exec(query, accessToken)
	if err != nil {
		return types.NewInternalError("internal issue, error code #1022")
	}
	return nil
}

func (c *authenticationRepository) GetTokenByAccessToken(accessToken string) (*models.Token, *types.Error) {
	query := `SELECT access_token, refresh_token, user_id, user_role, session_id, token_status, ip, agent, created_at, access_token_expire_at, refresh_token_expire_at FROM "tokens" WHERE access_token = $1`
	var token models.Token
	err := c.db.QueryRow(query, accessToken).Scan(&token.AccessToken, &token.RefreshToken, &token.UserId, &token.UserRole, &token.SessionId, &token.TokenStatus, &token.Ip, &token.Agent, &token.CreatedAt, &token.AccessTokenExpireAt, &token.RefreshTokenExpireAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, types.NewNotFoundError("token not found, error code #1023")
		}
		return nil, types.NewInternalError("internal issue, error code #1024")
	}
	return &token, nil
}

func (c *authenticationRepository) GetTokenByAccessTokenAndAgent(accessToken string, agent string) (*models.Token, *types.Error) {
	query := `SELECT access_token, refresh_token, user_id, user_role, session_id, token_status, ip, agent, created_at, access_token_expire_at, refresh_token_expire_at FROM "tokens" WHERE access_token = $1 AND agent = $2`
	var token models.Token
	err := c.db.QueryRow(query, accessToken, agent).Scan(&token.AccessToken, &token.RefreshToken, &token.UserId, &token.UserRole, &token.SessionId, &token.TokenStatus, &token.Ip, &token.Agent, &token.CreatedAt, &token.AccessTokenExpireAt, &token.RefreshTokenExpireAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, types.NewNotFoundError("token not found, error code #1025")
		}
		return nil, types.NewInternalError("internal issue, error code #1026")
	}
	return &token, nil
}

func (c *authenticationRepository) GetTokenByAccessTokenAndAgentAndUserId(accessToken string, agent string, userId int32) (*models.Token, *types.Error) {
	query := `SELECT access_token, refresh_token, user_id, user_role, session_id, token_status, ip, agent, created_at, access_token_expire_at, refresh_token_expire_at FROM "tokens" WHERE access_token = $1 AND agent = $2 AND user_id = $3`
	var token models.Token
	err := c.db.QueryRow(query, accessToken, agent, userId).Scan(&token.AccessToken, &token.RefreshToken, &token.UserId, &token.UserRole, &token.SessionId, &token.TokenStatus, &token.Ip, &token.Agent, &token.CreatedAt, &token.AccessTokenExpireAt, &token.RefreshTokenExpireAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, types.NewNotFoundError("token not found, error code #1027")
		}
		return nil, types.NewInternalError("internal issue, error code #1028")
	}
	return &token, nil
}

func (c *authenticationRepository) GetTokensByUserId(data *models.Pagination, userId int32) ([]models.Token, *types.Error) {
	query := `SELECT access_token, refresh_token, user_id, user_role, session_id, token_status, ip, agent, created_at, access_token_expire_at, refresh_token_expire_at FROM "tokens" WHERE user_id = $1 OFFSET $2 LIMIT $3`

	rows, err := c.db.Query(query, userId, data.Offset, data.Limit)
	if err != nil {
		return nil, types.NewInternalError("internal issue, error code #1029")
	}

	defer rows.Close()
	tokens := make([]models.Token, 0)
	for rows.Next() {
		var token models.Token
		if err := rows.Scan(
			&token.AccessToken,
			&token.RefreshToken,
			&token.UserId,
			&token.UserRole,
			&token.SessionId,
			&token.TokenStatus,
			&token.Ip,
			&token.Agent,
			&token.CreatedAt,
			&token.AccessTokenExpireAt,
			&token.RefreshTokenExpireAt); err != nil {
			return nil, types.NewInternalError("internal issue, error code #1030")
		}
		tokens = append(tokens, token)
	}
	return tokens, nil
}

func (c *authenticationRepository) GetTokens(data *models.Pagination) ([]models.Token, *types.Error) {
	query := `SELECT access_token, refresh_token, user_id, user_role, session_id, token_status, ip, agent, created_at, access_token_expire_at, refresh_token_expire_at FROM "tokens" OFFSET $1 LIMIT $2`

	rows, err := c.db.Query(query, data.Offset, data.Limit)
	if err != nil {
		return nil, types.NewInternalError("internal issue, error code #1031")
	}

	defer rows.Close()
	tokens := make([]models.Token, 0)
	for rows.Next() {
		var token models.Token
		if err := rows.Scan(
			&token.AccessToken,
			&token.RefreshToken,
			&token.UserId,
			&token.UserRole,
			&token.SessionId,
			&token.TokenStatus,
			&token.Ip,
			&token.Agent,
			&token.CreatedAt,
			&token.AccessTokenExpireAt,
			&token.RefreshTokenExpireAt); err != nil {
			return nil, types.NewInternalError("internal issue, error code #1032")
		}
		tokens = append(tokens, token)
	}
	return tokens, nil
}

func (c *authenticationRepository) DeleteTokenBySessionIdAdnUserId(sessionId int32, userId int32) *types.Error {
	query := `DELETE FROM "tokens" WHERE session_id = $1 AND user_id = $2`
	_, err := c.db.Exec(query, sessionId, userId)
	if err != nil {
		return types.NewInternalError("internal issue, error code #1033")
	}
	return nil
}

func (c *authenticationRepository) GetTokenByAccessTokenAndRefreshToken(accessToken string, refreshToken string) (*models.Token, *types.Error) {
	query := `SELECT access_token, refresh_token, user_id, user_role, session_id, token_status, ip, agent, created_at, access_token_expire_at, refresh_token_expire_at FROM "tokens" WHERE access_token = $1 AND refresh_token = $2`
	var token models.Token
	err := c.db.QueryRow(query, accessToken, refreshToken).Scan(&token.AccessToken, &token.RefreshToken, &token.UserId, &token.UserRole, &token.SessionId, &token.TokenStatus, &token.Ip, &token.Agent, &token.CreatedAt, &token.AccessTokenExpireAt, &token.RefreshTokenExpireAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, types.NewNotFoundError("token not found, error code #1034")
		}
		return nil, types.NewInternalError("internal issue, error code #1035")
	}
	return &token, nil
}

func (c *authenticationRepository) GetTokensCount() (int32, *types.Error) {
	query := `SELECT count(*) FROM "tokens"`
	var totalCount int32
	err := c.db.QueryRow(query).Scan(&totalCount)
	if err != nil {
		return 0, types.NewInternalError("internal issue, error code #1036")
	}
	return totalCount, nil
}

func (c *authenticationRepository) GetTokensByUserIdCount(userId int32) (int32, *types.Error) {
	query := `SELECT count(*) FROM "tokens" WHERE user_id = $1`
	var totalCount int32
	err := c.db.QueryRow(query, userId).Scan(&totalCount)
	if err != nil {
		return 0, types.NewInternalError("internal issue, error code #1037")
	}
	return totalCount, nil
}
