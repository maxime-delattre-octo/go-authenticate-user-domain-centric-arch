package controllers

import (
	"encoding/json"
	"go-authenticate-user-domain-centric-arch/usecases"
	"net/http"
)

type Controller struct {
	AuthenticateUser *usecases.AuthenticateUser
}

func (controller *Controller) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	var authenticateUserBodyRequest AuthenticateUserBodyRequest
	json.NewDecoder(request.Body).Decode(&authenticateUserBodyRequest)

	// TODO could be move into func of AuthenticateUserBodyRequest
	if authenticateUserBodyRequest.Name == "" || authenticateUserBodyRequest.Password == "" {
		http.Error(writer, "", http.StatusUnauthorized)
		return
	}

	token, authenticateError := controller.AuthenticateUser.Execute(authenticateUserBodyRequest.Name, authenticateUserBodyRequest.Password)

	if authenticateError != nil {
		http.Error(writer, "", http.StatusUnauthorized)
		return
	}

	tokenJsonBodyResponse, _ := json.Marshal(token)

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(tokenJsonBodyResponse)
}
