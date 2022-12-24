package handlers

import (
	skewauth "github.com/Skewax/backend/internal/auth"
	"github.com/Skewax/backend/pkg/swagger/server/models"
	"github.com/Skewax/backend/pkg/swagger/server/restapi/operations/authentication"
	"github.com/go-openapi/runtime/middleware"
)

func HandleTokenLogin(tokenChan chan *skewauth.UseTokenReq, params authentication.TokenLoginParams) middleware.Responder {
	reqChan := make(chan skewauth.AuthToken)
	tokenChan <- &skewauth.UseTokenReq{
		*params.Body.UserID,
		*params.Body.SessionToken,
		true,
		reqChan,
	}
	result := <-reqChan
	name := result.SessionToken()
	image := result.GoogleToken()
	body := models.LoginResponse{Error: "", SessionToken: "sessionToken", Timeout: 100, User: &models.User{Name: &name, Image: &image}}
	return authentication.NewTokenLoginOK().WithPayload(&body)
}

func HandleLogout(tokenChan chan *skewauth.ClearTokenReq, params authentication.LogoutParams) middleware.Responder {
	body := models.BasicResponse{Error: ""}
	return authentication.NewLogoutOK().WithPayload(&body)
}
