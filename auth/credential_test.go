package auth_test

import (
	"testing"

	"github.com/jairvercosa/auth-go/auth"
)

func TestSetPasswordUpdatePassword(t *testing.T) {
	password := "Test123"

	cred := auth.Credential{}
	cred.SetPassword(password)
	if cred.ValidatePassword(password) != true {
		t.Errorf("Credencial.SetPassword did not update password")
	}
}

func TestVerifyPasswordReturnsFalseWhenItDoesNotMatche(t *testing.T) {
	cred := auth.Credential{}
	cred.SetPassword("Test123")
	if cred.ValidatePassword("Test1234") == true {
		t.Errorf(
			"Credencial.ValidatePassowrd did not return false " +
				"when password does not match")
	}
}
