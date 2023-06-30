package service

import (
	"log"
	"receive-message-service/dto"
	"receive-message-service/message"

	"github.com/sirupsen/logrus"
)

type MessageService struct {
}

func (s MessageService) ProcessMessage(messageContent dto.ReceivedEventMessage) (code int64, err error) {
	log.Println("MessageService - ProcessMessage in")

	if messageContent.EventType == "donationRegistered" || messageContent.EventType == "unit-test-group-id" {
		logrus.Println("ProcessMessage - donationRegistered", int64(messageContent.Payload["donationId"].(float64)))
		content := message.Content{
			Mobile:     messageContent.Payload["mobile"].(string),
			DonationId: int64(messageContent.Payload["donationId"].(float64)),
			MemberId:   int64(messageContent.Payload["memberId"].(float64)),
		}

		code, err = message.SendDonationRegisteredMessage(content)

		return code, err
	}
	return code, err
}
