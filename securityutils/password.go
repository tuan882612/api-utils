package securityutils

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	if len(password) == 0 {
		return "", errors.New("password is empty")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

func ValidatePassword(hash *string, password string) (bool, error) {
	if hash == nil {
		return false, errors.New("password hash is nil")
	}

	if len(password) == 0 {
		return false, errors.New("password is empty")
	}

	bHash, bPassword := []byte(*hash), []byte(password)
	return bcrypt.CompareHashAndPassword(bHash, bPassword) == nil, nil
}
