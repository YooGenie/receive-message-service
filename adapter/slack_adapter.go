package adapter

import (
	"fmt"
	"os"
	"receive-message-service/config"

	"github.com/slack-go/slack"
)

func SendMessageToSlack() error {
	api := slack.New(config.Config.Slack) //토큰

	channelID, timestamp, err := api.PostMessage(
		os.Getenv("CHANNEL_ID"),
		slack.MsgOptionText("여기가 메시지 구나", false),
	)

	if err != nil {
		fmt.Println("에러 메시지 : ", err)
		return err
	}

	fmt.Printf("slack message post successfully %s at %s", channelID, timestamp)
	return nil
}
