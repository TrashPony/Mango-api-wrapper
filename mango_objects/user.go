package mango_objects

import (
	"encoding/json"
	"fmt"
)

type Users struct {
	Users []*User `json:"users"`
}

type User struct {
	General   *UserGeneral   `json:"general"`
	Telephony *UserTelephony `json:"telephony"`
	Extension string         `json:"extension"`
}

type UserGeneral struct {
	Name       string `json:"name"`
	Email      string `json:"email"`
	Department string `json:"department"`
	Position   string `json:"position"`
}

type UserTelephony struct {
	Extension    string     `json:"extension"`
	OutGoingLine string     `json:"outgoingline"`
	Numbers      []*Numbers `json:"numbers"`
}

type Numbers struct {
	Number   string `json:"number"`
	Protocol string `json:"protocol"`
	Order    int    `json:"order"`
	WaitSec  int    `json:"wait_sec"`
	Status   string `json:"status"`
}

func (u *User) ToJSON() string {
	callJson, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}

	return string(callJson)
}
