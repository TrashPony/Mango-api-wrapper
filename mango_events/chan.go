package mango_events

import (
	"github.com/TrashPony/Mango-api-wrapper/mango_objects"
)

var Events = eventsChanelInit()

const cacheCanal = 250

type eventsChanel struct {
	calls             chan *mango_objects.Call
	callsExit         chan *mango_objects.Call
	dtmf              chan *mango_objects.DTMF
	dtmfExit          chan *mango_objects.DTMF
	endCalls          chan *mango_objects.Call
	endCallsExit      chan *mango_objects.Call
	addedRecord       chan *mango_objects.Record
	addedRecordExit   chan *mango_objects.Record
	startedRecord     chan *mango_objects.Record
	startedRecordExit chan *mango_objects.Record
	sms               chan *mango_objects.SMS
	smsExit           chan *mango_objects.SMS
}

func eventsChanelInit() *eventsChanel {
	return &eventsChanel{
		calls:         make(chan *mango_objects.Call, cacheCanal),
		dtmf:          make(chan *mango_objects.DTMF, cacheCanal),
		endCalls:      make(chan *mango_objects.Call, cacheCanal),
		addedRecord:   make(chan *mango_objects.Record, cacheCanal),
		startedRecord: make(chan *mango_objects.Record, cacheCanal),
		sms:           make(chan *mango_objects.SMS, cacheCanal),
	}
}

func (e *eventsChanel) AddCall(call *mango_objects.Call) {
	if len(e.calls) < cacheCanal {
		// если информацию не считывают то мы ее не вносим
		// если надо больше то читайте каналы и сохраняйте например в бд
		e.calls <- call
	}
}

func (e *eventsChanel) AddDTMF(dtmf *mango_objects.DTMF) {
	if len(e.dtmf) < cacheCanal {
		e.dtmf <- dtmf
	}
}

func (e *eventsChanel) AddEndCall(call *mango_objects.Call) {
	if len(e.endCalls) < cacheCanal {
		e.endCalls <- call
	}
}

func (e *eventsChanel) AddRecord(record *mango_objects.Record) {
	if len(e.addedRecord) < cacheCanal {
		e.addedRecord <- record
	}
}

func (e *eventsChanel) AddStartRecord(record *mango_objects.Record) {
	if len(e.startedRecord) < cacheCanal {
		e.startedRecord <- record
	}
}

func (e *eventsChanel) AddSMS(sms *mango_objects.SMS) {
	if len(e.sms) < cacheCanal {
		e.sms <- sms
	}
}

func (e *eventsChanel) GetAddCallChan() <-chan *mango_objects.Call {
	if e.callsExit == nil {
		e.callsExit = make(chan *mango_objects.Call)

		go func() {
			for {
				call := <-e.calls
				e.callsExit <- call
			}
		}()
	}

	return e.callsExit
}

func (e *eventsChanel) GetDTMF() <-chan *mango_objects.DTMF {
	if e.dtmfExit == nil {
		e.dtmfExit = make(chan *mango_objects.DTMF)

		go func() {
			for {
				dtmf := <-e.dtmf
				e.dtmfExit <- dtmf
			}
		}()
	}

	return e.dtmfExit
}

func (e *eventsChanel) GetEndCalls() <-chan *mango_objects.Call {
	if e.endCallsExit == nil {
		e.endCallsExit = make(chan *mango_objects.Call)

		go func() {
			for {
				call := <-e.endCalls
				e.endCallsExit <- call
			}
		}()
	}

	return e.endCallsExit
}

func (e *eventsChanel) GetAddedRecords() <-chan *mango_objects.Record {
	if e.addedRecordExit == nil {
		e.addedRecordExit = make(chan *mango_objects.Record)

		go func() {
			for {
				record := <-e.addedRecord
				e.addedRecordExit <- record
			}
		}()
	}

	return e.addedRecordExit
}

func (e *eventsChanel) GetStartedRecord() <-chan *mango_objects.Record {
	if e.startedRecordExit == nil {
		e.startedRecordExit = make(chan *mango_objects.Record)

		go func() {
			for {
				record := <-e.startedRecord
				e.startedRecordExit <- record
			}
		}()
	}

	return e.startedRecordExit
}

func (e *eventsChanel) GetSMS() <-chan *mango_objects.SMS {
	if e.smsExit == nil {
		e.smsExit = make(chan *mango_objects.SMS)

		go func() {
			for {
				record := <-e.sms
				e.smsExit <- record
			}
		}()
	}

	return e.smsExit
}
