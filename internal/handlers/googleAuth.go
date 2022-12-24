package handlers

import (
	skewauth "github.com/Skewax/backend/internal/auth"
	"github.com/Skewax/backend/pkg/swagger/server/models"
	"github.com/Skewax/backend/pkg/swagger/server/restapi/operations/authentication"

	"github.com/go-openapi/runtime/middleware"
)

func HandleGoogleLogin(tokenChan chan *skewauth.GenTokenReq, params authentication.GoogleLoginParams) middleware.Responder {
	reqChan := make(chan skewauth.AuthToken)
	tokenChan <- &skewauth.GenTokenReq{
		AuthCode:     *params.Body.Code,
		GrantType:    *params.Body.Prompt,
		HostedDomain: *params.Body.HostedDomain,
		Result:       reqChan,
	}
	result := <-reqChan
	if !result.OK() {
		return authentication.NewGoogleLoginBadRequest().WithPayload(&models.BasicResponse{result.GetErrorCode() + ": " + result.GetErrorDesc()})
	}
	name := result.GoogleToken()
	image := result.SessionToken()
	return authentication.NewGoogleLoginOK().WithPayload(&models.LoginResponse{"", "token", 100, &models.User{&name, &image}})
}
