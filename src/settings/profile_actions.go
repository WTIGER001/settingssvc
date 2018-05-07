package settings

import (
	"fmt"
	"models"
	"os"
	"restapi/operations/preferences"

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

	id := params.ID
	version := 0
	if params.Version != nil {
		version = int(*params.Version)
	} else {
		version = store.LatestProfileVersion(id)
	}

	path := store.pathToProfileVersion(id, version)

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
		v, err := ExtractVersion(f)

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

	ids := params.ID

	profiles := make([]*models.Profile, 0, len(ids))
	for _, id := range ids {
		version := store.LatestProfileVersion(id)
		path := store.pathToProfileVersion(id, version)

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

// UpdateProfile creates a new version each time
func UpdateProfile(params preferences.UpdateProfileParams) middleware.Responder {
	fmt.Printf("Starting --> UpdateProfile id:%s\n", params.Body.ID)

	id := params.Body.ID
	lastVersion := store.LatestProfileVersion(id)
	version := lastVersion + 1
	fmt.Printf("\tUpdating version %d to %d\n", lastVersion, version)

	profile := params.Body
	profile.Version = int64(version)

	err := store.writeProfile(profile)
	if err != nil {
		return preferences.NewUpdateProfileInternalServerError()
	}

	response := preferences.NewUpdateProfileOK()
	return response
}
