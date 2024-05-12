package handler

import (
	"github.com/Uemerson/clean-go-api/internal/domain/user/contracts"
)

type UserHandler struct {
	s contracts.Service
}

func NewUserHandler(s contracts.Service) *UserHandler {
	return &UserHandler{s: s}
}
