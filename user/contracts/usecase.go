package contracts

import "github.com/Uemerson/clean-go-api/user/model"

type UseCase interface {
	Add(name, email, password string) (*model.User, error)
}
