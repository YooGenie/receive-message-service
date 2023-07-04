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
	config.InitConfig("../config/config.test.json")
	config.ConfigureEnvironment("../config")
}

func TestGetMessages(t *testing.T) {
	t.Run("(성공) 큐에 메시지 없을 때", func(t *testing.T) {
		sess := session.Must(session.NewSessionWithOptions(session.Options{
			SharedConfigState: session.SharedConfigEnable,
		}))

		queueURL := config.Config.AwsSqs.QueueUrl

		timeout := int64(10)
		msgResult, err := GetMessages(sess, &queueURL, &timeout)

		assert.Nil(t, msgResult.Messages)
		assert.Nil(t, err)

	})
	t.Run("(성공) 큐에 메시지 있을 때", func(t *testing.T) {
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
		msgResult, err := GetMessages(sess, &queueURL, &timeout)

		assert.Equal(t, "test", *msgResult.Messages[0].MessageAttributes["Env"].StringValue)
		assert.NotNil(t, msgResult.Messages)
		assert.Nil(t, err)

	})
}
