package mango

import (
	"github.com/TrashPony/Mango-api-wrapper/mango_client"
	"github.com/TrashPony/Mango-api-wrapper/mango_commands"
	"github.com/TrashPony/Mango-api-wrapper/mango_objects"
)

type Client struct {
	apiKey  string
	apiSing string
	apiUrl  string
}

func (c *Client) SetApiKey(key string) {
	c.apiKey = key
}

func (c *Client) SetApiSing(sing string) {
	c.apiSing = sing
}

func (c *Client) SetApiUrl(url string) {
	c.apiUrl = url
}

func (c *Client) GetBalance() *mango_client.Balance {
	return mango_client.GetBalance(c.apiKey, c.apiSing, c.apiUrl)
}

func (c *Client) EndCall(callID string) *mango_objects.Result {
	return mango_commands.EndCall(callID, c.apiKey, c.apiSing, c.apiUrl)
}

func (c *Client) InitCallOfGroup(from, to, lineNumber string) *mango_objects.Result {
	return mango_commands.InitCallOfGroup(from, to, lineNumber, c.apiKey, c.apiSing, c.apiUrl)
}

func (c *Client) InitCallOfUser(extension, callerNumber, toNumber, lineNumber, sipHeaders string) *mango_objects.Result {
	return mango_commands.InitCallOfUser(extension, callerNumber, toNumber, lineNumber, sipHeaders, c.apiKey, c.apiSing, c.apiUrl)
}

func (c *Client) RoutingCall(callID, toNumber, sipHeaders string) *mango_objects.Result {
	return mango_commands.RoutingCall(callID, toNumber, sipHeaders, c.apiKey, c.apiSing, c.apiUrl)
}

func (c *Client) SendSms(fromExtension, toNumber, smsSender, text string) *mango_objects.Result {
	return mango_commands.SendSms(fromExtension, toNumber, smsSender, text, c.apiKey, c.apiSing, c.apiUrl)
}

func (c *Client) TransferCall(callID, method, toNumber, initiator string) *mango_objects.Result {
	return mango_commands.TransferCall(callID, method, toNumber, initiator, c.apiKey, c.apiSing, c.apiUrl)
}

func (c *Client) TurnOnRecordingCall(callID, callPartyNumber string) *mango_objects.Result {
	return mango_commands.TurnOnRecordingCall(callID, callPartyNumber, c.apiKey, c.apiSing, c.apiUrl)
}
