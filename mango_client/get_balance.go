package mango_client

import (
	"encoding/json"
	"github.com/TrashPony/Mango-api-wrapper/mango_request"
)

type Balance struct {
	Balance    float64 `json:"balance"`
	Currency   string  `json:"currency"`
	ResponseAt string  `json:"response_at"`
	Result     float64 `json:"result"`
}

func GetBalance(apiKey, apiSing, apiUrl string) *Balance {

	jsonResp := mango_request.RequestToMango("{}", "account/balance", apiKey, apiSing, apiUrl)

	b := Balance{}
	err := json.Unmarshal([]byte(jsonResp), &b)
	if err != nil {
		panic(err)
	}

	return &b
}
