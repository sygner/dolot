package models

type ResetPasswordDTO struct {
	UserId          int32
	CurrentPassword string `json:"current_password"`
	NewPassword     string `json:"new_password"`
	Agent           string `json:"agent"`
	Ip              string `json:"ip"`
}
