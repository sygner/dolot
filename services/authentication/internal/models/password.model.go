package models

type UserPassword struct {
	UserId   int32  `json:"user_id"`
	Password string `json:"password"`
}
