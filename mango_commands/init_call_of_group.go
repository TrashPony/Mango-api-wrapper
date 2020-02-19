package mango_commands

import (
	"encoding/json"
	"fmt"
	"github.com/TrashPony/Mango-api-wrapper/mango_objects"
	"github.com/TrashPony/Mango-api-wrapper/mango_request"
	"github.com/google/uuid"
)

// Инициирование вызова от имени группы
func InitCallOfGroup(from, to, lineNumber, apiKey, apiSing, apiUrl string) *mango_objects.Result {
	uuidCall := uuid.New().String()

	// TODO из за того что тут требуются данные которые не матчатся с обычным звонком пришлось вот так делоть
	type GroupCall struct {
		CommandID  string `json:"command_id"`
		From       string `json:"from"`
		To         string `json:"to"`
		LineNumber string `json:"line_number"`
	}

	groupCall := GroupCall{
		CommandID:  uuidCall,
		From:       from,
		To:         to,
		LineNumber: lineNumber,
	}

	callJson, err := json.Marshal(groupCall)
	if err != nil {
		fmt.Println(err)
	}

	jsonResp := mango_request.RequestToMango(string(callJson), "commands/callback_group", apiKey, apiSing, apiUrl)

	return mango_request.ParseResult(jsonResp)
}
