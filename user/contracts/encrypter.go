package contracts

type Encrypter interface {
	BCrypt(text string) (string, error)
}
