package contracts

import (
	"github.com/Uemerson/clean-go-api/internal/configuration/exception"
	"github.com/Uemerson/clean-go-api/internal/domain/user/model"
)

type Service interface {
	Add(name, email, password string) (*model.User, *exception.Exception)
	Load() ([]model.User, *exception.Exception)
}
