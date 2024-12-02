package models

import (
	"fmt"
	"time"
)

type User struct {
	UserId          int32     `json:"user_id"`
	PhoneNumber     string    `json:"phone_number"`
	AccountUsername string    `json:"account_username"`
	Email           string    `json:"email"`
	UserRole        string    `json:"user_role"`
	UserStatus      string    `json:"user_status"`
	Provider        string    `json:"provider,omitempty"`
	IsSSO           bool      `json:"is_sso,omitempty"`
	CreatedAt       time.Time `json:"created_at"`
}

type UserRole int

// Declare user roles using iota
const (
	ADMIN UserRole = iota
	USER
	MODERATOR
)

// String method to get the string representation of the UserRole
func (role UserRole) String() string {
	switch role {
	case ADMIN:
		return fmt.Sprintf("Admin %s", time.Now().Format("2006-01-02 15:04:05"))
	case USER:
		return fmt.Sprintf("User %s", time.Now().Format("2006-01-02 15:04:05"))
	case MODERATOR:
		return fmt.Sprintf("Moderator %s", time.Now().Format("2006-01-02 15:04:05"))
	default:
		return fmt.Sprintf("Unkown %s", time.Now().Format("2006-01-02 15:04:05"))
	}
}

type UserStatus int

// Declare user roles using iota
const (
	ONGOING UserStatus = iota
	SUSPENDED
	BANNED
)

// String method to get the string representation of the UserRole
func (role UserStatus) String() string {
	switch role {
	case ONGOING:
		return fmt.Sprintf("OnGoing %s", time.Now().Format("2006-01-02 15:04:05"))
	case SUSPENDED:
		return fmt.Sprintf("Suspended %s", time.Now().Format("2006-01-02 15:04:05"))
	case BANNED:
		return fmt.Sprintf("Banned %s", time.Now().Format("2006-01-02 15:04:05"))
	default:
		return fmt.Sprintf("Unkown %s", time.Now().Format("2006-01-02 15:04:05"))
	}
}
