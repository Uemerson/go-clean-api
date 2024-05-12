package service

import (
	"github.com/Uemerson/clean-go-api/internal/configuration/exception"
	"github.com/Uemerson/clean-go-api/internal/domain/user/model"
)

func (s *Service) Load() ([]model.User, *exception.Exception) {
	u, err := s.repo.Load()
	return u, err
}
