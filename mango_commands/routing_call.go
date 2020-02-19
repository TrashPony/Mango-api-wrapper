package mango_commands

import (
	"github.com/TrashPony/Mango-api-wrapper/mango_objects"
	"github.com/TrashPony/Mango-api-wrapper/mango_request"
	"github.com/google/uuid"
)

//Маршрутизация вызова
func RoutingCall(callID, toNumber, sipHeaders string) *mango_objects.Result {
	uuidCall := uuid.New().String()

	call := mango_objects.Call{
		CommandID:  uuidCall,
		CallID:     callID,
		ToNumber:   toNumber,
		SipHeaders: sipHeaders,
	}

	jsonResp := mango_request.RequestToMango(call.ToJSON(), "commands/route")
	return mango_request.ParseResult(jsonResp)
}
