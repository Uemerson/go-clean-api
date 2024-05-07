package repository

import "github.com/Uemerson/clean-go-api/internal/domain/user/infra/repository/entity"

func (r *Repository) Migration() {
	r.db.AutoMigrate(&entity.User{})
	r.db.Migrator().CreateTable()
}
