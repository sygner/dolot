package models

type UserDTO struct {
	Email           string  `json:"email"`
	UserId          int32   `json:"user_id"`
	UserRole        string  `json:"user_role"`
	UserStatus      string  `json:"user_status"`
	AccountUsername string  `json:"account_username"`
	Password        *string `json:"password"`
	Provider        *string `json:"provider"`
	IsSSO           bool    `json:"is_sso"`
	Agent           string  `json:"agent"`
	Ip              string  `json:"ip"`
}

type LoginDTO struct {
	Value       string  `json:"value"`
	Password    string  `json:"password"`
	Agent       string  `json:"agent"`
	Ip          string  `json:"ip"`
	Provider    *string `json:"provider"`
	IsSSO       bool    `json:"is_sso"`
	LoginMethod int32   `json:"login_method"`
}

type ResetPasswordDTO struct {
	UserId          int32  `json:"user_id"`
	CurrentPassword string `json:"current_password"`
	NewPassword     string `json:"new_password"`
	Agent           string `json:"agent"`
	Ip              string `json:"ip"`
}
