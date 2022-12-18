package main

import (
	"log"

	"github.com/Skewax/backend/pkg/swagger/server/restapi"
	"github.com/Skewax/backend/internal/handlers"
	"github.com/go-openapi/loads"
	// "github.com/go-openapi/runtime/middleware"

	"github.com/Skewax/backend/pkg/swagger/server/restapi/operations"
	"github.com/Skewax/backend/pkg/swagger/server/restapi/operations/authentication"
	"github.com/Skewax/backend/pkg/swagger/server/restapi/operations/files"
	// "golang.org/x/oauth2"
	// "golang.org/x/oauth2/google"
	// "google.golang.org/api/drive/v3"
)

func main() {

  swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
  if(err != nil) {
    log.Fatalln(err)
  }

  api := operations.NewSkewaxBackendAPI(swaggerSpec)
	server := restapi.NewServer(api)

  defer func() {
  	if err := server.Shutdown(); err != nil {
  		log.Fatalln(err)
  	}
  }()

	server.Port = 8080

	//attaching handler functions to API
	api.AuthenticationGoogleLoginHandler = authentication.GoogleLoginHandlerFunc(handlers.HandleGoogleLogin)
	api.AuthenticationTokenLoginHandler = authentication.TokenLoginHandlerFunc(handlers.HandleTokenLogin)
	api.AuthenticationLogoutHandler = authentication.LogoutHandlerFunc(handlers.HandleLogout)

	api.FilesGetFilesHandler = files.GetGetFilesHandlerFunc(handlers.HandleGetFiles)
	api.FilesReadFileHandler = files.GetReadFileHandlerFunc(handlers.HandleReadFile)
	api.FilesUpdateFileHandler = files.PostUpdateFileHandlerFunc(handlers.HandleUpdateFile)
	api.FilesCreateFileHandler = files.PostCreateFileHandlerFunc(handlers.HandleCreateFile)
	api.FilesDeleteFileHandler = files.PostDeleteFileHandlerFunc(handlers.HandleDeleteFile)

	if err := server.Serve();  err != nil {
		log.Fatalln(err)
	}
}




