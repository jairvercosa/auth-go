package auth

import uuid "github.com/satori/go.uuid"

type Credential struct {
	Uuid     uuid.UUID
	Username string
	Email    string
	Active   bool
	password Password
}

func (c *Credential) SetPassword(val string) {
	c.password.Value = val
	c.password.Secure()
}

func (c *Credential) ValidatePassword(pass string) bool {
	return c.password.Verify(pass)
}
