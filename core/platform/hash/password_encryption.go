package hash

import (
	"attempt4/core/internal"
	"golang.org/x/crypto/bcrypt"
)

func EncryptPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)

	if err != nil {
		return "", internal.FailInHash
	}
	return string(hash), nil
}
func CompareEncryptedPasswords(password1 string, password2 string) error {
	err := bcrypt.CompareHashAndPassword([]byte(password1), []byte(password2))
	if err != nil {
		return internal.InvalidPassword
	}
	return nil
}
