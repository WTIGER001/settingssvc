package store

import (
	"github.com/wtiger001/settingssvc/models"

	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

// FileStore ...
type FileStore struct {
	dir string
}

// DeleteAll removes all items from the store... big reset
func (fs *FileStore) DeleteAll() error {
	d, err := os.Open(fs.dir)
	if err != nil {
		return err
	}
	defer d.Close()
	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(fs.dir, name))
		if err != nil {
			return err
		}
	}
	return nil
}

// DELETE

// NewFileStore creates a new filestore based on a root directory.
func NewFileStore(rootdir string) *FileStore {
	var store = &FileStore{
		dir: rootdir,
	}
	return store
}

// InitStore ...
func (fs *FileStore) InitStore() error {

	return nil
}

// CloseStore ...
func (fs *FileStore) CloseStore() error {
	return nil
}

// TypeExists ...
func (fs *FileStore) TypeExists(id string) (bool, error) {
	path := fs.pathToType(id)
	return fs.exists(path)
}

// DefinitionExists ...
func (fs *FileStore) DefinitionExists(id string) (bool, error) {
	path := fs.pathToDefinition(id)
	return fs.exists(path)
}

// CategoryExists ...
func (fs *FileStore) CategoryExists(id string) (bool, error) {
	path := fs.pathToCategory(id)
	return fs.exists(path)
}

// OwnerExists ...
func (fs *FileStore) OwnerExists(id string) (bool, error) {
	path := fs.pathToOwner(id)
	return fs.exists(path)
}

// ProfileExists ...
func (fs *FileStore) ProfileExists(id string, version int) (bool, error) {
	path := fs.pathToProfileVersion(id, version)
	return fs.exists(path)
}

// DeleteDefintion ...
func (fs *FileStore) DeleteDefintion(id string) error {
	path := fs.pathToDefinition(id)
	return os.Remove(path)
}

// StoreProfile ...
func (fs *FileStore) StoreProfile(item *models.Profile) error {
	fmt.Printf("\tWriting Profile \n")

	path := fs.pathToProfileVersionOwner(item.ID, int(item.Version), item.Owner)
	fmt.Printf("Writing profile to file %s\n", path)

	return fs.write(path, item)
}

// ReadProfile ...
func (fs *FileStore) ReadProfile(ID string, version int) (*models.Profile, error) {
	path := fs.pathToProfileVersion(ID, version)

	fmt.Printf("Reading File %s\n", path)
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("Error reading file: %s\n", err.Error())
		return nil, err
	}
	fmt.Printf("\tRead %d bytes\n", len(bytes))

	// Unmarshall "most" of this object.
	result := &models.Profile{}
	err = json.Unmarshal(bytes, result)
	if err != nil {
		return nil, err
	}
	fmt.Printf("\tUnmarshalled %s\n", result.ID)

	// now read the value
	err = ParseProfile(bytes, result)
	fmt.Printf("\tComplete %s\n", result.ID)

	return result, nil
}

// StoreOwner ...
func (fs *FileStore) StoreOwner(item *models.PreferenceOwner) error {

	path := fs.pathToOwner(item.ID)
	fmt.Printf("Writing Owner to file %s\n", path)

	return fs.write(path, item)
}

// ReadOwner ...
func (fs *FileStore) ReadOwner(ID string) (*models.PreferenceOwner, error) {
	path := fs.pathToOwner(ID)

	fmt.Printf("Reading File\n")
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("Errpr reading file: %s\n", err.Error())
		return nil, err
	}

	// Unmarshall "most" of this object.
	result := &models.PreferenceOwner{}
	err = json.Unmarshal(bytes, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// StoreCategory ...
func (fs *FileStore) StoreCategory(item *models.Category) error {

	path := fs.pathToCategory(item.Name)
	fmt.Printf("Writing Category to file %s\n", path)

	return fs.write(path, item)
}

// ReadCategory ...
func (fs *FileStore) ReadCategory(ID string) (*models.Category, error) {
	path := fs.pathToCategory(ID)
	return readCategoryPath(path)
}

func readCategoryPath(path string) (*models.Category, error) {
	fmt.Printf("Reading File\n")
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("Errpr reading file: %s\n", err.Error())
		return nil, err
	}

	// Unmarshall "most" of this object.
	result := &models.Category{}
	err = json.Unmarshal(bytes, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// StoreType ...
func (fs *FileStore) StoreType(item *models.OwnerType) error {

	path := fs.pathToType(item.ID)
	fmt.Printf("Writing Type to file %s\n", path)

	return fs.write(path, item)
}

// ReadType ...
func (fs *FileStore) ReadType(ID string) (*models.OwnerType, error) {
	path := fs.pathToType(ID)
	return readTypePath(path)
}

func readTypePath(path string) (*models.OwnerType, error) {
	fmt.Printf("Reading File %s\n", path)
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("Errpr reading file: %s\n", err.Error())
		return nil, err
	}

	// Unmarshall "most" of this object.
	result := &models.OwnerType{}
	err = json.Unmarshal(bytes, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// StoreDefintion ...
func (fs *FileStore) StoreDefintion(item *models.PreferenceDefinition) error {

	path := fs.pathToDefinition(item.ID)
	fmt.Printf("Writing Type to file %s\n", path)

	return fs.write(path, item.JsonData)
}

// ReadDefinition ...
func (fs *FileStore) ReadDefinition(ID string) (*models.PreferenceDefinition, error) {
	path := fs.pathToDefinition(ID)

	fmt.Printf("Reading File\n")
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("Errpr reading file: %s\n", err.Error())
		return nil, err
	}

	// Unmarshall "most" of this object.
	result := &models.PreferenceDefinition{}

	err = ParsePreferenceDefinition(bytes, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func readDefinitionPath(path string) (*models.PreferenceDefinition, error) {
	fmt.Printf("Reading File\n")
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("Errpr reading file: %s\n", err.Error())
		return nil, err
	}

	// Unmarshall "most" of this object.
	result := &models.PreferenceDefinition{}

	err = ParsePreferenceDefinition(bytes, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DeleteCategory ...
func (fs *FileStore) DeleteCategory(id string) error {
	path := fs.pathToCategory(id)
	return os.Remove(path)
}

// DeleteType ...
func (fs *FileStore) DeleteType(id string) error {
	path := fs.pathToType(id)
	return os.Remove(path)
}

// DeleteOwner ...
func (fs *FileStore) DeleteOwner(id string) error {
	path := fs.pathToOwner(id)
	return os.Remove(path)
}

// DeleteProfile ...
func (fs *FileStore) DeleteProfile(id string) error {
	files, err := fs.listProfiles(id)
	if err != nil {
		return err
	}
	for _, f := range files {
		os.Remove(f)
	}
	return nil
}

func (fs *FileStore) exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	return false, err
}

func (fs *FileStore) pathToDefinition(id string) string {
	return filepath.Join(fs.dir, "definitions", id+".definition.json")
}

func (fs *FileStore) pathToType(id string) string {
	return filepath.Join(fs.dir, "types", id+".type.json")
}

func (fs *FileStore) pathToCategory(id string) string {
	return filepath.Join(fs.dir, "categories", id+".category.json")
}

func (fs *FileStore) pathToOwner(id string) string {
	return filepath.Join(fs.dir, "owners", id+".owner.json")
}

func (fs *FileStore) pathToProfileVersion(id string, version int) string {
	file, _ := fs.listProfileVersion(id, version)
	return file
}
func (fs *FileStore) pathToProfileVersionOwner(id string, version int, owner string) string {
	return fmt.Sprintf("%s/profiles/%s.%s.v%d.profile.json", fs.dir, owner, id, version)
}

func (fs *FileStore) pathToProfile(id string) string {
	maxVersion, owner := fs.latestProfileVersion(id)
	return fs.pathToProfileVersionOwner(id, maxVersion, owner)
}

// ProfileVersions ...
func (fs *FileStore) ProfileVersions(id string) ([]int, error) {
	files, err := fs.listProfiles(id)
	if err != nil {
		return nil, err
	}
	versions := make([]int, len(files))
	for i, file := range files {
		_, versions[i], _, _ = extractIDVersionOwner(file)
	}
	return versions, nil
}

// LatestVersion ...
func (fs *FileStore) LatestVersion(id string) (int, error) {
	v, _ := fs.latestProfileVersion(id)
	return v, nil
}

// LatestProfileVersion finds the largest version for a single profile.
// A return of 0 means that no version was found
func (fs *FileStore) latestProfileVersion(id string) (int, string) {
	fmt.Printf("Listing Profiles\n")
	var owner string
	files, err := fs.listProfiles(id)
	if err != nil {
		return 0, ""
	}
	fmt.Printf("\tfound %d\n", len(files))

	var max int
	for _, f := range files {
		_, v, _, err := extractIDVersionOwner(f)
		fmt.Printf("\tfound version %d\n", v)

		if err != nil {
			fmt.Printf("error decoding version string\n")
		} else {
			if v > max {
				max = v
			}
		}
	}

	return max, owner
}

//ProfileVersionDate gets the date of a profile version
func (fs *FileStore) profileVersionDate(profileName string) (time.Time, error) {
	path := fs.dir + "/profile/" + profileName
	return fs.getFileDate(path)
}

func (fs *FileStore) getFileDate(path string) (time.Time, error) {
	stat, err := os.Stat(path)
	if err != nil {
		return time.Now(), err
	}
	return stat.ModTime(), nil
}

//ExtractVersionOwner gets the version from a filename or path
func extractIDVersionOwner(filename string) (string, int, string, error) {
	fmt.Printf("Extract Version %s\n", filename)

	parts := strings.Split(filename, ".")
	for i, part := range parts {
		fmt.Printf("\t Part: %d: %s\n", i, part)
	}
	versionstr := parts[len(parts)-3]
	fmt.Printf("\t%s\n", versionstr)

	versionstr = versionstr[1:]
	fmt.Printf("\t%s\n", versionstr)

	v, err := strconv.Atoi(versionstr)
	if err != nil {
		return "", 0, "", err
	}
	fmt.Printf("\t%d\n", v)

	return parts[1], v, parts[0], nil
}

//ListProfiles for an id
func (fs *FileStore) listProfiles(id string) ([]string, error) {
	part := fs.dir + "/profiles/*." + id + ".v*.profile.json"
	files, err := filepath.Glob(part)
	return files, err
}

// ListProfilesForOwner ...
func (fs *FileStore) listProfilesForOwner(owner string) ([]string, error) {
	part := fmt.Sprintf("%s/profiles/%s.*.*.profile.json", fs.dir, owner)

	files, err := filepath.Glob(part)
	return files, err
}

// ListProfileVersion ...
func (fs *FileStore) listProfileVersion(id string, version int) (string, error) {
	// part := fs.dir + "/profiles/*." + id + ".v" + .profile.json"
	part := fmt.Sprintf("%s/profiles/*.%s.v%d.profile.json", fs.dir, id, version)
	files, err := filepath.Glob(part)
	if len(files) == 1 {
		// return fmt.Sprintf("%s/profiles/%s", fs.dir, files[0]), err
		return files[0], err
	}
	return "", err
}

// func (fs *FileStore) ListProfiles(id string) ([]string, error) {
// 	part := fs.dir + "/profiles/" + id + "*.profile.json"
// 	files, err := filepath.Glob(part)
// 	return files, err
// }

func (fs *FileStore) makeDir(path string) error {
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		fmt.Printf("Could not create path : %s \n", dir)
		return err
	}
	return nil
}

func (fs *FileStore) write(path string, data interface{}) error {
	fmt.Printf("Writing : %s\n", path)

	if err := fs.makeDir(path); err != nil {
		fmt.Printf("Could not create path : %s\n", path)
		return err
	}

	bytes, _ := json.Marshal(data)
	err := ioutil.WriteFile(path, bytes, 0644)
	return err
}

// ParsePreferenceDefinition ...
func ParsePreferenceDefinition(bytes []byte, data *models.PreferenceDefinition) error {

	// Unmarshall "most" of this object.
	if err := json.Unmarshal(bytes, data); err != nil {
		return err
	}

	// Now unmarshal the schema and layout fields manually
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(bytes, &objMap)
	if err != nil {
		return err
	}
	data.JsonData = objMap

	// // Parse schema object
	// if schMsg, ok := objMap["schema"]; ok {
	// 	// schema := string(*schMsg)
	// 	data.Schema = schMsg
	// }

	// // Layout schema object
	// if layMsg, ok := objMap["layout"]; ok {
	// 	// layout := string(*layMsg)
	// 	data.Layout = layMsg
	// }

	fmt.Printf("ID: %s\n", data.ID)
	fmt.Printf("Cat: %s\n", data.Category)
	fmt.Printf("Name: %s\n", data.Name)
	// fmt.Printf("schema: %s\n", data.Schema)
	// fmt.Printf("layout: %s\n", data.Layout)
	return nil
}

// ParseProfile ...
func ParseProfile(bytes []byte, data *models.Profile) error {

	// Unmarshall "most" of this object.
	if err := json.Unmarshal(bytes, data); err != nil {
		return err
	}

	// Now unmarshal the schema and layout fields manually
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(bytes, &objMap)
	if err != nil {
		return err
	}
	data.JsonData = objMap

	return nil
}

// ReadAllDefinitions ...
func (fs *FileStore) ReadAllDefinitions() ([]*models.PreferenceDefinition, error) {
	fmt.Printf("Starting --> ReadAllDefinitions\n")

	dir := filepath.Join(fs.dir, "definitions")

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	fmt.Printf("\tFound Files: %d\n", len(files))
	arr := make([]*models.PreferenceDefinition, 0, len(files))
	for _, f := range files {
		// Read the item
		fmt.Printf("\tChecking FIle: %s\n", f)

		if !f.IsDir() {
			myPath := filepath.Join(fs.dir, "categories", f.Name())
			fmt.Printf("\\rReading: %s\n", myPath)

			item, err := readDefinitionPath(myPath)
			if err != nil {
				fmt.Printf("\tBad Type: %s\n", f.Name())
			} else {
				fmt.Printf("\tAdding Type: %s\n", item.Name)
				arr = append(arr, item)
			}
		}
	}

	fmt.Printf("\tComplet with %d categories found\n", len(arr))
	return arr, nil
}

// ReadAllCategories ...
func (fs *FileStore) ReadAllCategories() ([]*models.Category, error) {
	fmt.Printf("Starting --> ReadAllCategories\n")

	dir := filepath.Join(fs.dir, "categories")

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	fmt.Printf("\tFound Files: %d\n", len(files))
	arr := make([]*models.Category, 0, len(files))
	for _, f := range files {
		// Read the item
		fmt.Printf("\tChecking FIle: %s\n", f)

		if !f.IsDir() {
			myPath := filepath.Join(fs.dir, "categories", f.Name())
			fmt.Printf("\\rReading: %s\n", myPath)

			item, err := readCategoryPath(myPath)
			if err != nil {
				fmt.Printf("\tBad Type: %s\n", f.Name())
			} else {
				fmt.Printf("\tAdding Type: %s\n", item.Name)
				arr = append(arr, item)
			}
		}
	}

	fmt.Printf("\tComplet with %d categories found\n", len(arr))
	return arr, nil
}

// ReadAllTypes ...
func (fs *FileStore) ReadAllTypes() ([]*models.OwnerType, error) {
	fmt.Printf("Starting --> ReadAllCategories\n")

	dir := filepath.Join(fs.dir, "types")

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	fmt.Printf("\tFound Files: %d\n", len(files))
	arr := make([]*models.OwnerType, 0, len(files))
	for _, f := range files {
		// Read the item
		fmt.Printf("\tChecking FIle: %s\n", f)

		if !f.IsDir() {
			myPath := filepath.Join(fs.dir, "categories", f.Name())
			fmt.Printf("\\rReading: %s\n", myPath)

			item, err := readTypePath(myPath)
			if err != nil {
				fmt.Printf("\tBad Type: %s\n", f.Name())
			} else {
				fmt.Printf("\tAdding Type: %s\n", item.Name)
				arr = append(arr, item)
			}
		}
	}

	fmt.Printf("\tComplet with %d types found\n", len(arr))
	return arr, nil
}

// ReadLatestProfile ...
func (fs *FileStore) ReadLatestProfile(id string) (*models.Profile, error) {
	v, err := fs.LatestVersion(id)
	if err != nil {
		return nil, err
	}
	return fs.ReadProfile(id, int(v))
}

// ReadLatestProfilesForOwner ...
func (fs *FileStore) ReadLatestProfilesForOwner(owner string) ([]*models.Profile, error) {
	files, err := fs.listProfilesForOwner(owner)
	if err != nil {
		return nil, err
	}
	if len(files) == 0 {
		return nil, nil
	}
	fmt.Printf("\tFound :%d\n", len(files))

	// Get the latest versions
	mapProfiles := make(map[string]*record)
	for _, f := range files {
		fmt.Printf("\tWorking on :%s\n", f)

		item := parse(f)
		other, ok := mapProfiles[item.id]
		if !ok {
			fmt.Printf("\tAdding %s\n", f)
			mapProfiles[item.id] = item
		} else if other.version < item.version {
			fmt.Printf("\tReplacing %s with %s\n", other.fname, item.fname)
			mapProfiles[item.id] = item
		} else {
			fmt.Printf("\tBAD %s with %s\n", other.fname, item.fname)
			fmt.Printf("\tBAD %d with %d\n", other.version, item.version)
		}
	}

	// Now load the profiles
	profiles := make([]*models.Profile, 0)
	for _, v := range mapProfiles {
		// path := store.pathToProfileVersionOwner(v.id, v.version, v.owner)
		fmt.Printf("\tGoing to read %s\n", v.fname)
		if id, v, _, err := extractIDVersionOwner(v.fname); err == nil {
			if profile, err := fs.ReadProfile(id, v); err == nil {
				profiles = append(profiles, profile)
			}
		}
	}
	fmt.Printf("\tPreparing Results %d items\n", len(mapProfiles))
	return profiles, err
}

type record struct {
	version int
	id      string
	fname   string
	owner   string
}

func parse(fname string) *record {
	parts := strings.Split(fname, ".")
	v, _ := strconv.Atoi(parts[2][1:])
	fmt.Printf("\tPARTS: %s\n", parts)

	return &record{
		owner:   parts[0],
		id:      parts[1],
		version: v,
		fname:   fname,
	}
}
