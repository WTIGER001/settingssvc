package restapi

import (
	"net/http"
	"github.com/wtiger001/settingssvc/restapi/operations"

	loads "github.com/go-openapi/loads"
)

func getAPI() (*operations.SettingsAPI, error) {
	swaggerSpec, err := loads.Analyzed(SwaggerJSON, "")
	if err != nil {
		return nil, err
	}
	api := operations.NewSettingsAPI(swaggerSpec)
	return api, nil
}

// GetAPIHandler ...
func GetAPIHandler() (http.Handler, error) {
	api, err := getAPI()
	if err != nil {
		return nil, err
	}
	h := configureAPI(api)
	err = api.Validate()
	if err != nil {
		return nil, err
	}
	return h, nil
}
