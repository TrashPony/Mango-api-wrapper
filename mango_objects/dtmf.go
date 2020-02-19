package mango_objects

import "time"

type DTMF struct {
	Seq       int       `json:"seq"`
	DTMF      string    `json:"dtmf"`
	Timestamp int64     `json:"timestamp"`
	Time      time.Time `json:"time"`
	CallID    string    `json:"call_id"`
	EntryID   string    `json:"entry_id"`
	Location  string    `json:"location"`
}
