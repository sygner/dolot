package utils

import (
	"dolott_authentication/internal/types"
	"fmt"
	"regexp"
	"strings"
	"time"
)

// ValidatePassword checks if the password meets the specified criteria
func ValidatePassword(password string) (bool, *types.Error) {
	// Check the length
	if len(password) < 8 {
		return false, types.NewBadRequestError("password must be at least 8 characters long, error code #9011")
	}

	// Check for at least one lowercase letter
	hasLower, err := regexp.MatchString(`[a-z]`, password)
	if err != nil {
		fmt.Println(err)
		return false, types.NewInternalError("internal issue, error code #9007")
	}
	if !hasLower {
		return false, types.NewBadRequestError("password must contain at least one lowercase letter, error code #9012")
	}

	// Check for at least one uppercase letter
	hasUpper, err := regexp.MatchString(`[A-Z]`, password)
	if err != nil {
		fmt.Println(err)
		return false, types.NewInternalError("internal issue, error code #9008")
	}
	if !hasUpper {
		return false, types.NewBadRequestError("password must contain at least one uppercase letter, error code #9013")
	}

	// Check for at least one digit
	hasDigit, err := regexp.MatchString(`[0-9]`, password)
	if err != nil {
		fmt.Println(err)
		return false, types.NewInternalError("internal issue, error code #9009")
	}
	if !hasDigit {
		return false, types.NewBadRequestError("password must contain at least one digit, error code #9014")
	}

	// Check for at least one special character
	hasSpecial, err := regexp.MatchString(`[\W_]`, password)
	if err != nil {
		fmt.Println(err)
		return false, types.NewInternalError("internal issue, error code 9010")
	}
	if !hasSpecial {
		return false, types.NewBadRequestError("password must contain at least one special character, error code #9015")
	}

	// If all checks pass, the password is valid
	return true, nil
}

func ValidateUserStatus(userStatus string) *types.Error {
	data := strings.Split(userStatus, " ")[0]
	if data == "Suspended" {
		return types.NewPermissionDeniedError("account suspended, try to open a ticket, error code #9016")
	} else if data == "Banned" {
		return types.NewPermissionDeniedError("account banned, error code #9017")
	}
	return nil
}

func ValidateTokenStatus(tokenStatus string) *types.Error {
	data := strings.Split(tokenStatus, " ")[0]
	if data == "Dead" {
		return types.NewPermissionDeniedError("token is dead, error code #9018")
	} else if data == "Permanent_Banned" {
		return types.NewPermissionDeniedError("token banned, error code #9019")
	}
	return nil
}

func ValidateTokenExpireTime(accessTokenExpireAt time.Time) *types.Error {
	if time.Now().After(accessTokenExpireAt) {
		return types.NewBadRequestError("token expired, error code #9029")
	}

	return nil
}
