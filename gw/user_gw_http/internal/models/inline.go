package models

type Pagination struct {
	Offset int32 `json:"offset"`
	Limit  int32 `json:"limit"`
	Total  bool  `json:"total,omitempty"`
}

type Users struct {
	Users []User `json:"users"`
	Total *int32 `json:"total,omitempty"`
}

type Tokens struct {
	Tokens []Token `json:"tokens"`
	Total  *int32  `json:"total,omitempty"`
}

type LoginHistories struct {
	LoginHistories []int64 `json:"login_histories"`
	Total          *int32  `json:"total,omitempty"`
}

type VerifyResponse struct {
	Token Token  `json:"token"`
	Value string `json:"value"`
}
