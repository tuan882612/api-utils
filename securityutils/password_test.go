package securityutils

import "testing"

func Test_HashPassword(t *testing.T) {
	password := "password"
	hashedPassword, err := HashPassword(password)

	if err != nil {
		t.Errorf("Error hashing password: %v", err)
	}

	cond, err := ValidatePassword(&hashedPassword, password)

	if err != nil {
		t.Errorf("Error comparing password: %v", err)
	}

	if cond != true {
		t.Errorf("Error comparing password: %v", err)
	}
}

func Test_ValidatePassword(t *testing.T) {
	hashed := "$2a$10$gEH4ShlEMylHTjYuSU7NReC0OIXdFKBFfIDPzFNYreXUFOazKTN1O"

	cond, err := ValidatePassword(&hashed, "password")

	if err != nil {
		t.Errorf("Error comparing password: %v", err)
	}

	if cond != true {
		t.Errorf("Error comparing password: %v", err)
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
	var hash *string

	_, err := ValidatePassword(hash, "password")
	if err == nil {
		t.Errorf("Should return an error: %v", err)
	}

	_, err = ValidatePassword(hash, "")
	if err == nil {
		t.Errorf("Should return an error: %v", err)
	}
}
