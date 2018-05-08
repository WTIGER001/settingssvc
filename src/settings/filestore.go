package settings

import (
	"models"

	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

var store = &FileStore{
	dir: "c:/dev/fs",
}

// FileStore ...
type FileStore struct {
	dir string
}

func (fs *FileStore) writeProfile(item *models.Profile) error {
	fmt.Printf("\tWriting Profile \n")

	path := fs.pathToProfileVersion(item.ID, int(item.Version))
	fmt.Printf("Writing Proflie to file %s\n", path)

	return fs.write(path, item)
}

func (fs *FileStore) readProfile(path string) (*models.Profile, error) {
	fmt.Printf("Reading File\n")
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("Error reading file: %s\n", err.Error())
		return nil, err
	}

	// Unmarshall "most" of this object.
	result := &models.Profile{}
	err = json.Unmarshal(bytes, result)
	if err != nil {
		return nil, err
	}

	// now read the value
	err = readProfile(bytes, result)

	return result, nil
}

func (fs *FileStore) writeOwner(item *models.PreferenceOwner) error {

	path := fs.pathToOwner(item.ID)
	fmt.Printf("Writing Owner to file %s\n", path)

	return fs.write(path, item)
}

func (fs *FileStore) readOwner(path string) (*models.PreferenceOwner, error) {
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

func (fs *FileStore) writeCategory(item *models.Category) error {

	path := fs.pathToCategory(item.Name)
	fmt.Printf("Writing Category to file %s\n", path)

	return fs.write(path, item)
}

func (fs *FileStore) readCategory(path string) (*models.Category, error) {
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

func (fs *FileStore) writeType(item *models.OwnerType) error {

	path := fs.pathToType(item.ID)
	fmt.Printf("Writing Type to file %s\n", path)

	return fs.write(path, item)
}

func (fs *FileStore) readType(path string) (*models.OwnerType, error) {
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

func (fs *FileStore) writeDefinition(item *models.PreferenceDefinition) error {

	path := fs.pathToDefinition(item.ID)
	fmt.Printf("Writing Type to file %s\n", path)

	return fs.write(path, item.JsonData)
}

func (fs *FileStore) readDefinition(path string) (*models.PreferenceDefinition, error) {
	fmt.Printf("Reading File\n")
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("Errpr reading file: %s\n", err.Error())
		return nil, err
	}

	// Unmarshall "most" of this object.
	result := &models.PreferenceDefinition{}

	err = readPreferenceDefinition(bytes, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (fs *FileStore) exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	fmt.Printf("Error : %s", err.Error())
	return false
}

func (fs *FileStore) pathToDefinition(id string) string {
	return fs.dir + "/definitions/" + id + ".definition.json"
}
func (fs *FileStore) pathToType(id string) string {
	return fs.dir + "/types/" + id + ".type.json"
}
func (fs *FileStore) pathToCategory(id string) string {
	return fs.dir + "/categories/" + id + ".category.json"
}
func (fs *FileStore) pathToOwner(id string) string {
	return fs.dir + "/owners/" + id + ".owner.json"
}
func (fs *FileStore) pathToProfileVersion(id string, version int) string {
	return fmt.Sprintf("%s/profiles/%s.v%d.profile.json", fs.dir, id, version)
}
func (fs *FileStore) pathToProfile(id string) string {
	maxVersion := fs.LatestProfileVersion(id)
	return fs.pathToProfileVersion(id, maxVersion)
}

// LatestProfileVersion finds the largest version for a single profile.
// A return of 0 means that no version was found
func (fs *FileStore) LatestProfileVersion(id string) int {
	fmt.Printf("Listing Profiles\n")

	files, err := fs.ListProfiles(id)
	if err != nil {
		return 0
	}
	fmt.Printf("\tfound %d\n", len(files))

	var max int
	for _, f := range files {
		v, err := ExtractVersion(f)
		fmt.Printf("\tfound version %d\n", v)

		if err != nil {
			fmt.Printf("error decoding version string\n")
		} else {
			if v > max {
				max = v
			}
		}
	}

	return max
}

//ProfileVersionDate gets the date of a profile version
func (fs *FileStore) ProfileVersionDate(profileName string) (time.Time, error) {
	path := fs.dir + "/owners/" + profileName
	return fs.getFileDate(path)
}

func (fs *FileStore) getFileDate(path string) (time.Time, error) {
	stat, err := os.Stat(path)
	if err != nil {
		return time.Now(), err
	}
	return stat.ModTime(), nil
}

// ExtractVersion gets the version from a filename or path
func ExtractVersion(filename string) (int, error) {
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
		return 0, err
	}
	fmt.Printf("\t%d\n", v)

	return v, nil
}

//ListProfiles for an id
func (fs *FileStore) ListProfiles(id string) ([]string, error) {
	part := fs.dir + "/profiles/" + id + "*.profile.json"
	files, err := filepath.Glob(part)
	return files, err
}

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
