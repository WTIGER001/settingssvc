package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/wtiger001/settingssvc/restapi/actions"

	errors "github.com/go-openapi/errors"

	graceful "github.com/tylerb/graceful"

	"github.com/wtiger001/settingssvc/restapi/operations"
	"github.com/wtiger001/settingssvc/restapi/operations/configuration"
	"github.com/wtiger001/settingssvc/restapi/operations/preferences"

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
	api.JSONConsumer = actions.JSONConsumer()
	api.JSONProducer = actions.JSONProducer()

	// Definitions
	api.ConfigurationAddDefinitionHandler = configuration.AddDefinitionHandlerFunc(actions.AddDefinition)
	api.ConfigurationGetDefinitionHandler = configuration.GetDefinitionHandlerFunc(actions.GetDefinition)
	api.ConfigurationDeleteDefinitionHandler = configuration.DeleteDefinitionHandlerFunc(actions.DeleteDefinition)
	api.ConfigurationUpdateDefinitionHandler = configuration.UpdateDefinitionHandlerFunc(actions.UpdateDefinition)
	api.ConfigurationGetDefinitionsHandler = configuration.GetDefinitionsHandlerFunc(actions.GetDefinitions)

	// Types
	api.ConfigurationGetOwnerTypesHandler = configuration.GetOwnerTypesHandlerFunc(actions.GetTypes)
	api.ConfigurationAddOwnerTypeHandler = configuration.AddOwnerTypeHandlerFunc(actions.AddType)
	api.ConfigurationDeleteTypeHandler = configuration.DeleteTypeHandlerFunc(actions.DeleteType)
	api.ConfigurationUpdateOwnerTypeHandler = configuration.UpdateOwnerTypeHandlerFunc(actions.UpdateType)
	api.ConfigurationGetTypeHandler = configuration.GetTypeHandlerFunc(actions.GetType)

	// Categories
	api.ConfigurationAddCategoryHandler = configuration.AddCategoryHandlerFunc(actions.AddCategory)
	api.ConfigurationDeleteCategoryHandler = configuration.DeleteCategoryHandlerFunc(actions.DeleteCategory)
	api.ConfigurationGetCategoriesHandler = configuration.GetCategoriesHandlerFunc(actions.GetCategories)

	// Configuration
	// api.ConfigurationGetConfigHandler = configuration.GetConfigHandlerFunc(settings.GetConfig)

	// Owner
	api.PreferencesDeleteOwnerHandler = preferences.DeleteOwnerHandlerFunc(actions.DeleteOwner)
	api.PreferencesUpdateOwnerHandler = preferences.UpdateOwnerHandlerFunc(actions.UpdateOwner)
	api.PreferencesGetOwnerHandler = preferences.GetOwnerHandlerFunc(actions.GetOwner)

	// Profile
	api.PreferencesDeleteProfileHandler = preferences.DeleteProfileHandlerFunc(actions.DeleteProfile)
	api.PreferencesGetMyActiveProfileHandler = preferences.GetMyActiveProfileHandlerFunc(actions.GetMyActiveProfile)
	api.PreferencesGetProfileHandler = preferences.GetProfileHandlerFunc(actions.GetProfile)
	api.PreferencesGetProfileVersionsHandler = preferences.GetProfileVersionsHandlerFunc(actions.GetProfileVersions)
	api.PreferencesGetProfilesHandler = preferences.GetProfilesHandlerFunc(actions.GetProfiles)
	api.PreferencesUpdateProfileHandler = preferences.UpdateProfileHandlerFunc(actions.UpdateProfile)

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
	actions.SetupStore()

	corsHandler := cors.New(cors.Options{
		Debug:          false,
		AllowedHeaders: []string{"*"},
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "DELETE", "PUT", "PATCH"},
		MaxAge:         1000,
	})
	return corsHandler.Handler(handler)
	// return handler
}
