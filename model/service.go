package model

type Service interface {
	Signup(auth *Auth) (*SignupResponse, error)
}

type SignupResponse struct {
	ID           int64  `json:"id"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
