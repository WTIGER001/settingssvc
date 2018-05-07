package settings

import (
	"fmt"
	"os"
	// "restapi/operations/configuration"
	"restapi/operations/preferences"

	"github.com/go-openapi/runtime/middleware"
)

// GetOwner reads a definition identified by an id
func GetOwner(params preferences.GetOwnerParams) middleware.Responder {
	id := params.ID
	fmt.Printf("getOwner id:%s\n", id)

	path := store.pathToOwner(id)
	fmt.Printf("getOwner path:%s\n", path)

	if !store.exists(path) {
		return preferences.NewGetOwnerNotFound()
	}

	item, err := store.readOwner(path)
	if err != nil {
		return preferences.NewGetOwnerInternalServerError()
	}

	var r = preferences.NewGetOwnerOK()
	r.SetPayload(item)
	return r
}

//UpdateOwner adds or updates an owner
func UpdateOwner(params preferences.UpdateOwnerParams) middleware.Responder {
	fmt.Print("Updating Owner\n")

	item := params.Body

	err := store.writeOwner(item)
	if err != nil {
		return preferences.NewUpdateOwnerInternalServerError()
	}

	// Return success
	var r = preferences.NewUpdateOwnerOK()
	return r
}

// DeleteOwner deletes a definition
func DeleteOwner(params preferences.DeleteOwnerParams) middleware.Responder {
	id := params.ID

	// Check for duplicate ID
	path := store.pathToOwner(id)
	if !store.exists(path) {
		fmt.Print("Not Found\n")
		return preferences.NewDeleteOwnerNotFound()
	}

	// Delete the file
	err := os.Remove(path)
	if err != nil {
		return preferences.NewDeleteOwnerInternalServerError()
	}

	return preferences.NewDeleteOwnerOK()
}
