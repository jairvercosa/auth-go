package auth

import (
	"log"

	"github.com/raja/argon2pw"
)

type Password struct {
	Value string
}

func (p *Password) encrypt(val string) string {
	hashedPassword, err := argon2pw.GenerateSaltedHash(val)
	if err != nil {
		log.Fatalf("Unable to generate the password hash")
	}

	return hashedPassword
}

func (p *Password) Secure() {
	hash := p.encrypt(p.Value)
	p.Value = hash
}

func (p *Password) Verify(pass string) bool {
	valid, err := argon2pw.CompareHashWithPassword(p.Value, pass)
	if err != nil {
		return false
	}

	return valid
}
