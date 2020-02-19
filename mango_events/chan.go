package mango_events

import "github.com/TrashPony/Mango-api-wrapper/mango_objects"

var Events = eventsChanelInit()

type eventsChanel struct {
	Calls       chan *mango_objects.Call
	DTMF        chan *mango_objects.DTMF
	EndCalls    chan *mango_objects.Call
	AddRecord   chan *mango_objects.Record
	StartRecord chan *mango_objects.Record
	SMS         chan *mango_objects.SMS
}

func eventsChanelInit() *eventsChanel {
	return &eventsChanel{
		Calls:       make(chan *mango_objects.Call),
		DTMF:        make(chan *mango_objects.DTMF),
		EndCalls:    make(chan *mango_objects.Call),
		AddRecord:   make(chan *mango_objects.Record),
		StartRecord: make(chan *mango_objects.Record),
		SMS:         make(chan *mango_objects.SMS),
	}
}
