package mango_commands

import (
	"github.com/TrashPony/Mango-api-wrapper/mango_objects"
	"github.com/TrashPony/Mango-api-wrapper/mango_request"
	"github.com/google/uuid"
)

// Включение записи разговора
func TurnOnRecordingCall(callID, callPartyNumber string) *mango_objects.Result {
	uuidCall := uuid.New().String()

	call := mango_objects.Call{
		CommandID:       uuidCall,
		CallID:          callID,
		CallPartyNumber: callPartyNumber,
	}

	jsonResp := mango_request.RequestToMango(call.ToJSON(), "commands/recording/start")
	return mango_request.ParseResult(jsonResp)
}
