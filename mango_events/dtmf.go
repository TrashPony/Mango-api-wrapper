package mango_events

import (
	"encoding/json"
	"github.com/TrashPony/Mango-api-wrapper/mango_objects"
	"github.com/TrashPony/Mango-api-wrapper/mango_request"
	"net/http"
	"time"
)

func PressDTMF(w http.ResponseWriter, r *http.Request) {
	request := mango_request.ParseRequest(r)

	dtmf := mango_objects.DTMF{}
	err := json.Unmarshal([]byte(request.Json), &dtmf)
	if err != nil {
		panic(err)
	}

	dtmf.Time = time.Unix(dtmf.Timestamp, 0)
	Events.DTMF <- &dtmf
}
