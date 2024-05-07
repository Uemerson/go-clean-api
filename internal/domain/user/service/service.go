package service

import (
	"github.com/Uemerson/clean-go-api/internal/domain/user/contracts"
)

type Service struct {
	repo    contracts.Repository
	encrypt contracts.Encrypter
	helper  contracts.Helper
}

func NewService(r contracts.Repository, e contracts.Encrypter, h contracts.Helper) *Service {
	return &Service{
		repo:    r,
		encrypt: e,
		helper:  h,
	}
}
