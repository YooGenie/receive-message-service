package adapter

import (
	"errors"
	"receive-message-service/config"
	"receive-message-service/dto"
	"testing"
)

func init() {
	config.InitConfig("../config/config.test.json")
	config.ConfigureEnvironment("../config")
}

func TestSendMessage(t *testing.T) {
	t.Run("(성공)", func(t *testing.T) {
		message := map[string]interface{}{
			"donationId": float64(293858),
			"memberId":   float64(1564),
			"mobile":     "01050560712",
		}

		ReceivedEventMessage := dto.ReceivedEventMessage{
			MessageHandle: "",
			EventType:     "unit-test-group-id",
			Env:           config.Config.Environment,
			ServiceName:   config.Config.ServiceName,
			Payload:       message,
		}

		var content = Content{
			Message:     ReceivedEventMessage,
			Error:       errors.New("400번 에러"),
			Code:        400,
			Env:         config.Config.Environment,
			ServiceName: config.Config.ServiceName,
			Content:     "에러입니다",
		}

		err := SendMessage(content)
		if err != nil {
			return
		}
	})
}
