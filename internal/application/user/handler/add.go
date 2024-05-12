package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Uemerson/clean-go-api/internal/configuration/exception"
)

type UserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (uh *UserHandler) Add(w http.ResponseWriter, r *http.Request) {
	var ur UserRequest
	w.Header().Set("Content-Type", "application/json")
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
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		err := exception.InternalServerError()
		w.WriteHeader(err.StatusCode)
		json.NewEncoder(w).Encode(err)
	}
}
