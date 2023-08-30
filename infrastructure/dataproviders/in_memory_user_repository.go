package dataproviders

import "go-authenticate-user-domain-centric-arch/domain"
import _ "go-authenticate-user-domain-centric-arch/usecases"

type InMemoryUserRepository struct{}

func (inMemoryUserRepository *InMemoryUserRepository) Find(username string) *domain.User {
	if username == "dexter" {
		return domain.NewUser("dexter", "killer")
	}

	return nil
}
