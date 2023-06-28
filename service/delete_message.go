package service

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func DeleteMessage(sess *session.Session, queueURL *string, messageHandle *string) error {
	svc := sqs.New(sess)

	fmt.Println("messageHandle : ", queueURL)

	_, err := svc.DeleteMessage(&sqs.DeleteMessageInput{
		QueueUrl:      queueURL,
		ReceiptHandle: messageHandle,
	})

	if err != nil {
		return err
	}

	return nil
}
