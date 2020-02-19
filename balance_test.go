package mango

import (
	"github.com/TrashPony/Mango-api-wrapper/mango_client"
	"testing"
)

func TestGetBalance(t *testing.T) {
	InitMangoCallHandle("", "", "", "")
	balance := mango_client.GetBalance()
	println(balance.Balance)
}
