package auth_test

import (
	"testing"

	"github.com/jairvercosa/auth-go/auth"
)

func TestSecureUpdatesValueAmount(t *testing.T) {
	pass_val := "Test123"
	pass := auth.Password{pass_val}
	pass.Secure()

	if pass.Value == pass_val {
		t.Errorf("Password.secure does not updates value amount")
	}
}

func TestVerifyReturnsTrueWhenPasswordIsTheSame(t *testing.T) {
	pass_val := "Test123"
	pass := auth.Password{pass_val}
	pass.Secure()

	if !pass.Verify(pass_val) {
		t.Errorf("Verify does not return true when password matches")
	}
}

func TestVerifyReturnsFalseWhenPasswordIsNotTheSame(t *testing.T) {
	pass_val := "Test123"
	pass := auth.Password{pass_val}
	pass.Secure()

	if pass.Verify("Teste") {
		t.Errorf("Verify does not return true when password matches")
	}
}
