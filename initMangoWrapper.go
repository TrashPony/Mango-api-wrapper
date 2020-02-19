package mango

import (
	"github.com/TrashPony/Mango-api-wrapper/mango_events"
	"net/http"
	"sync"
)

var ClientMango *Client
var once sync.Once

func InitMangoCallHandle(port, apiKey, sing, url string) *Client {

	once.Do(func() {

		ClientMango = &Client{}
		ClientMango.SetApiKey(apiKey)
		ClientMango.SetApiSing(sing)
		ClientMango.SetApiUrl(url)

		// Уведомление о вызове
		http.HandleFunc("/events/call", mango_events.EventCall)

		// Уведомление о нажатиях DTMF клавиш
		http.HandleFunc("/events/dtmf", mango_events.PressDTMF)

		// Уведомление о завершении вызова
		http.HandleFunc("/events/summary", mango_events.EndCall)

		// Уведомление о записи разговора
		http.HandleFunc("/events/recording", mango_events.RecordStart)

		// Уведомление о помещении записи разговора в облачное хранилище
		http.HandleFunc("/events/record/added", mango_events.RecordAdd)

		// Уведомление о результате отправки SMS
		http.HandleFunc("/events/sms", mango_events.ResultSMS)

		go func() {
			err := http.ListenAndServe(":"+port, nil)
			if err != nil {
				println(err.Error())
			}
		}()
	})

	return ClientMango
}
