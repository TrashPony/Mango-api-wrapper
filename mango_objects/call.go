package mango_objects

import (
	"encoding/json"
	"fmt"
	"time"
)

type Call struct {
	EntryID           string    `json:"entry_id"`
	CallID            string    `json:"call_id"`
	Timestamp         int64     `json:"timestamp"`
	Time              time.Time `json:"-"`
	Seq               int       `json:"seq"`
	CallState         string    `json:"call_state"`
	Location          string    `json:"location"`
	From              *From     `json:"from"`
	To                *To       `json:"to"`
	DCT               *DCT      `json:"dct"`
	DisconnectReason  int       `json:"disconnect_reason"`
	CommandID         string    `json:"command_id"`
	TaskID            string    `json:"task_id"`
	CallbackInitiator string    `json:"callback_initiator"`
	CallDirection     int       `json:"call_direction"`
	LineNumber        string    `json:"line_number"`
	CreateTimeStamp   int64     `json:"create_time"`
	CreateTime        time.Time `json:"-"`
	ForwardTimeStamp  int64     `json:"forward_time"`
	ForwardTime       time.Time `json:"-"`
	TalkTimeStamp     int64     `json:"talk_time"`
	TalkTime          time.Time `json:"-"`
	EndTimeStamp      int64     `json:"end_time"`
	EndTime           time.Time `json:"-"`
	EntryResult       int       `json:"entry_result"`
	ToNumber          string    `json:"to_number"`
	SipHeaders        string    `json:"sip_headers"`
	CallPartyNumber   string    `json:"call_party_number"`
	Method            string    `json:"method"`
	Initiator         string    `json:"initiator"`
}

type From struct {
	Extension       string `json:"extension"`
	Number          string `json:"number"`
	TakenFromCallID string `json:"taken_from_call_id"`
}

type To struct {
	Extension  string `json:"extension"`
	Number     string `json:"number"`
	LineNumber string `json:"line_number"`
	AcdGroup   string `json:"acd_group"`
}

type DCT struct {
	Number string `json:"number"`
	Type   int    `json:"type"`
}

func (c *Call) ToJSON() string {
	callJson, err := json.Marshal(c)
	if err != nil {
		fmt.Println(err)
	}

	return string(callJson)
}
