package restapi

import (
	"crypto/tls"
	"net/http"

	"settings"

	errors "github.com/go-openapi/errors"

	graceful "github.com/tylerb/graceful"

	"restapi/operations"
	"restapi/operations/configuration"
	"restapi/operations/preferences"

	"github.com/rs/cors"
)

// This file is safe to edit. Once it exists it will not be overwritten

//go:generate swagger generate server --target .. --name settings --spec ..\..\..\settings\swagger\swagger.yaml

func configureFlags(api *operations.SettingsAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.SettingsAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	// api.JSONConsumer = runtime.JSONConsumer()
	// api.JSONProducer = runtime.JSONProducer()
	api.JSONConsumer = settings.JSONConsumer()
	api.JSONProducer = settings.JSONProducer()

	// Definitions
	api.ConfigurationAddDefinitionHandler = configuration.AddDefinitionHandlerFunc(settings.AddDefinition)
	api.ConfigurationGetDefinitionHandler = configuration.GetDefinitionHandlerFunc(settings.GetDefinition)
	api.ConfigurationDeleteDefinitionHandler = configuration.DeleteDefinitionHandlerFunc(settings.DeleteDefinition)
	api.ConfigurationUpdateDefinitionHandler = configuration.UpdateDefinitionHandlerFunc(settings.UpdateDefinition)
	api.ConfigurationGetDefinitionsHandler = configuration.GetDefinitionsHandlerFunc(settings.GetDefinitions)

	// Types
	api.ConfigurationGetOwnerTypesHandler = configuration.GetOwnerTypesHandlerFunc(settings.GetTypes)
	api.ConfigurationAddOwnerTypeHandler = configuration.AddOwnerTypeHandlerFunc(settings.AddType)
	api.ConfigurationDeleteTypeHandler = configuration.DeleteTypeHandlerFunc(settings.DeleteType)
	api.ConfigurationUpdateOwnerTypeHandler = configuration.UpdateOwnerTypeHandlerFunc(settings.UpdateType)
	api.ConfigurationGetTypeHandler = configuration.GetTypeHandlerFunc(settings.GetType)

	// Categories
	api.ConfigurationAddCategoryHandler = configuration.AddCategoryHandlerFunc(settings.AddCategory)
	api.ConfigurationDeleteCategoryHandler = configuration.DeleteCategoryHandlerFunc(settings.DeleteCategory)
	api.ConfigurationGetCategoriesHandler = configuration.GetCategoriesHandlerFunc(settings.GetCategories)

	// Configuration
	// api.ConfigurationGetConfigHandler = configuration.GetConfigHandlerFunc(settings.GetConfig)

	// Owner
	api.PreferencesDeleteOwnerHandler = preferences.DeleteOwnerHandlerFunc(settings.DeleteOwner)
	api.PreferencesUpdateOwnerHandler = preferences.UpdateOwnerHandlerFunc(settings.UpdateOwner)
	api.PreferencesGetOwnerHandler = preferences.GetOwnerHandlerFunc(settings.GetOwner)

	// Profile
	api.PreferencesDeleteProfileHandler = preferences.DeleteProfileHandlerFunc(settings.DeleteProfile)
	api.PreferencesGetMyActiveProfileHandler = preferences.GetMyActiveProfileHandlerFunc(settings.GetMyActiveProfile)
	api.PreferencesGetProfileHandler = preferences.GetProfileHandlerFunc(settings.GetProfile)
	api.PreferencesGetProfileVersionsHandler = preferences.GetProfileVersionsHandlerFunc(settings.GetProfileVersions)
	api.PreferencesGetProfilesHandler = preferences.GetProfilesHandlerFunc(settings.GetProfiles)
	api.PreferencesUpdateProfileHandler = preferences.UpdateProfileHandlerFunc(settings.UpdateProfile)

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
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *graceful.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	corsHandler := cors.New(cors.Options{
		Debug:          false,
		AllowedHeaders: []string{"*"},
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{},
		MaxAge:         1000,
	})
	return corsHandler.Handler(handler)
	// return handler
}
