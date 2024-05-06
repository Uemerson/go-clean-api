package errors

import "net/http"

type Err struct {
	Data       interface{} `json:"data"`
	StatusCode int         `json:"-"`
}

func ErrInternalServerError() *Err {
	return &Err{
		Data:       "internal server error",
		StatusCode: http.StatusInternalServerError,
	}
}
