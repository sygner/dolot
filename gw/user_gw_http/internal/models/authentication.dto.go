package models

type SignupDTO struct {
	Email    string  `json:"email"`
	Password *string `json:"password,omitempty"`
	Provider *string `json:"provider,omitempty"`
	Username *string `json:"username,omitempty"`
	IsSSO    bool    `json:"is_sso"`
	Ip       string
	Agent    string
}

type SigninDTO struct {
	Value         string  `json:"value"`
	Signin_Method int32   `json:"signin_method"`
	Provider      *string `json:"provider,omitempty"`
	IsSSO         bool    `json:"is_sso"`
	Ip            string
	Agent         string
	Password      *string `json:"password,omitempty"`
}

type VerifyDTO struct {
	Code          string  `json:"code"`
	Verify_Method int32   `json:"verify_method"`
	NewPassword   *string `json:"new_password,omitempty"`
	// IsSSO         bool    `json:"is_sso,omitempty"`
	Ip    string
	Agent string
}
