package main

import (
	"log"

	"github.com/Skewax/backend/pkg/swagger/server/restapi"
	"github.com/go-openapi/loads"
	// "github.com/go-openapi/runtime/middleware"

	"github.com/Skewax/backend/pkg/swagger/server/restapi/operations"
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
	server.Port = 8080

  defer func() {
  	if err := server.Shutdown(); err != nil {
  		log.Fatalln(err)
  	}
  }()


	if err := server.Serve();  err != nil {
		log.Fatalln(err)
	}
}



