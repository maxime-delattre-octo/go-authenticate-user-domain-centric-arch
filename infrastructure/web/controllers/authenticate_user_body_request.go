package controllers

type AuthenticateUserBodyRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}
