package dto

type ReceivedEventMessage struct {
	MessageHandle string                 `json:"messageHandle"`
	Env           string                 `json:"env"`
	Module        string                 `json:"module"`
	EventType     string                 `json:"eventType"`
	Payload       map[string]interface{} `json:"payload"`
}
