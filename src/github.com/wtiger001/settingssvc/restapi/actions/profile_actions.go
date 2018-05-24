package actions

import (
	"fmt"
	"github.com/wtiger001/settingssvc/models"
	"github.com/wtiger001/settingssvc/restapi/operations/preferences"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteProfile deletes all versions of a profile
func DeleteProfile(params preferences.DeleteProfileParams) middleware.Responder {
	fmt.Printf("Starting --> DeleteProfile id:%s\n", params.ID)

	id := params.ID

	err := store.DeleteProfile(id)
	if err != nil {
		fmt.Printf("\tError: %s\n", err.Error())
		return preferences.NewDeleteProfileInternalServerError()
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
		v, err := store.LatestVersion(id)
		if err != nil {
			fmt.Printf("\tError: %s\n", err.Error())
			return preferences.NewGetProfileInternalServerError()
		}
		version = v
	}

	profile, err := store.ReadProfile(id, version)
	if err != nil {
		fmt.Printf("\tError: %s\n", err.Error())
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
	vArr, err := store.ProfileVersions(id)
	if err != nil {
		fmt.Printf("\tError: %s\n", err.Error())
		return preferences.NewDeleteProfileInternalServerError()
	}
	fmt.Printf("\tFound %d files\n", len(vArr))

	table := make([]*models.ProfileVersion, 0, len(vArr))
	for _, value := range vArr {
		version := &models.ProfileVersion{}
		version.Version = int64(value)
		table = append(table, version)
	}

	versions.ID = id
	versions.Versions = table
	response := preferences.NewGetProfileVersionsOK()
	response.SetPayload(versions)
	return response
}

//GetProfiles returns all the latest profiles in the list
func GetProfiles(params preferences.GetProfilesParams) middleware.Responder {
	fmt.Printf("Starting --> GetProfiles id:%s and %d\n", params.ID, len(params.ID))

	if len(params.ID) > 0 {
		fmt.Printf("\tIn GetPRofiles by ID\n")

		ids := params.ID

		profiles := make([]*models.Profile, 0, len(ids))
		for _, id := range ids {
			profile, err := store.ReadLatestProfile(id)
			if err != nil {
				fmt.Printf("\tError: %s\n", err.Error())
				continue
			}
			profiles = append(profiles, profile)
		}

		response := preferences.NewGetProfilesOK()
		response.SetPayload(profiles)
		return response
	} else if params.Ownerid != nil {

		fmt.Printf("Starting --> GetProfiles Owner:%s\n", *params.Ownerid)
		owner := *params.Ownerid
		profiles, err := store.ReadLatestProfilesForOwner(owner)
		if err != nil {
			fmt.Printf("\tError: %s\n", err.Error())
			return preferences.NewGetProfilesInternalServerError()
		}

		response := preferences.NewGetProfilesOK()
		response.SetPayload(profiles)
		return response
	}
	return preferences.NewGetProfilesBadRequest()

}

// UpdateProfile creates a new version each time
func UpdateProfile(params preferences.UpdateProfileParams) middleware.Responder {
	fmt.Printf("Starting --> UpdateProfile id:%s\n", params.Body.ID)

	err := store.StoreProfile(params.Body)
	if err != nil {
		fmt.Printf("\tError: %s\n", err.Error())
		return preferences.NewUpdateProfileInternalServerError()
	}

	response := preferences.NewUpdateProfileOK()
	return response
}
