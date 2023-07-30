package securityutils

import "testing"

func Test_HashPassword(t *testing.T) {
	password := "password"
	hashedPassword, err := HashPassword(password)

	if err != nil {
		t.Errorf("Error hashing password: %v", err)
	}

	if err = ValidatePassword(hashedPassword, password); err != nil {
		t.Errorf("Error validating password: %v", err)
	}

}

func Test_ValidatePassword(t *testing.T) {
	hashed := "$2a$10$gEH4ShlEMylHTjYuSU7NReC0OIXdFKBFfIDPzFNYreXUFOazKTN1O"

	if err := ValidatePassword(hashed, "password"); err != nil {
		t.Errorf("Error validating password: %v", err)
	}

}

func Test_HashPassword_ErrorHandling(t *testing.T) {
	password := ""
	_, err := HashPassword(password)

	if err == nil {
		t.Errorf("Should return an error: %v", err)
	}
}

func Test_ValidatePassword_ErrorHandling(t *testing.T) {
	var hash string

	if err := ValidatePassword(hash, "password"); err == nil {
		t.Errorf("Should return an error: %v", err)
	}

	if err := ValidatePassword(hash, ""); err == nil {
		t.Errorf("Should return an error: %v", err)
	}
}
