package usecases

import (
	"go-authenticate-user-domain-centric-arch/domain"
	"go-authenticate-user-domain-centric-arch/domain/errors"
	"go-authenticate-user-domain-centric-arch/usecases/ports"
)

type AuthenticateUser struct {
	UserRepository         ports.UserRepository
	Random32CharsGenerator ports.Random32CharsGenerator
}

func (authenticateUser *AuthenticateUser) Execute(username string, userPassword string) (*domain.Token, error) {
	user := authenticateUser.UserRepository.Find(username)

	if user == nil {
		return nil, &errors.UnauthorizedOperationError{}
	}

	authenticateError := user.Authenticate(userPassword)

	if authenticateError != nil {
		return nil, authenticateError
	}

	random32Chars := authenticateUser.Random32CharsGenerator.Execute()

	token := domain.NewToken(random32Chars)

	return token, nil
}
