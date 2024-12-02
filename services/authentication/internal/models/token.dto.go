package models

type TokenDTO struct {
	UserId   string `json:"user_id"`
	UserRole string `json:"user_role"`
	Agent    string `json:"agent"`
	Ip       string `json:"ip"`
}

type VerifyTokenDTO struct {
	AccessToken string `json:"access_token"`
	UserId      int32  `json:"user_id"`
	Agent       string `json:"agent"`
}

type RenewTokenDTO struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Agent        string `json:"agent"`
	Ip           string `json:"ip"`
}
