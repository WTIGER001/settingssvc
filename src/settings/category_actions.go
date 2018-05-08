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

// GetCategories reads all owner types
func GetCategories(params configuration.GetCategoriesParams) middleware.Responder {
	result := configuration.NewGetCategoriesOK()

	fmt.Printf("Starting --> GetCategories\n")
	dir := filepath.Join(store.dir, "categories")
	fmt.Printf("\tDir: %s\n", dir)

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Printf("\tInternal Error: %s\n", err.Error())
		return configuration.NewGetOwnerTypesInternalServerError()
	}

	fmt.Printf("\tFound Files: %d\n", len(files))
	arr := make([]*models.Category, 0, len(files))
	for _, f := range files {
		// Read the item
		fmt.Printf("\tChecking FIle: %s\n", f)

		if !f.IsDir() {
			myPath := filepath.Join(store.dir, "categories", f.Name())
			fmt.Printf("\\rReading: %s\n", myPath)

			item, err := store.readCategory(myPath)
			if err != nil {
				fmt.Printf("\tBad Type: %s\n", f.Name())
			} else {
				fmt.Printf("\tAdding Type: %s\n", item.Name)
				arr = append(arr, item)
			}
		}
	}

	fmt.Printf("\tComplet with %d categories found\n", len(arr))

	result.SetPayload(arr)
	return result
}

// AddCategory ...
func AddCategory(params configuration.AddCategoryParams) middleware.Responder {
	fmt.Printf("Starting --> AddType, ID:%s\n", params.Category.Name)

	item := params.Category

	// Perform other validation
	// if (def.ID == "") {

	// }

	// Save as String
	err := store.writeCategory(item)
	if err != nil {
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

	id := params.Name

	// Check for duplicate ID
	path := store.pathToCategory(id)
	if !store.exists(path) {
		fmt.Print("Not Found\n")
		return configuration.NewDeleteCategoryNotFound()
	}

	// Delete the file
	err := os.Remove(path)
	if err != nil {
		return configuration.NewDeleteCategoryInternalServerError()
	}

	return configuration.NewDeleteCategoryOK()
}
