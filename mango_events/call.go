package mango_events

import (
	"encoding/json"
	"github.com/TrashPony/Mango-api-wrapper/mango_objects"
	"github.com/TrashPony/Mango-api-wrapper/mango_request"
	"net/http"
	"time"
)

func EventCall(w http.ResponseWriter, r *http.Request) {
	request := mango_request.ParseRequest(r)

	eventCall := mango_objects.Call{}
	err := json.Unmarshal([]byte(request.Json), &eventCall)
	if err != nil {
		panic(err)
	}

	eventCall.Time = time.Unix(eventCall.Timestamp, 0)

	Events.AddCall(&eventCall)
}
