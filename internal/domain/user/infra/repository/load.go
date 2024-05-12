package repository

import (
	"github.com/Uemerson/clean-go-api/internal/configuration/exception"
	"github.com/Uemerson/clean-go-api/internal/domain/user/infra/repository/entity"
	"github.com/Uemerson/clean-go-api/internal/domain/user/model"
)

func (r *Repository) Load() ([]model.User, *exception.Exception) {
	var users []entity.User
	result := r.db.Find(&users)
	if result.Error != nil {
		return nil, exception.InternalServerError()
	}
	var u []model.User
	for _, r := range users {
		u = append(u, model.User{
			Id:        r.Id,
			Name:      r.Name,
			Email:     r.Email,
			CreatedAt: r.CreatedAt,
		})
	}
	return u, nil
}
