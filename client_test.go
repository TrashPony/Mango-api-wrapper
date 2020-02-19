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
