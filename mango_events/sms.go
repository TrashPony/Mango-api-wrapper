package mango_events

import (
	"encoding/json"
	"github.com/TrashPony/Mango-api-wrapper/mango_objects"
	"github.com/TrashPony/Mango-api-wrapper/mango_request"
	"net/http"
	"time"
)

func ResultSMS(w http.ResponseWriter, r *http.Request) {
	request := mango_request.ParseRequest(r)

	sms := mango_objects.SMS{}
	err := json.Unmarshal([]byte(request.Json), &sms)
	if err != nil {
		panic(err)
	}

	sms.Time = time.Unix(sms.Timestamp, 0)
	Events.SMS <- &sms
}
