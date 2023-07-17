package adapter

import (
	"fmt"
	"receive-message-service/config"
	"receive-message-service/dto"

	"github.com/slack-go/slack"
)

type Content struct {
	Message     dto.ReceivedEventMessage
	Error       error
	Code        int64
	Env         string
	ServiceName string
	Content     string
}

func SendMessage(content Content) error {

	contents := fmt.Sprintf("# 기부 정보\n"+
		"* 환경 : %v \n"+
		"* 메시지 타입: %v\n"+
		"* 에러 메시지: %v\n"+
		"* 보낸 서비스 : %v\n"+
		"* ErrorCode : %d\n"+
		"* Content : %v\n"+
		"* Payload : %v\n",
		content.Env,
		content.Message.EventType,
		content.Error, content.ServiceName, content.Code, content.Content, content.Message.Payload)

	api := slack.New(config.Config.Slack.Token) //토큰

	channelID, timestamp, err := api.PostMessage(
		config.Config.Slack.ChannelID,
		slack.MsgOptionText(contents, false),
	)

	if err != nil {
		fmt.Println("에러 메시지 : ", err)
		return err
	}

	fmt.Printf("slack message post successfully %s at %s", channelID, timestamp)

	return nil
}
