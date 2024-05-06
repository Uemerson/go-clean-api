package exception

import "net/http"

type Exception struct {
	Data       interface{} `json:"data"`
	StatusCode int         `json:"-"`
}

func InternalServerError() *Exception {
	return &Exception{
		Data:       "internal server error",
		StatusCode: http.StatusInternalServerError,
	}
}

func BadRequest(data interface{}) *Exception {
	return &Exception{
		Data:       data,
		StatusCode: http.StatusBadRequest,
	}
}

func Conflict(data interface{}) *Exception {
	return &Exception{
		Data:       data,
		StatusCode: http.StatusConflict,
	}
}
