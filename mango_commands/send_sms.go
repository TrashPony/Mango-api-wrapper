package mango_commands

import (
	"github.com/TrashPony/Mango-api-wrapper/mango_objects"
	"github.com/TrashPony/Mango-api-wrapper/mango_request"
	"github.com/google/uuid"
)

//Отправка SMS
func SendSms(fromExtension, toNumber, smsSender, text string) *mango_objects.Result {
	uuidCall := uuid.New().String()

	sms := mango_objects.SMS{
		CommandID:     uuidCall,
		Text:          text,
		FromExtension: fromExtension,
		ToNumber:      toNumber,
		SmsSender:     smsSender,
	}

	jsonResp := mango_request.RequestToMango(sms.ToJSON(), "commands/sms")
	return mango_request.ParseResult(jsonResp)
}
