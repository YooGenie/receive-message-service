package main

import (
	"fmt"
	"receive-message-service/consumer"
	"receive-message-service/dto"
	"time"
)

func main() {
	queueURL := ""

	megCh := make(chan dto.ReceivedEventMessage)

	sqsConsumer := consumer.SqsConsumer{
		QueueUrl:            queueURL,
		MaxNumberOfMessages: 10,
		VisibilityTimeout:   time.Second * 30,
	}

	go sqsConsumer.Consume(megCh)
	for {
		select {
		case message := <-megCh:
			fmt.Println("메시지 내용 : ", message)
		}
	}
}
