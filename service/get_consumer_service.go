package service

import (
	"errors"
	"receive-message-service/adapter"
	"receive-message-service/config"
	"receive-message-service/dto"

	log "github.com/sirupsen/logrus"
)

func ForConsumerService(message dto.ReceivedEventMessage) {
	log.Println("ForConsumerService")
	var err error
	var code int64

	config.Config.Environment = message.Env

	if !message.Test {
		code, err = MessageService{}.ProcessMessage(message)
	}
	if message.Test {
		err = errors.New("에러 메시지 입니다.")
		code = 400
	}

	var content adapter.Content
	content.Message = message
	content.ServiceName = message.ServiceName
	content.Error = err
	content.Env = message.Env
	content.Code = code

	if err != nil {

		if code >= 400 && code < 500 {
			// 400트번대 에러
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
		err = adapter.SendMessage(content)
		if err != nil {
			log.Println(err)
		}

	} else {
		// 성공한 경우
		err = DeleteMessage(&config.Config.AwsSqs.QueueUrl, &message.MessageHandle)
		if err != nil {
			log.Println("Got an error deleting the message:", err)
		}
	}
}
