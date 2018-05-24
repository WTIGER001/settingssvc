package actions

import (
	"fmt"
	"github.com/wtiger001/settingssvc/restapi/operations/configuration"

	"github.com/go-openapi/runtime/middleware"
)

// GetType reads a definition identified by an id
func GetType(params configuration.GetTypeParams) middleware.Responder {
	fmt.Printf("Starting --> GetType, ID:%s\n", params.ID)

	id := params.ID
	fmt.Printf("getType id:%s\n", id)

	exists, err := store.TypeExists(id)
	if err != nil {
		fmt.Printf("\tError: %s\n", err.Error())
		return configuration.NewGetTypeInternalServerError()
	}
	if !exists {
		return configuration.NewGetTypeNotFound()
	}

	item, err := store.ReadType(id)
	if err != nil {
		fmt.Printf("\tError: %s\n", err.Error())

		return &configuration.AddDefinitionMethodNotAllowed{}
	}

	var r = configuration.NewGetTypeOK()
	r.SetPayload(item)
	return r
}

// GetTypes reads all owner types
func GetTypes(params configuration.GetOwnerTypesParams) middleware.Responder {
	result := configuration.NewGetOwnerTypesOK()

	fmt.Printf("Starting --> GetTypes\n")

	arr, err := store.ReadAllTypes()
	if err != nil {
		fmt.Printf("\tError: %s\n", err.Error())
		return configuration.NewGetOwnerTypesInternalServerError()
	}
	fmt.Printf("\tComplet with %d types found\n", len(arr))

	result.SetPayload(arr)
	return result
}

// AddType ...
func AddType(params configuration.AddOwnerTypeParams) middleware.Responder {
	fmt.Printf("Starting --> AddType, ID:%s\n", params.Body.ID)

	item := params.Body

	// Check for duplicate ID
	exists, err := store.TypeExists(item.ID)
	if err != nil {
		fmt.Printf("\tError: %s\n", err.Error())
		return configuration.NewAddOwnerTypeInternalServerError()
	}
	if exists {
		fmt.Printf("\tAlready exists: %s\n", item.ID)
		return configuration.NewAddOwnerTypeInternalServerError()
	}

	// Save as String
	err = store.StoreType(item)
	if err != nil {
		fmt.Printf("\tError: %s\n", err.Error())
		return configuration.NewAddOwnerTypeInternalServerError()
	}

	// Return success
	var r = configuration.NewAddOwnerTypeOK()
	return r
}

// UpdateType updates an existing definition...
func UpdateType(params configuration.UpdateOwnerTypeParams) middleware.Responder {
	fmt.Printf("Starting --> UpdateType, ID:%s\n", params.Body.ID)

	item := params.Body

	// Save as String
	err := store.StoreType(item)
	if err != nil {
		fmt.Printf("\tError: %s\n", err.Error())
		return configuration.NewUpdateOwnerTypeInternalServerError()
	}

	// Return success
	var r = configuration.NewUpdateOwnerTypeOK()
	r.SetPayload(item)
	return r
}

// DeleteType deletes a definition
func DeleteType(params configuration.DeleteTypeParams) middleware.Responder {
	fmt.Printf("Starting --> DeleteType, ID:%s\n", params.ID)

	id := params.ID

	// Check for duplicate ID
	exists, err := store.TypeExists(id)
	if err != nil {
		fmt.Printf("\tError: %s\n", err.Error())
		return configuration.NewDeleteTypeInternalServerError()
	}
	if !exists {
		return configuration.NewDeleteTypeNotFound()
	}

	// Delete the file
	err = store.DeleteType(id)
	if err != nil {
		fmt.Printf("\tError: %s\n", err.Error())
		return configuration.NewDeleteTypeInternalServerError()
	}

	return configuration.NewDeleteTypeOK()
}
