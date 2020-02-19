package mango_objects

import (
	"encoding/json"
	"fmt"
	"time"
)

type SMS struct {
	CommandID     string    `json:"command_id"`
	Timestamp     int64     `json:"timestamp"`
	Time          time.Time `json:"-"`
	Reason        int       `json:"reason"`
	Text          string    `json:"text"`
	FromExtension string    `json:"from_extension"`
	ToNumber      string    `json:"to_number"`
	SmsSender     string    `json:"sms_sender"`
}

func (s *SMS) ToJSON() string {
	callJson, err := json.Marshal(s)
	if err != nil {
		fmt.Println(err)
	}

	return string(callJson)
}
