package model

type Repository interface {
	CreateAuth(auth *Auth) (int64, error)
}
