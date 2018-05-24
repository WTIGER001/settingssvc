package actions

import (
	"fmt"
	"github.com/wtiger001/settingssvc/restapi/operations/configuration"

	"github.com/go-openapi/runtime/middleware"
)

// GetCategories reads all owner types
func GetCategories(params configuration.GetCategoriesParams) middleware.Responder {
	result := configuration.NewGetCategoriesOK()

	fmt.Printf("Starting --> GetCategories\n")

	arr, err := store.ReadAllCategories()

	if err != nil {
		fmt.Printf("\tError: %s\n", err.Error())
		return configuration.NewGetOwnerTypesInternalServerError()
	}

	result.SetPayload(arr)
	return result
}

// AddCategory ...
func AddCategory(params configuration.AddCategoryParams) middleware.Responder {
	fmt.Printf("Starting --> AddCategory, ID:%s\n", params.Category.Name)

	item := params.Category

	// Save as String
	err := store.StoreCategory(item)
	if err != nil {
		fmt.Printf("\tError: %s\n", err.Error())
		return configuration.NewAddCategoryInternalServerError()
	}

	// Return success
	var r = configuration.NewAddCategoryOK()
	// r.SetPayload(item)
	return r
}

// DeleteCategory deletes a definition
func DeleteCategory(params configuration.DeleteCategoryParams) middleware.Responder {
	fmt.Printf("Starting --> DeleteCategory, ID:%s\n", params.Name)

	// Check for duplicate ID
	exist, err := store.CategoryExists(params.Name)
	if err != nil {
		fmt.Printf("\tError: %s\n", err.Error())
		return configuration.NewDeleteCategoryInternalServerError()
	}
	if !exist {
		fmt.Print("Not Found\n")
		return configuration.NewDeleteCategoryNotFound()
	}

	// Delete the file
	store.DeleteCategory(params.Name)
	if err != nil {
		fmt.Printf("\tError: %s\n", err.Error())
		return configuration.NewDeleteCategoryInternalServerError()
	}

	return configuration.NewDeleteCategoryOK()
}
