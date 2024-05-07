package encrypter

import (
	"github.com/Uemerson/clean-go-api/internal/configuration/exception"
	"golang.org/x/crypto/bcrypt"
)

func (e *Encrypter) BCrypt(text string) (string, *exception.Exception) {
	hashedText, err := bcrypt.GenerateFromPassword([]byte(text), bcrypt.DefaultCost)
	if err != nil {
		return "", exception.InternalServerError()
	}
	return string(hashedText), nil
}
