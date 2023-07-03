package service

import (
	"encoding/json"
	"receive-message-service/config"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func init() {
	config.InitConfig("../config/config_test.json")
	config.ConfigureEnvironment("../config")
}

func TestDeleteMessage(t *testing.T) {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := sqs.New(sess)

	message := map[string]interface{}{
		"donationId": float64(293858),
		"memberId":   float64(1564),
		"mobile":     "01000000000",
	}

	eventType := "unit-test-group-id"

	queueURL := config.Config.AwsSqs.QueueUrl
	messageGroupId := eventType
	messageDeduplicationId := uuid.New().String()

	messageBody, _ := json.Marshal(message)

	var _, err = svc.SendMessage(&sqs.SendMessageInput{
		MessageAttributes: map[string]*sqs.MessageAttributeValue{
			"EventType": {
				DataType:    aws.String("String"),
				StringValue: aws.String(eventType),
			},
			"Env": {
				DataType:    aws.String("String"),
				StringValue: aws.String("test"),
			},
			"Module": {
				DataType:    aws.String("String"),
				StringValue: aws.String("receive-message-service"),
			},
		},
		MessageBody:            aws.String(string(messageBody)),
		QueueUrl:               &queueURL,
		MessageGroupId:         &messageGroupId,
		MessageDeduplicationId: &messageDeduplicationId,
	})

	if err != nil {
		t.Fail()
	}
	timeout := int64(10)
	messages, err := GetMessages(sess, &queueURL, &timeout)
	if err != nil {
		return
	}

	messageHandle := *messages.Messages[0].ReceiptHandle
	err = DeleteMessage(&queueURL, &messageHandle)

	assert.NotNil(t, messageHandle)
	assert.Nil(t, err)
}
