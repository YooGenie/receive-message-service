package service

import (
	"receive-message-service/config"
	"receive-message-service/dto"

	log "github.com/sirupsen/logrus"
)

func ForConsumerService(message dto.ReceivedEventMessage) {
	log.Println("ForConsumerService")
	var err error
	var code int64

	code, err = MessageService{}.ProcessMessage(message)

	if err != nil {
		if code >= 400 && code < 500 {
			// 400번대 에러
		}
		if code >= 500 && code < 600 {
			// 500번대 에러
		}

		// 시도 끝에 삭제
		err = DeleteMessage(&config.Config.AwsSqs.QueueUrl, &message.MessageHandle)
		if err != nil {
			log.Println("Got an error deleting the message:", err)
		}
		// 에러 로그 DB 저장
		// 슬랙 호출
	} else {
		// 성공한 경우
		err = DeleteMessage(&config.Config.AwsSqs.QueueUrl, &message.MessageHandle)
		if err != nil {
			log.Println("Got an error deleting the message:", err)
		}
	}
}
