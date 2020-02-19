package mango_commands

import (
	"github.com/TrashPony/Mango-api-wrapper/mango_objects"
	"github.com/TrashPony/Mango-api-wrapper/mango_request"
	"github.com/google/uuid"
)

//Инициирование вызова от имени сотрудника
func InitCallOfUser(extension, callerNumber, toNumber, lineNumber, sipHeaders, apiKey, apiSing, apiUrl string) *mango_objects.Result {
	uuidCall := uuid.New().String()

	call := mango_objects.Call{
		CommandID: uuidCall,
		From: &mango_objects.From{
			Extension: extension,
			Number:    callerNumber,
		},
		LineNumber: lineNumber,
		ToNumber:   toNumber,
		SipHeaders: sipHeaders,
	}

	jsonResp := mango_request.RequestToMango(call.ToJSON(), "commands/callback", apiKey, apiSing, apiUrl)

	return mango_request.ParseResult(jsonResp)
}
