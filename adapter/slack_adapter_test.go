package adapter

import (
	"receive-message-service/config"
	"testing"
)

func init() {
	config.InitConfig("../config/config.test.json")
	config.ConfigureEnvironment("../config")
}

func TestSendMessage(t *testing.T) {
	t.Run("(성공) 큐에 메시지 있을 때", func(t *testing.T) {
		err := SendMessage()
		if err != nil {
			return
		}
	})
}
