package contracts

import (
	"github.com/Uemerson/clean-go-api/internal/configuration/exception"
)

type Encrypter interface {
	BCrypt(text string) (string, *exception.Exception)
}
