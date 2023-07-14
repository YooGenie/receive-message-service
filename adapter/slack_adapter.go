package adapter

import (
	"fmt"
	"receive-message-service/config"

	"github.com/slack-go/slack"
)

func SendMessage() error {
	api := slack.New(config.Config.Slack.Token) //토큰

	channelID, timestamp, err := api.PostMessage(
		config.Config.Slack.ChannelID,
		slack.MsgOptionText("여기가 메시지 구나", false),
	)

	if err != nil {
		fmt.Println("에러 메시지 : ", err)
		return err
	}

	fmt.Printf("slack message post successfully %s at %s", channelID, timestamp)

	return nil
}
