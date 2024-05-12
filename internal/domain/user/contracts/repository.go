package contracts

import (
	"github.com/Uemerson/clean-go-api/internal/configuration/exception"
	"github.com/Uemerson/clean-go-api/internal/domain/user/model"
)

type Repository interface {
	Add(u *model.User) *exception.Exception
	LoadByEmail(email string) (*model.User, *exception.Exception)
	Load() ([]model.User, *exception.Exception)
	Migration()
}
