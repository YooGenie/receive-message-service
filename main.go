package main

import (
	"fmt"
	"receive-message-service/config"
	"receive-message-service/consumer"
	"receive-message-service/dto"
	"receive-message-service/service"
	"time"
)

func main() {
	config.InitConfig("./config/config.json")
	config.ConfigureEnvironment("./config")

	megCh := make(chan dto.ReceivedEventMessage)

	sqsConsumer := consumer.SqsConsumer{
		QueueUrl:            config.Config.AwsSqs.QueueUrl,
		MaxNumberOfMessages: 10,
		VisibilityTimeout:   time.Second * 30,
	}

	go sqsConsumer.Consume(megCh)
	for {
		select {
		case message := <-megCh:
			// 채널에 있는 값을 메시지에 넣어서
			service.ForConsumerService(message)
			fmt.Println("메시지 내용 : ", message)
		}
	}
}
