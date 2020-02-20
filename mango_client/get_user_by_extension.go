package mango_client

import (
	"encoding/json"
	"github.com/TrashPony/Mango-api-wrapper/mango_objects"
	"github.com/TrashPony/Mango-api-wrapper/mango_request"
)

func GetUserByExtension(extension, apiKey, apiSing, apiUrl string) *mango_objects.User {

	user := mango_objects.User{
		Extension: extension,
	}

	jsonResp := mango_request.RequestToMango(user.ToJSON(), "config/users/request", apiKey, apiSing, apiUrl)

	users := mango_objects.Users{}

	err := json.Unmarshal([]byte(jsonResp), &users)
	if err != nil {
		panic(err)
	}

	if len(users.Users) == 0 {
		return nil
	}
	return users.Users[0]
}
