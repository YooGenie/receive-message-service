package dto

type ReceivedEventMessage struct {
	MessageHandle string                 `json:"messageHandle"`
	Env           string                 `json:"env"`
	ServiceName   string                 `json:"serviceName"`
	EventType     string                 `json:"eventType"`
	Payload       map[string]interface{} `json:"payload"`
	Test          bool
}
