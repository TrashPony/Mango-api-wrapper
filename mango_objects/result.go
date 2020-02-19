package mango_objects

type Result struct {
	CommandID string `json:"command_id"`
	Result    int    `json:"result"`
	Name      string `json:"name"`
	Message   string `json:"message"`
}
