package service

import (
	"github.com/tckthecreator/clean_arch_go/model"
	serviceutils "github.com/tckthecreator/clean_arch_go/service/utils"
)

type service struct {
	repo model.Repository
}

func NewService(repo model.Repository) model.Service {
	return &service{
		repo,
	}
}

func (s *service) Signup(auth *model.Auth) (*model.SignupResponse, error) {
	rt, err := serviceutils.GenerateRefreshToken()
	if err != nil {
		return nil, err
	}

	id, err := s.repo.CreateAuth(auth)
	if err != nil {
		return nil, err
	}

	at, err := serviceutils.GenerateAccessToken(id)
	if err != nil {
		return nil, err
	}

	return &model.SignupResponse{
		ID:           id,
		RefreshToken: rt,
		AccessToken:  at,
	}, err
}
