package contracts

import "github.com/Uemerson/clean-go-api/user/entity"

type Repository interface {
	Add(u *entity.UserDto) error
	LoadByEmail(email string) (*entity.UserDto, error)
}
