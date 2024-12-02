package models

import (
	"time"
)

type Token struct {
	AccessToken          string    `json:"access_token"`
	RefreshToken         string    `json:"refresh_token"`
	UserId               int32     `json:"user_id"`
	UserRole             string    `json:"user_role"`
	SessionId            int32     `json:"session_id"`
	TokenStatus          string    `json:"token_status"`
	Ip                   string    `json:"ip"`
	Agent                string    `json:"agent"`
	CreatedAt            time.Time `json:"created_at"`
	AccessTokenExpireAt  time.Time `json:"access_token_expire_at"`
	RefreshTokenExpireAt time.Time `json:"refresh_token_expire_at"`
}

type TokenStatus int

// Declare user roles using iota
const (
	LIVE TokenStatus = iota
	DEAD
	PERMANENT_BANNED
)

// String method to get the string representation of the UserRole
func (status TokenStatus) String() string {
	switch status {
	case LIVE:
		return "Live"
	case DEAD:
		return "Dead"
	case PERMANENT_BANNED:
		return "Permanent_Banned"
	default:
		return "Unkown"
	}
}
