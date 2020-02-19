package mango_objects

import "time"

type Record struct {
	CallID         string    `json:"call_id"`
	ProductID      int       `json:"product_id"`
	UserID         int       `json:"user_id"`
	Timestamp      int64     `json:"timestamp"`
	Time           time.Time `json:"-"`
	RecordingID    string    `json:"recording_id"`
	RecordingState string    `json:"recording_state"`
	Seq            int       `json:"seq"`
	EntryID        string    `json:"entry_id"`
	Extension      string    `json:"extension"`
	CompletionCode int       `json:"completion_code"`
	Recipient      string    `json:"recipient"`
	CommandID      string    `json:"command_id"`
}
