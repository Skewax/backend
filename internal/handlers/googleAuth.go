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
	switch {
	case result.OK():
		//TODO: get user data from google on OK
		//TODO: add access and refresh tokens both to cache and database
		name := "TEMPNAME"
		imageUrl := "TEMPURL"
		return authentication.NewGoogleLoginOK().WithPayload(&models.LoginResponse{
			Error:        "",
			SessionToken: result.SessionToken(),
			Timeout:      result.GoogleTimeout(),
			User: &models.User{
				Name:  &name,
				Image: &imageUrl,
			},
		})
	default:
		model := models.BasicResponse{
			Error: result.GetErrorDesc(),
		}
		switch result.GetErrorCode() {
		case "400":
			return authentication.NewGoogleLoginBadRequest().WithPayload(&model)
		case "401":
			return authentication.NewGoogleLoginUnauthorized().WithPayload(&model)
		case "408":
			return authentication.NewGoogleLoginRequestTimeout().WithPayload(&model)
		default:
			return authentication.NewGoogleLoginInternalServerError().WithPayload(&model)
		}
	}
}
