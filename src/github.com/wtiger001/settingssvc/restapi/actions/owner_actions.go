package actions

import (
	"fmt"
	// "restapi/operations/configuration"
	"github.com/wtiger001/settingssvc/restapi/operations/preferences"

	"github.com/go-openapi/runtime/middleware"
)

// GetOwner reads a definition identified by an id
func GetOwner(params preferences.GetOwnerParams) middleware.Responder {
	id := params.ID
	fmt.Printf("getOwner id:%s\n", id)

	exists, err := store.OwnerExists(id)
	if err != nil {
		fmt.Printf("\tError: %s\n", err.Error())
		return preferences.NewGetOwnerInternalServerError()
	}
	if !exists {
		return preferences.NewGetOwnerNotFound()
	}

	item, err := store.ReadOwner(id)
	if err != nil {
		fmt.Printf("\tError: %s\n", err.Error())
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

	err := store.StoreOwner(item)
	if err != nil {
		fmt.Printf("\tError: %s\n", err.Error())
		return preferences.NewUpdateOwnerInternalServerError()
	}

	// Return success
	var r = preferences.NewUpdateOwnerOK()
	return r
}

// DeleteOwner deletes a definition
func DeleteOwner(params preferences.DeleteOwnerParams) middleware.Responder {
	id := params.ID

	exists, err := store.OwnerExists(id)
	if err != nil {
		fmt.Printf("\tError: %s\n", err.Error())
		return preferences.NewGetOwnerInternalServerError()
	}
	if !exists {
		return preferences.NewGetOwnerNotFound()
	}

	err = store.DeleteOwner(id)
	if err != nil {
		fmt.Printf("\tError: %s\n", err.Error())
		return preferences.NewGetOwnerInternalServerError()
	}

	return preferences.NewDeleteOwnerOK()
}
