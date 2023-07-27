package model

type Auth struct {
	Username     string `json:"username"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	Country      string `json:"country"`
	RefreshToken string
}
