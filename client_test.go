package mango

import (
	"testing"
)

func TestGetBalance(t *testing.T) {
	client := InitMangoCallHandle("", "", "", "")
	balance := client.GetBalance()
	println(balance.Balance)
}

func TestSMS(t *testing.T) {
	client := InitMangoCallHandle("", "", "", "")
	result := client.SendSms("", "", "", "")
	println(result.Result)
}

func TestGetUser(t *testing.T) {
	client := InitMangoCallHandle("", "", "", "")
	user := client.GetUserByExtension("200")
	println(user)
}

func TestGetUsers(t *testing.T) {
	client := InitMangoCallHandle("", "", "", "")
	users := client.GetUsers()
	println(users)
}
