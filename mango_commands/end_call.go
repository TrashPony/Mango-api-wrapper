package mango_commands

import (
	"github.com/TrashPony/Mango-api-wrapper/mango_objects"
	"github.com/TrashPony/Mango-api-wrapper/mango_request"
	"github.com/google/uuid"
)

//Завершение вызова
func EndCall(callID, apiKey, apiSing, apiUrl string) *mango_objects.Result {
	uuidCall := uuid.New().String()

	call := mango_objects.Call{
		CommandID: uuidCall,
		CallID:    callID,
	}

	jsonResp := mango_request.RequestToMango(call.ToJSON(), "commands/call/hangup", apiKey, apiSing, apiUrl)
	return mango_request.ParseResult(jsonResp)
}
