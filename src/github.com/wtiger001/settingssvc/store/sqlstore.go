package store

import (
	"fmt"
	"github.com/wtiger001/settingssvc/models"
	"os"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"github.com/jinzhu/inflection"
)

// SQLStore ...
type SQLStore struct {
	db *pg.DB
}

// NewSQLStore Creates a new SQL Store
func NewSQLStore() *SQLStore {
	var store = &SQLStore{}
	return store
}

// InitStore ...
func (sql *SQLStore) InitStore() error {
	sql.db = pg.Connect(&pg.Options{
		User:     "settings",
		Password: "settings",
	})

	
	schemaName := os.Getenv("SETTINGS_SCHEMA")
	fmt.Printf("Schema Name %s\n", schemaName)
	if schemaName != "" {
		orm.SetTableNameInflector(func(s string) string {
			fmt.Printf("OVERRIDING SCHEMA NAMES\n")
			return schemaName + "." + inflection.Plural(s)
		})
	}

	err := sql.createSchema()
	return err
}

func (sql *SQLStore) createSchema() error {
	options := &orm.CreateTableOptions{IfNotExists: true}

	if err := sql.db.CreateTable((*models.OwnerType)(nil), options); err != nil {
		fmt.Println("Creating table for OwnerType")
		return err
	}
	if err := sql.db.CreateTable((*models.PreferenceDefinition)(nil), options); err != nil {
		fmt.Println("Creating table for PreferenceDefinition")

		return err
	}
	if err := sql.db.CreateTable((*models.PreferenceOwner)(nil), options); err != nil {
		fmt.Println("Creating table for PreferenceOwner")
		return err
	}
	if err := sql.db.CreateTable((*models.Profile)(nil), options); err != nil {
		fmt.Println("Creating table for Profile")
		return err
	}
	if err := sql.db.CreateTable((*models.Category)(nil), options); err != nil {
		fmt.Println("Creating table for Category")
		return err
	}
	return nil
}

// CloseStore ...
func (sql *SQLStore) CloseStore() error {
	return sql.db.Close()
}

// StoreDefintion ...
func (sql *SQLStore) StoreDefintion(def *models.PreferenceDefinition) error {
	fmt.Printf("StoreDefintion %s\n", def.ID)

	// layout, _ := def.JsonData["layout"]
	// schema, _ := def.JsonData["schema"]

	// // This guy is strange
	// item := &DefinitionAlternante{
	// 	ID: def.ID,
	// 	// OwnerType: def.,
	// 	Name:     def.Name,
	// 	Order:    def.Order,
	// 	Category: def.Category,
	// 	Layout:   layout,
	// 	Schema:   schema,
	// }
	_, err := sql.db.Model(def).
		OnConflict("(id) DO UPDATE").
		// Set("ownertype = EXCLUDED.ownertype").
		Set("category = EXCLUDED.category").
		// Set("layout = EXCLUDED.layout").
		// Set("schema = EXCLUDED.schema").
		Set("Json_Data = EXCLUDED.Json_Data").
		Insert()

	return err
}

// StoreCategory ...
func (sql *SQLStore) StoreCategory(cat *models.Category) error {
	_, err := sql.db.Model(cat).
		OnConflict("DO NOTHING").
		Insert()

	return err
}

// StoreType ...
func (sql *SQLStore) StoreType(typ *models.OwnerType) error {
	_, err := sql.db.Model(typ).
		OnConflict("(id) DO UPDATE").
		Set("name = EXCLUDED.name").
		Set("description = EXCLUDED.description").
		Insert()

	return err
}

// StoreOwner ...
func (sql *SQLStore) StoreOwner(own *models.PreferenceOwner) error {
	_, err := sql.db.Model(own).
		OnConflict("(id) DO UPDATE").
		Set("type = EXCLUDED.type").
		Set("active = EXCLUDED.active").
		Insert()

	return err
}

// StoreProfile ...
func (sql *SQLStore) StoreProfile(profile *models.Profile) error {
	_, err := sql.db.Model(profile).
		OnConflict("(id) DO UPDATE").
		Set("name = EXCLUDED.name").
		Set("owner = EXCLUDED.owner").
		Insert()

	return err
}

// TypeExists ...
func (sql *SQLStore) TypeExists(id string) (bool, error) {
	item := &models.OwnerType{ID: id}
	count, err := sql.db.Model(item).Count()
	return count > 0, err
}

// DefinitionExists ...
func (sql *SQLStore) DefinitionExists(id string) (bool, error) {
	item := &models.PreferenceDefinition{ID: id}
	count, err := sql.db.Model(item).Count()
	return count > 0, err
}

// CategoryExists ...
func (sql *SQLStore) CategoryExists(id string) (bool, error) {
	item := &models.Category{Name: id}
	count, err := sql.db.Model(item).Count()
	return count > 0, err
}

// OwnerExists ...
func (sql *SQLStore) OwnerExists(id string) (bool, error) {
	item := &models.PreferenceOwner{ID: id}
	count, err := sql.db.Model(item).Count()
	return count > 0, err
}

// ProfileExists ..
func (sql *SQLStore) ProfileExists(id string, version int) (bool, error) {
	item := &models.Profile{ID: id, Version: int64(version)}
	count, err := sql.db.Model(item).Count()
	return count > 0, err
}

// ProfileVersions ...
func (sql *SQLStore) ProfileVersions(id string) ([]int, error) {
	item := &models.Profile{ID: id}
	var items []models.Profile

	err := sql.db.Model(item).Select(&items)
	if err == nil {
		fmt.Printf("In Versions\n")

		versions := make([]int, len(items))
		for i, profile := range items {
			versions[i] = int(profile.Version)
		}
		return versions, nil
	}
	fmt.Printf("Error  %s\n", err.Error())
	return nil, err
}

// LatestVersion ..
func (sql *SQLStore) LatestVersion(id string) (int, error) {
	versions, err := sql.ProfileVersions(id)
	if err != nil && len(versions) > 0 {
		max := versions[0]
		for _, v := range versions {
			if max < v {
				max = v
			}
		}
		return int(max), nil
	}
	return 0, err
}

// DeleteDefintion ...
func (sql *SQLStore) DeleteDefintion(id string) error {
	item := &models.PreferenceDefinition{ID: id}
	return sql.db.Delete(item)
}

// DeleteCategory ...
func (sql *SQLStore) DeleteCategory(id string) error {
	fmt.Printf("Deleteing Category Name: %s\n", id)
	item := &models.Category{ID: id}
	return sql.db.Delete(item)
}

// DeleteType ...
func (sql *SQLStore) DeleteType(id string) error {
	item := &models.OwnerType{ID: id}
	return sql.db.Delete(item)
}

// DeleteOwner ...
func (sql *SQLStore) DeleteOwner(id string) error {
	item := &models.PreferenceOwner{ID: id}
	return sql.db.Delete(item)
}

// DeleteProfile ...
func (sql *SQLStore) DeleteProfile(id string) error {
	item := &models.Profile{ID: id}
	return sql.db.Delete(item)
}

// ReadAllDefinitions ...
func (sql *SQLStore) ReadAllDefinitions() ([]*models.PreferenceDefinition, error) {
	fmt.Println("Selecting all definiions")
	var items []*models.PreferenceDefinition
	err := sql.db.Model(&items).Select()
	if err != nil {
		fmt.Printf("Error Drring Select %s\n", err.Error())
	}

	return items, err
}

// ReadAllCategories ...
func (sql *SQLStore) ReadAllCategories() ([]*models.Category, error) {
	var items []*models.Category
	err := sql.db.Model(&items).Select()
	return items, err
}

// ReadAllTypes ...
func (sql *SQLStore) ReadAllTypes() ([]*models.OwnerType, error) {
	var items []*models.OwnerType
	err := sql.db.Model(&items).Select()
	return items, err
}

// ReadDefinition ...
func (sql *SQLStore) ReadDefinition(id string) (*models.PreferenceDefinition, error) {
	item := &models.PreferenceDefinition{ID: id}
	err := sql.db.Select(item)
	return item, err
}

// ReadCategory ...
func (sql *SQLStore) ReadCategory(id string) (*models.Category, error) {
	item := &models.Category{Name: id}
	err := sql.db.Select(item)
	return item, err
}

// ReadType ...
func (sql *SQLStore) ReadType(id string) (*models.OwnerType, error) {
	item := &models.OwnerType{ID: id}
	err := sql.db.Select(item)
	return item, err
}

// ReadOwner ...
func (sql *SQLStore) ReadOwner(id string) (*models.PreferenceOwner, error) {
	item := &models.PreferenceOwner{ID: id}
	err := sql.db.Select(item)
	return item, err
}

// ReadProfile ...
func (sql *SQLStore) ReadProfile(id string, version int) (*models.Profile, error) {
	item := &models.Profile{ID: id}
	err := sql.db.Select(item)
	return item, err
}

// ReadLatestProfile ...
func (sql *SQLStore) ReadLatestProfile(id string) (*models.Profile, error) {
	v, err := sql.LatestVersion(id)
	if err != nil {
		return nil, err
	}
	return sql.ReadProfile(id, int(v))
}

// ReadLatestProfilesForOwner ..
func (sql *SQLStore) ReadLatestProfilesForOwner(owner string) ([]*models.Profile, error) {
	// get all the profiles for the owner... but we only want the latest
	var profiles []*models.Profile
	err := sql.db.Model(&profiles).
		Where("owner = ?", owner).
		Group("id").
		Having("MAX(version) = version").
		Select()

	if err != nil {
		return nil, err
	}
	return profiles, err
}

// DeleteAll removes all items
func (sql *SQLStore) DeleteAll() error {
	var categories *models.Category
	var profiles *models.Profile
	var owners *models.PreferenceOwner
	var types *models.OwnerType
	var definitions *models.PreferenceDefinition

	if _, err := sql.db.Model(categories).Where("1=1").Delete(); err != nil {
		return err
	}

	if _, err := sql.db.Model(profiles).Where("1=1").Delete(); err != nil {
		return err
	}
	if _, err := sql.db.Model(owners).Where("1=1").Delete(); err != nil {
		return err
	}
	if _, err := sql.db.Model(types).Where("1=1").Delete(); err != nil {
		return err
	}
	if _, err := sql.db.Model(definitions).Where("1=1").Delete(); err != nil {
		return err
	}

	return nil
}
