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

	path := fs.pathToProfileVersionOwner(item.ID, int(item.Version), item.Owner)
	fmt.Printf("Writing profile to file %s\n", path)

	return fs.write(path, item)
}

func (fs *FileStore) readProfile(path string) (*models.Profile, error) {
	fmt.Printf("Reading File\n")
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
	err = readProfile(bytes, result)
	fmt.Printf("\tComplete %s\n", result.ID)

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
	file, _ := fs.ListProfileVersion(id, version)
	return file
}
func (fs *FileStore) pathToProfileVersionOwner(id string, version int, owner string) string {
	return fmt.Sprintf("%s/profiles/%s.%s.v%d.profile.json", fs.dir, owner, id, version)
}
func (fs *FileStore) pathToProfile(id string) string {
	maxVersion, owner := fs.LatestProfileVersion(id)
	return fs.pathToProfileVersionOwner(id, maxVersion, owner)
}

// LatestProfileVersion finds the largest version for a single profile.
// A return of 0 means that no version was found
func (fs *FileStore) LatestProfileVersion(id string) (int, string) {
	fmt.Printf("Listing Profiles\n")
	var owner string
	files, err := fs.ListProfiles(id)
	if err != nil {
		return 0, ""
	}
	fmt.Printf("\tfound %d\n", len(files))

	var max int
	for _, f := range files {
		v, _, err := ExtractVersionOwner(f)
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
func (fs *FileStore) ProfileVersionDate(profileName string) (time.Time, error) {
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
func ExtractVersionOwner(filename string) (int, string, error) {
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
		return 0, "", err
	}
	fmt.Printf("\t%d\n", v)

	return v, parts[0], nil
}

//ListProfiles for an id
func (fs *FileStore) ListProfiles(id string) ([]string, error) {
	part := fs.dir + "/profiles/*." + id + ".v*.profile.json"
	files, err := filepath.Glob(part)
	return files, err
}

// ListProfilesForOwner ...
func (fs *FileStore) ListProfilesForOwner(owner string) ([]string, error) {
	part := fmt.Sprintf("%s/profiles/%s.*.*.profile.json", fs.dir, owner)

	files, err := filepath.Glob(part)
	return files, err
}

// ListProfileVersion ...
func (fs *FileStore) ListProfileVersion(id string, version int) (string, error) {
	// part := fs.dir + "/profiles/*." + id + ".v" + .profile.json"
	part := fmt.Sprintf("%s/profiles/*.%s.v%d.profile.json", fs.dir, id, version)
	files, err := filepath.Glob(part)
	if len(files) == 1 {
		return fmt.Sprintf("%s/profiles/%s", fs.dir, files[0]), err
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
