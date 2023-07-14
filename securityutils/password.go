package securityutils

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (*string, error) {
	if len(password) == 0 {
		return nil, errors.New("password is empty")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	strHash := string(hash)
	return &strHash, err
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
