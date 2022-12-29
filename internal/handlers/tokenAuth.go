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
		Uid:        *params.Body.UserID,
		Token:      *params.Body.SessionToken,
		CanReplace: true,
		Result:     reqChan,
	}
	result := <-reqChan
	switch {
	case result.OK():
		//TODO remove validated token, create and cache new
		//TODO get google data
		name := "TEMPNAME"
		imageUrl := "TEMPURL"
		return authentication.NewTokenLoginOK().WithPayload(&models.LoginResponse{
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
			return authentication.NewTokenLoginBadRequest().WithPayload(&model)
		default:
			return authentication.NewTokenLoginInternalServerError().WithPayload(&model)
		}
	}
}

func HandleLogout(tokenChan chan *skewauth.ClearTokenReq, params authentication.LogoutParams) middleware.Responder {
	body := models.BasicResponse{Error: ""}
	return authentication.NewLogoutOK().WithPayload(&body)
}
