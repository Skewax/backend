package main

import (
	"log"
	"os"
	"strconv"
	"time"

	skewauth "github.com/Skewax/backend/internal/auth"
	"github.com/Skewax/backend/internal/handlers"
	"github.com/Skewax/backend/pkg/swagger/server/restapi"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"

	"github.com/Skewax/backend/pkg/swagger/server/restapi/operations"
	"github.com/Skewax/backend/pkg/swagger/server/restapi/operations/authentication"
	"github.com/Skewax/backend/pkg/swagger/server/restapi/operations/files"

	"github.com/joho/godotenv"
)

type CachedSessionData struct {
	googleToken string //access token not refresh token
	uid         string
	cacheExp    time.Time
	accessExp   time.Time
	sessionExp  time.Time
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalln("No .env file found")
		panic("no .env file")
	}
}

func main() {

	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewSkewaxBackendAPI(swaggerSpec)
	server := restapi.NewServer(api)

	defer func() {
		if err := server.Shutdown(); err != nil {
			log.Fatalln(err)
		}
	}()

	//ENVIRONMENT VARIABLE LOADING
	port, exists := os.LookupEnv("SKXPORT")
	intPort, validPort := strconv.Atoi(port)
	if !exists || validPort != nil {
		log.Println("No valid port found, defaulting to 8080")
		server.Port = 8080
	} else {
		server.Port = intPort
	}
	cacheTime, exists := os.LookupEnv("SKXCACHE")
	intCacheTime, err := strconv.Atoi(cacheTime)
	if !exists || err != nil {
		log.Println("No valid data cache time found, defaulting to 10 minutes")
		intCacheTime = 10
	}
	ConnStr, exists := os.LookupEnv("SKXCONN")
	if !exists {
		log.Fatalln("No database connection string")
		panic("no database connection string found")
	}
	GClientId, exists := os.LookupEnv("SKXGCID")
	if !exists {
		log.Fatalln("No Google Client ID")
		panic("missing google client ID")
	}
	GClientSecret, exists := os.LookupEnv("SKXGCSECRET")
	if !exists {
		log.Fatalln("No Google Client Secret")
		panic("missing google client secret")
	}

	//channels for interacting with authentication service
	GenChan := make(chan (*skewauth.GenTokenReq))
	TokenChan := make(chan (*skewauth.UseTokenReq))
	ClearTokenChan := make(chan (*skewauth.ClearTokenReq))
	QuitAuthService := make(chan bool)
	defer func() { QuitAuthService <- true }()

	authconf := skewauth.AuthConfig{
		CacheTime:     (time.Duration(intCacheTime) * time.Minute),
		ConnStr:       ConnStr,
		GClientId:     GClientId,
		GClientSecret: GClientSecret,
	}
	go skewauth.Begin(QuitAuthService, TokenChan, GenChan, ClearTokenChan, &authconf)

	//attaching handler functions to API

	api.AuthenticationGoogleLoginHandler = authentication.GoogleLoginHandlerFunc(
		func(params authentication.GoogleLoginParams) middleware.Responder {
			return handlers.HandleGoogleLogin(GenChan, params)
		})
	api.AuthenticationTokenLoginHandler = authentication.TokenLoginHandlerFunc(
		func(params authentication.TokenLoginParams) middleware.Responder {
			return handlers.HandleTokenLogin(TokenChan, params)
		})
	api.AuthenticationLogoutHandler = authentication.LogoutHandlerFunc(
		func(params authentication.LogoutParams) middleware.Responder {
			return handlers.HandleLogout(ClearTokenChan, params)
		})

	api.FilesGetFilesHandler = files.GetFilesHandlerFunc(handlers.HandleGetFiles)
	api.FilesReadFileHandler = files.ReadFileHandlerFunc(handlers.HandleReadFile)
	api.FilesUpdateFileHandler = files.UpdateFileHandlerFunc(handlers.HandleUpdateFile)
	api.FilesCreateFileHandler = files.CreateFileHandlerFunc(handlers.HandleCreateFile)
	api.FilesDeleteFileHandler = files.DeleteFileHandlerFunc(handlers.HandleDeleteFile)

	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}

}
