package adapter

import (
	"fmt"
	"os"
	"receive-message-service/config"

	"github.com/slack-go/slack"
)

func SendMessageToSlack() error {
	s := "[utils/SendMessageToSlack]"
	api := slack.New(config.Config.Slack)

	channelID, timestamp, err := api.PostMessage(
		os.Getenv("CHANNEL_ID"),
		slack.MsgOptionText("alert! you must fix it!", false),
	)

	if err != nil {
		fmt.Printf("%s %v\n", s, err)
		return err
	}

	fmt.Printf("slack message post successfully %s at %s", channelID, timestamp)
	return nil
}
