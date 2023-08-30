package web

import (
	"go-authenticate-user-domain-centric-arch/infrastructure/web/controllers"
	"net/http"
)

type Router struct{}

func NewRouter(webController *controllers.Controller) {
	http.Handle("/api/auth", webController)
	// could be moved to a WebServer func
	http.ListenAndServe(":8080", nil)
}
