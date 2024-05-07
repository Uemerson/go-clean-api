package repository

import (
	"errors"

	"github.com/Uemerson/clean-go-api/internal/configuration/exception"
	"github.com/Uemerson/clean-go-api/internal/domain/user/infra/repository/entity"
	"github.com/Uemerson/clean-go-api/internal/domain/user/model"
	"gorm.io/gorm"
)

func (r *Repository) LoadByEmail(email string) (*model.User, *exception.Exception) {
	var u entity.User
	result := r.db.First(&u, "email = ?", email)
	switch {
	case errors.Is(result.Error, gorm.ErrRecordNotFound):
		return nil, nil
	case result.Error != nil:
		return nil, exception.InternalServerError()
	}
	return &model.User{
		Id:        u.Id,
		Name:      u.Name,
		Email:     u.Email,
		Password:  u.Password,
		CreatedAt: u.CreatedAt}, nil
}
