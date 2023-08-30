package main

import (
	"go-authenticate-user-domain-centric-arch/infrastructure/dataproviders"
	"go-authenticate-user-domain-centric-arch/infrastructure/web"
	"go-authenticate-user-domain-centric-arch/infrastructure/web/controllers"
	"go-authenticate-user-domain-centric-arch/usecases"
)

// dependency injection
func main() {
	inMemoryRepository := dataproviders.InMemoryUserRepository{}
	random32CharsGenerator := dataproviders.GoRandom32CharsGenerator{}
	authenticateUser := usecases.AuthenticateUser{&inMemoryRepository, &random32CharsGenerator}
	webController := controllers.Controller{&authenticateUser}
	web.NewRouter(&webController)
}
