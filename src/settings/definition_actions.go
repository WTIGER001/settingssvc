package settings

import (
	"fmt"
	"io/ioutil"
	"models"
	"os"
	"restapi/operations/configuration"

	"path/filepath"

	"github.com/go-openapi/runtime/middleware"
)

// HACK IN MODEL
//HACK
// JsonData map[string]*json.RawMessage `json:"-"`

// GetDefinition reads a definition identified by an id
func GetDefinition(params configuration.GetDefinitionParams) middleware.Responder {
	id := params.ID
	fmt.Printf("getDefinitions id:%s\n", id)

	path := store.pathToDefinition(id)
	fmt.Printf("getDefinitions path:%s\n", path)

	if !store.exists(path) {
		return &configuration.GetDefinitionNotFound{}
	}

	item, err := store.readDefinition(path)
	if err != nil {
		return &configuration.AddDefinitionMethodNotAllowed{}
	}

	var r = &configuration.GetDefinitionOK{}
	r.SetPayload(item)
	return r
}

// GetDefinitions reads a definition identified by an id
func GetDefinitions(params configuration.GetDefinitionsParams) middleware.Responder {
	result := configuration.NewGetDefinitionsOK()

	dir := filepath.Join(store.dir, "definitions")

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return configuration.NewGetDefinitionsInternalServerError()
	}

	arr := make([]*models.PreferenceDefinition, 0, len(files))
	for _, f := range files {
		// Read the item
		if !f.IsDir() {
			mypath := filepath.Join(store.dir, "definitions", f.Name())
			item, err := store.readDefinition(mypath)
			if err != nil {
				fmt.Print("BAD DATA %s", f.Name)
			} else {
				arr = append(arr, item)
			}
		}
	}

	result.SetPayload(arr)
	return result
}

// AddDefinition ...
func AddDefinition(params configuration.AddDefinitionParams) middleware.Responder {
	fmt.Print("Adding definition\n")

	def := params.Body

	// Check for duplicate ID
	if store.exists(store.pathToDefinition(def.ID)) {
		fmt.Print("Duplicate Found\n")
		return configuration.NewAddDefinitionMethodNotAllowed()
	}

	// Perform other validation
	// if (def.ID == "") {

	// }

	// Save as String
	err := store.writeDefinition(def)
	if err != nil {
		return &configuration.AddDefinitionMethodNotAllowed{}
	}

	// Return success
	var r = configuration.NewAddDefinitionOK()
	r.SetPayload(def)
	return r
}

// UpdateDefinition updates an existing definition...
func UpdateDefinition(params configuration.UpdateDefinitionParams) middleware.Responder {
	fmt.Print("Updating definition\n")

	def := params.Body

	// Check for duplicate ID
	// if store.exists(store.pathToDefinition(def.ID)) && !overwrite {
	// 	fmt.Print("Duplicate Found\n")
	// 	return configuration.NewAddDefinitionMethodNotAllowed()
	// }

	// Perform other validation
	// if (def.ID == "") {

	// }

	// Save as String
	err := store.writeDefinition(def)
	if err != nil {
		return configuration.NewUpdateDefinitionInternalServerError()
	}

	// Return success
	var r = configuration.NewUpdateDefinitionOK()
	r.SetPayload(def)
	return r
}

// DeleteDefinition deletes a definition
func DeleteDefinition(params configuration.DeleteDefinitionParams) middleware.Responder {
	id := params.ID

	// Check for duplicate ID
	path := store.pathToDefinition(id)
	if !store.exists(path) {
		fmt.Print("Not Found\n")
		return configuration.NewDeleteDefinitionNotFound()
	}

	// Delete the file
	err := os.Remove(path)
	if err != nil {
		return configuration.NewDeleteDefinitionInternalServerError()
	}

	return configuration.NewDeleteDefinitionOK()
}
