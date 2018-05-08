package settings

import (
	"fmt"
	"io/ioutil"
	"models"
	"os"
	"path/filepath"
	"restapi/operations/configuration"

	"github.com/go-openapi/runtime/middleware"
)

// GetType reads a definition identified by an id
func GetType(params configuration.GetTypeParams) middleware.Responder {
	fmt.Printf("Starting --> GetType, ID:%s\n", params.ID)

	id := params.ID
	fmt.Printf("getType id:%s\n", id)

	path := store.pathToDefinition(id)
	fmt.Printf("getDefinitions path:%s\n", path)

	if !store.exists(path) {
		return &configuration.GetDefinitionNotFound{}
	}

	item, err := store.readType(path)
	if err != nil {
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
	dir := filepath.Join(store.dir, "types")
	fmt.Printf("\tDir: %s\n", dir)

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Printf("\tInternal Error: %s\n", err.Error())
		return configuration.NewGetOwnerTypesInternalServerError()
	}

	fmt.Printf("\tFound Files: %d\n", len(files))
	arr := make([]*models.OwnerType, 0, len(files))
	for _, f := range files {
		// Read the item
		fmt.Printf("\tChecking FIle: %s\n", f)

		if !f.IsDir() {
			myPath := filepath.Join(store.dir, "types", f.Name())
			fmt.Printf("\\rReading: %s\n", myPath)

			item, err := store.readType(myPath)
			if err != nil {
				fmt.Printf("\tBad Type: %s\n", f.Name())
			} else {
				fmt.Printf("\tAdding Type: %s\n", item.ID)
				arr = append(arr, item)
			}
		}
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
	if store.exists(store.pathToDefinition(item.ID)) {
		fmt.Print("Duplicate Found\n")
		return configuration.NewAddOwnerTypeInternalServerError()
	}

	// Perform other validation
	// if (def.ID == "") {

	// }

	// Save as String
	err := store.writeType(item)
	if err != nil {
		return configuration.NewAddOwnerTypeInternalServerError()
	}

	// Return success
	var r = configuration.NewAddOwnerTypeOK()
	// r.SetPayload(item)
	return r
}

// UpdateType updates an existing definition...
func UpdateType(params configuration.UpdateOwnerTypeParams) middleware.Responder {
	fmt.Printf("Starting --> UpdateType, ID:%s\n", params.Body.ID)

	item := params.Body

	// Check for duplicate ID
	// if store.exists(store.pathToDefinition(def.ID)) && !overwrite {
	// 	fmt.Print("Duplicate Found\n")
	// 	return configuration.NewAddDefinitionMethodNotAllowed()
	// }

	// Perform other validation
	// if (def.ID == "") {

	// }

	// Save as String
	err := store.writeType(item)
	if err != nil {
		return configuration.NewAddOwnerTypeInternalServerError()
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
	path := store.pathToType(id)
	if !store.exists(path) {
		fmt.Print("Not Found\n")
		return configuration.NewDeleteTypeNotFound()
	}

	// Delete the file
	err := os.Remove(path)
	if err != nil {
		return configuration.NewDeleteTypeInternalServerError()
	}

	return configuration.NewDeleteTypeOK()
}
