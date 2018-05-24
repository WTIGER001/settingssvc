package store

import (
	"github.com/wtiger001/settingssvc/models"
)

//ItemStore is the basic interface for the ability to store and retrieve the
type ItemStore interface {
	// STORE
	StoreDefintion(def *models.PreferenceDefinition) error
	StoreCategory(cat *models.Category) error
	StoreType(typ *models.OwnerType) error
	StoreOwner(typ *models.PreferenceOwner) error
	StoreProfile(typ *models.Profile) error

	//GET
	ReadDefinition(id string) (*models.PreferenceDefinition, error)
	ReadCategory(id string) (*models.Category, error)
	ReadType(id string) (*models.OwnerType, error)
	ReadOwner(id string) (*models.PreferenceOwner, error)
	ReadProfile(id string, version int) (*models.Profile, error)
	ReadLatestProfile(id string) (*models.Profile, error)
	ReadLatestProfilesForOwner(owner string) ([]*models.Profile, error)

	ReadAllDefinitions() ([]*models.PreferenceDefinition, error)
	ReadAllCategories() ([]*models.Category, error)
	ReadAllTypes() ([]*models.OwnerType, error)

	// Exists
	TypeExists(id string) (bool, error)
	DefinitionExists(id string) (bool, error)
	CategoryExists(id string) (bool, error)
	OwnerExists(id string) (bool, error)
	ProfileExists(id string, version int) (bool, error)

	// DELETE
	DeleteDefintion(id string) error
	DeleteCategory(id string) error
	DeleteType(id string) error
	DeleteOwner(id string) error
	DeleteProfile(id string) error

	ProfileVersions(id string) ([]int, error)
	LatestVersion(id string) (int, error)

	DeleteAll() error
	InitStore() error
	CloseStore() error
}
