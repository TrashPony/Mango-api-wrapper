package mango_client

import "time"

type State struct {
	DateFrom int64  `json:"date_from"`
	DateTo   int64  `json:"date_to"`
	Fields   string `json:"fields"`
}

func GetState(dateFrom, dateTo time.Time, fields string) {
	//State := State{
	//	DateFrom: dateFrom.Unix(),
	//	DateTo:   dateTo.Unix(),
	//	Fields:   fields,
	//}
}
