package service

import (
	"github.com/aws/aws-sdk-go/aws/session"
)

func DeleteMessage(sess *session.Session, queueURL *string, messageHandle *string) error {
	//svc := sqs.New(sess)
	//
	//_, err := svc.DeleteMessage(&sqs.DeleteMessageInput{
	//	QueueUrl:      queueURL,
	//	ReceiptHandle: messageHandle,
	//})
	//
	//if err != nil {
	//	return err
	//}

	return nil
}
