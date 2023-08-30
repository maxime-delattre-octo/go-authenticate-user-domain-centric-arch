package domain

import "go-authenticate-user-domain-centric-arch/domain/errors"

type User struct {
	name     string
	password string
}

func (user User) Authenticate(password string) error {
	if user.password != password {
		return errors.NewUnauthorizedOperationError()
	}

	return nil
}

func NewUser(name string, password string) *User {
	return &User{
		name:     name,
		password: password,
	}
}
