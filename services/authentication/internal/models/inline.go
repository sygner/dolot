package models

import "time"

type Pagination struct {
	Offset int32 `json:"offset"`
	Limit  int32 `json:"limit"`
	Total  bool  `json:"total"`
}

type Users struct {
	Users []User `json:"users"`
	Total *int32 `json:"total"`
}

type Tokens struct {
	Tokens []Token `json:"tokens"`
	Total  *int32  `json:"total"`
}

type LoginHistories struct {
	LoginHistories []time.Time `json:"login_histories"`
	Total          *int32      `json:"total"`
}

type VerifyDTO struct {
	Code         string  `json:"code"`
	Agent        string  `json:"agent"`
	Ip           string  `json:"ip"`
	VerifyMethod int32   `json:"verify_method"`
	NewPassword  *string `json:"new_password"`
}

type VerifyResponse struct {
	Token Token  `json:"token"`
	Value string `json:"value"`
}
