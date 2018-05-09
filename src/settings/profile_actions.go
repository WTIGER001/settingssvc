package settings

import (
	"fmt"
	"models"
	"os"
	"restapi/operations/preferences"
	"strconv"
	"strings"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteProfile deletes all versions of a profile
func DeleteProfile(params preferences.DeleteProfileParams) middleware.Responder {
	fmt.Printf("Starting --> DeleteProfile id:%s\n", params.ID)

	id := params.ID
	files, err := store.ListProfiles(id)
	if err != nil {
		return preferences.NewDeleteProfileInternalServerError()
	}
	for _, f := range files {
		os.Remove(f)
	}

	return preferences.NewDeleteProfileOK()
}

// GetMyActiveProfile looks at the JWT to determine the user id and then gives back the profile under "Active"
func GetMyActiveProfile(params preferences.GetMyActiveProfileParams) middleware.Responder {

	return middleware.NotImplemented("operation preferences.GetMyActiveProfile has not yet been implemented")
}

//GetProfile gets a single profile
func GetProfile(params preferences.GetProfileParams) middleware.Responder {
	fmt.Printf("Starting --> GetProfile id:%s, version:%d\n", params.ID, params.Version)

	var path string
	id := params.ID
	version := 0
	owner := ""
	if params.Version != nil {
		version = int(*params.Version)
	} else {
		version, owner = store.LatestProfileVersion(id)
	}

	if owner != "" {
		path = store.pathToProfileVersionOwner(id, version, owner)
	} else {
		path = store.pathToProfileVersion(id, version)
	}

	if !store.exists(path) {
		return preferences.NewGetProfileNotFound()
	}

	profile, err := store.readProfile(path)
	if err != nil {
		return preferences.NewGetProfileInternalServerError()
	}

	response := preferences.NewGetProfileOK()
	response.SetPayload(profile)
	return response
}

//GetProfileVersions returns a list of all versions for a given id
func GetProfileVersions(params preferences.GetProfileVersionsParams) middleware.Responder {
	fmt.Printf("Starting --> GetProfileVersions id:%s\n", params.ID)

	versions := &models.ProfileVersions{}

	id := params.ID
	files, err := store.ListProfiles(id)
	if err != nil {
		return preferences.NewDeleteProfileInternalServerError()
	}
	fmt.Printf("\tFound %d files\n", len(files))

	table := make([]*models.ProfileVersion, 0, len(files))
	for _, f := range files {
		version := &models.ProfileVersion{}
		v, _, err := ExtractVersionOwner(f)

		if err == nil {
			fmt.Printf("\tFound Version %d\n", v)
			version.Version = int64(v)
			table = append(table, version)
			// _, err := store.ProfileVersionDate(f)
			// if err != nil {
			// 	table = append(table, version)
			// }
		} else {
			fmt.Printf("\tError: %s\n", err.Error())
		}
	}
	versions.ID = id
	versions.Versions = table
	response := preferences.NewGetProfileVersionsOK()
	response.SetPayload(versions)
	return response
}

//GetProfiles returns all the latest profiles in the list
func GetProfiles(params preferences.GetProfilesParams) middleware.Responder {
	fmt.Printf("Starting --> GetProfiles id:%s\n", params.ID)

	if len(params.ID) > 0 {
		ids := params.ID

		profiles := make([]*models.Profile, 0, len(ids))
		for _, id := range ids {
			version, owner := store.LatestProfileVersion(id)
			path := store.pathToProfileVersionOwner(id, version, owner)

			if !store.exists(path) {
				// SKIP this on
				continue
			}

			profile, err := store.readProfile(path)
			if err != nil {
				// SKip this one
				continue
			}

			profiles = append(profiles, profile)
		}

		response := preferences.NewGetProfilesOK()
		response.SetPayload(profiles)
		return response
	}
	fmt.Printf("Starting --> GetProfiles Owner:%s\n", *params.Ownerid)

	profiles := make([]*models.Profile, 0, 10)
	files, err := store.ListProfilesForOwner(*params.Ownerid)
	if err != nil {
		return preferences.NewGetProfilesInternalServerError()
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
	for _, v := range mapProfiles {
		// path := store.pathToProfileVersionOwner(v.id, v.version, v.owner)
		fmt.Printf("\tGoing to read %s\n", v.fname)
		profile, err := store.readProfile(v.fname)
		if err != nil {
			// SKip this one
			continue
		}
		profiles = append(profiles, profile)
	}
	fmt.Printf("\tPreparing Results %d items\n", len(mapProfiles))

	response := preferences.NewGetProfilesOK()
	response.SetPayload(profiles)
	return response

}

// UpdateProfile creates a new version each time
func UpdateProfile(params preferences.UpdateProfileParams) middleware.Responder {
	fmt.Printf("Starting --> UpdateProfile id:%s\n", params.Body.ID)

	id := params.Body.ID
	lastVersion, _ := store.LatestProfileVersion(id)
	version := lastVersion + 1
	fmt.Printf("\tUpdating version %d to %d\n", lastVersion, version)

	profile := params.Body
	profile.Version = int64(version)

	// newProfile := !store.exists(store.pathToProfile(profile.ID))

	err := store.writeProfile(profile)
	if err != nil {
		return preferences.NewUpdateProfileInternalServerError()
	}

	response := preferences.NewUpdateProfileOK()
	return response
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
