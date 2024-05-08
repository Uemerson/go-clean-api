package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Uemerson/clean-go-api/internal/configuration/exception"
	"github.com/Uemerson/clean-go-api/internal/domain/user/contracts"
)

type UserHandler struct {
	s contracts.Service
}

func NewUserHandler(s contracts.Service) *UserHandler {
	return &UserHandler{s: s}
}

type UserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (uh *UserHandler) Add(w http.ResponseWriter, r *http.Request) {
	var ur UserRequest
	if err := json.NewDecoder(r.Body).Decode(&ur); err != nil {
		err := exception.InternalServerError()
		w.WriteHeader(err.StatusCode)
		json.NewEncoder(w).Encode(err)
		return
	}
	user, err := uh.s.Add(ur.Name, ur.Email, ur.Password)
	if err != nil {
		w.WriteHeader(err.StatusCode)
		json.NewEncoder(w).Encode(err)
		return
	}
	if err := json.NewEncoder(w).Encode(user); err != nil {
		err := exception.InternalServerError()
		w.WriteHeader(err.StatusCode)
		json.NewEncoder(w).Encode(err)
	}
}
