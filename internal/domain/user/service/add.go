package service

import (
	"fmt"

	"github.com/Uemerson/clean-go-api/internal/configuration/exception"
	"github.com/Uemerson/clean-go-api/internal/domain/user/model"
)

func (s *Service) Add(name, email, password string) (*model.User, *exception.Exception) {
	p := []string{}
	for r, v := range map[string]string{
		"name":     name,
		"email":    email,
		"password": password,
	} {
		if v == "" {
			p = append(p, fmt.Sprintf("missing param: %s", r))
		}
	}
	if len(p) > 0 {
		return nil, exception.BadRequest(p)
	}
	if u, _ := s.repo.LoadByEmail(email); u != nil {
		return nil, exception.Conflict("email already exist")
	}
	hashedPassword, err := s.encrypt.BCrypt(password)
	if err != nil {
		return nil, err
	}
	u := model.User{
		Id:       s.helper.UUID(),
		Name:     name,
		Email:    email,
		Password: hashedPassword,
	}
	if err := s.repo.Add(&u); err != nil {
		return nil, err
	}
	return &model.User{
		Id:       u.Id,
		Name:     u.Name,
		Email:    u.Email,
		Password: u.Password,
	}, nil
}
