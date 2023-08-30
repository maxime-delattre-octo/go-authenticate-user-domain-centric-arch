package usecases

import "go-authenticate-user-domain-centric-arch/domain"

type UserRepository interface {
	Find(username string) *domain.User
}
