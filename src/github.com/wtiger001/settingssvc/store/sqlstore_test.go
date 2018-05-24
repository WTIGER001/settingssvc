package store

import (
	"fmt"
	"github.com/wtiger001/settingssvc/models"
	"os"
	// "restapi/actions"
	// stor "store"
	util "github.com/wtiger001/settingssvc/test"
	"testing"
)

// var ts *Server

// TestMain global testing
var store ItemStore
var fstore FileStore

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	shutdown()
	os.Exit(code)
}

// setup uses a test database user
func setup() {
	fmt.Println("SETUP")
	os.Setenv("SETTINGS_SCHEMA", "testsettings")
	os.Setenv("SETTINGS_STORE", "sql")
	// actions.SetupStore()
	store = NewSQLStore()
	if err := store.InitStore(); err != nil {
		fmt.Printf("Error in Setup %s\n", err.Error())
		panic(err)
	}

	testdir := "c:/dev/projects/settingssvc/data"
	fstore = *NewFileStore(testdir)

	// Delete everything
	// store = actions.GetStore()
	if err := store.DeleteAll(); err != nil {
		fmt.Printf("Error in Setup %s\n", err.Error())
		panic(err)
	}

	// Create Sample
	// generateSampleTypes()
}

func shutdown() {
	store.CloseStore()
	// ts.Close()
	fmt.Println("SHUTDOWN")
}
func TestTypes(t *testing.T) {
	items, err := store.ReadAllTypes()
	util.AssertNilE(t, err)
	if len(items) != 0 {
		t.Fatalf("Too many items %d\n", len(items))
	}

	// Now load a type
	sample := &models.OwnerType{
		ID:          "sample",
		Name:        "SAMPLE NAME",
		Description: "Sample desc",
	}
	err = store.StoreType(sample)
	if len(items) > 0 {
		t.Fatalf("Too many items %d\n", len(items))
	}

	// Count again
	items, err = store.ReadAllTypes()
	util.AssertNilE(t, err)

	if len(items) != 1 {
		t.Fatalf("Too many items %d\n", len(items))
	}
	util.AssertEqual(t, sample, items[0], "Type Equal")

	// Delete
	store.DeleteType(sample.ID)
	util.AssertNilE(t, err)

	items, err = store.ReadAllTypes()
	util.AssertNilE(t, err)

	// Items have to be empty
	if len(items) > 0 {
		t.Fatalf("Too many items %d\n", len(items))
	}
}

func TestCategories(t *testing.T) {
	items, err := store.ReadAllCategories()
	util.AssertNilE(t, err)

	// Items have to be empty
	if len(items) > 0 {
		t.Fatalf("Too many items %d\n", len(items))
	}

	// Now load a type
	sample := &models.Category{
		ID:    "sample",
		Name:  "SAMPLE NAME",
		Order: 999,
	}
	err = store.StoreCategory(sample)
	if len(items) > 0 {
		t.Fatalf("Too many items %d\n", len(items))
	}

	// Count again
	items, err = store.ReadAllCategories()
	util.AssertNilE(t, err)

	if len(items) != 1 {
		t.Fatalf("Too many items %d\n", len(items))
	}
	util.AssertEqual(t, sample, items[0], "Type Equal")

	// Delete
	store.DeleteCategory(sample.ID)
	if err != nil {
		t.Fatalf("Error getting types: %s\n", err.Error())
	}

	items, err = store.ReadAllCategories()
	util.AssertNilE(t, err)

	// Items have to be empty
	if len(items) > 0 {
		t.Fatalf("Too many items %d\n", len(items))
	}
}

func TestDefinitions(t *testing.T) {
	items, err := store.ReadAllDefinitions()
	util.AssertNilE(t, err)

	// Items have to be empty
	if len(items) > 0 {
		t.Fatalf("Too many items %d\n", len(items))
	}

	// Now load a type
	sample := &models.PreferenceDefinition{
		ID:    "sample",
		Name:  "SAMPLE NAME",
		Order: 999,
	}
	err = store.StoreDefintion(sample)
	if len(items) > 0 {
		t.Fatalf("Too many items %d\n", len(items))
	}

	// Count again
	items, err = store.ReadAllDefinitions()
	util.AssertNilE(t, err)

	if len(items) != 1 {
		t.Fatalf("Too many items %d\n", len(items))
	}
	util.AssertEqual(t, sample, items[0], "Type Equal")

	// Delete
	store.DeleteDefintion(sample.ID)
	if err != nil {
		t.Fatalf("Error getting types: %s\n", err.Error())
	}

	items, err = store.ReadAllDefinitions()
	util.AssertNilE(t, err)

	// Items have to be empty
	if len(items) > 0 {
		t.Fatalf("Too many items %d\n", len(items))
	}
}
