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

func GetBalance() *Balance {

	jsonResp := mango_request.RequestToMango("{}", "account/balance")

	b := Balance{}
	err := json.Unmarshal([]byte(jsonResp), &b)
	if err != nil {
		panic(err)
	}

	return &b
}
