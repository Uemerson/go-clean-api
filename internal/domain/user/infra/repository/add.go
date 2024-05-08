package repository

import (
	"github.com/Uemerson/clean-go-api/internal/configuration/exception"
	"github.com/Uemerson/clean-go-api/internal/domain/user/infra/repository/entity"
	"github.com/Uemerson/clean-go-api/internal/domain/user/model"
)

func (r *Repository) Add(u *model.User) *exception.Exception {
	result := r.db.Create(&entity.User{
		Name:      u.Name,
		Email:     u.Email,
		Password:  u.Password,
		CreatedAt: u.CreatedAt,
	})
	if result.Error != nil {
		return exception.InternalServerError()
	}
	return nil
}
