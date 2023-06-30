package service

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func DeleteMessage(queueURL *string, messageHandle *string) error {
	sess := session.Must(session.NewSessionWithOptions(
		session.Options{
			SharedConfigState: session.SharedConfigEnable,
		}))

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
