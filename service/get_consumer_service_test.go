package service

import (
	"receive-message-service/config"
	"receive-message-service/dto"
	"testing"
)

func init() {
	config.InitConfig("../config/config.test.json")
	config.ConfigureEnvironment("../config")
}

func TestForConsumerService_ok(t *testing.T) {
	message := map[string]interface{}{
		"donationId": float64(293858),
		"memberId":   float64(1564),
		"mobile":     "01000000000",
	}

	eventType := "unit-test-group-id"
	ReceivedEventMessage := dto.ReceivedEventMessage{
		MessageHandle: "",
		EventType:     eventType,
		Env:           config.Config.Environment,
		ServiceName:   config.Config.ServiceName,
		Payload:       message,
		Test:          true,
	}

	ForConsumerService(ReceivedEventMessage)
}
