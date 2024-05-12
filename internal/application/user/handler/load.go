package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Uemerson/clean-go-api/internal/configuration/exception"
	"github.com/Uemerson/clean-go-api/internal/domain/user/model"
)

func (uh *UserHandler) Load(w http.ResponseWriter, r *http.Request) {
	users, err := uh.s.Load()
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(err.StatusCode)
		json.NewEncoder(w).Encode(err)
		return
	}
	if users == nil {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode([]model.User{})
		return
	}
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(users); err != nil {
		err := exception.InternalServerError()
		w.WriteHeader(err.StatusCode)
		json.NewEncoder(w).Encode(err)
	}
}
