package actions

import (
	"fmt"
	"github.com/wtiger001/settingssvc/restapi/operations/configuration"

	"github.com/go-openapi/runtime/middleware"
)

// HACK IN MODEL
//HACK
// JsonData map[string]*json.RawMessage `json:"-"`

// GetDefinition reads a definition identified by an id
func GetDefinition(params configuration.GetDefinitionParams) middleware.Responder {
	id := params.ID
	fmt.Printf("getDefinition id:%s\n", id)

	exists, err := store.DefinitionExists(id)
	if err != nil {
		fmt.Printf("\tError: %s\n", err.Error())
		return configuration.NewGetDefinitionInternalServerError()
	}
	if !exists {
		return configuration.NewGetDefinitionNotFound()
	}

	item, err := store.ReadDefinition(id)
	if err != nil {
		fmt.Printf("\tError: %s\n", err.Error())
		return configuration.NewGetDefinitionInternalServerError()
	}

	var r = configuration.NewGetDefinitionOK()
	r.SetPayload(item)
	return r
}

// GetDefinitions reads a definition identified by an id
func GetDefinitions(params configuration.GetDefinitionsParams) middleware.Responder {
	fmt.Printf("getDefinitions\n")

	arr, err := store.ReadAllDefinitions()
	if err != nil {
		fmt.Printf("\tError: %s\n", err.Error())
		return configuration.NewGetDefinitionsInternalServerError()
	}

	result := configuration.NewGetDefinitionsOK()
	result.SetPayload(arr)
	return result
}

// AddDefinition ...
func AddDefinition(params configuration.AddDefinitionParams) middleware.Responder {
	fmt.Print("Adding definition\n")

	item := params.Body

	exists, err := store.DefinitionExists(item.ID)
	if err != nil {
		fmt.Printf("\tError: %s\n", err.Error())
		return configuration.NewAddDefinitionInternalServerError()
	}
	if exists {
		return configuration.NewAddDefinitionInternalServerError()
	}

	// Save as String
	err = store.StoreDefintion(item)
	if err != nil {
		fmt.Printf("\tError: %s\n", err.Error())
		return configuration.NewAddDefinitionInternalServerError()
	}

	// Return success
	var r = configuration.NewAddDefinitionOK()
	r.SetPayload(item)
	return r
}

// UpdateDefinition updates an existing definition...
func UpdateDefinition(params configuration.UpdateDefinitionParams) middleware.Responder {
	fmt.Print("Updating definition\n")

	item := params.Body

	// Save as String
	err := store.StoreDefintion(item)
	if err != nil {
		fmt.Printf("\tError: %s\n", err.Error())
		return configuration.NewAddDefinitionInternalServerError()
	}

	// Return success
	var r = configuration.NewAddDefinitionOK()
	r.SetPayload(item)
	return r
}

// DeleteDefinition deletes a definition
func DeleteDefinition(params configuration.DeleteDefinitionParams) middleware.Responder {
	id := params.ID

	err := store.DeleteDefintion(id)
	if err != nil {
		fmt.Printf("\tError: %s\n", err.Error())
		return configuration.NewDeleteDefinitionInternalServerError()
	}

	return configuration.NewDeleteDefinitionOK()
}
