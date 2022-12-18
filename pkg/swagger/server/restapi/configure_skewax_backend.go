// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/Skewax/backend/pkg/swagger/server/restapi/operations"
	"github.com/Skewax/backend/pkg/swagger/server/restapi/operations/authentication"
	"github.com/Skewax/backend/pkg/swagger/server/restapi/operations/files"
)

//go:generate swagger generate server --target ../../server --name SkewaxBackend --spec ../../swagger.yml --principal interface{} --exclude-main

func configureFlags(api *operations.SkewaxBackendAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.SkewaxBackendAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	if api.FilesGetGetFilesHandler == nil {
		api.FilesGetGetFilesHandler = files.GetGetFilesHandlerFunc(func(params files.GetGetFilesParams) middleware.Responder {
			return middleware.NotImplemented("operation files.GetGetFiles has not yet been implemented")
		})
	}
	if api.FilesGetReadFileHandler == nil {
		api.FilesGetReadFileHandler = files.GetReadFileHandlerFunc(func(params files.GetReadFileParams) middleware.Responder {
			return middleware.NotImplemented("operation files.GetReadFile has not yet been implemented")
		})
	}
	if api.AuthenticationGetTokenLoginHandler == nil {
		api.AuthenticationGetTokenLoginHandler = authentication.GetTokenLoginHandlerFunc(func(params authentication.GetTokenLoginParams) middleware.Responder {
			return middleware.NotImplemented("operation authentication.GetTokenLogin has not yet been implemented")
		})
	}
	if api.FilesPostCreateFileHandler == nil {
		api.FilesPostCreateFileHandler = files.PostCreateFileHandlerFunc(func(params files.PostCreateFileParams) middleware.Responder {
			return middleware.NotImplemented("operation files.PostCreateFile has not yet been implemented")
		})
	}
	if api.FilesPostDeleteFileHandler == nil {
		api.FilesPostDeleteFileHandler = files.PostDeleteFileHandlerFunc(func(params files.PostDeleteFileParams) middleware.Responder {
			return middleware.NotImplemented("operation files.PostDeleteFile has not yet been implemented")
		})
	}
	if api.AuthenticationPostLogoutHandler == nil {
		api.AuthenticationPostLogoutHandler = authentication.PostLogoutHandlerFunc(func(params authentication.PostLogoutParams) middleware.Responder {
			return middleware.NotImplemented("operation authentication.PostLogout has not yet been implemented")
		})
	}
	if api.AuthenticationPostNewLoginHandler == nil {
		api.AuthenticationPostNewLoginHandler = authentication.PostNewLoginHandlerFunc(func(params authentication.PostNewLoginParams) middleware.Responder {
			return middleware.NotImplemented("operation authentication.PostNewLogin has not yet been implemented")
		})
	}
	if api.FilesPostUpdateFileHandler == nil {
		api.FilesPostUpdateFileHandler = files.PostUpdateFileHandlerFunc(func(params files.PostUpdateFileParams) middleware.Responder {
			return middleware.NotImplemented("operation files.PostUpdateFile has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
