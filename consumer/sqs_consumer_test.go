package consumer

import (
	"encoding/json"
	"fmt"
	"log"
	"receive-message-service/dto"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

const TestQueueURL = ""

func TestSqsConsumer_Consume_메시지_수신_1회(t *testing.T) {
	//give
	c := make(chan dto.ReceivedEventMessage)
	consumer := SqsConsumer{
		QueueUrl:            TestQueueURL,
		MaxNumberOfMessages: 1,
		VisibilityTimeout:   time.Second * 2,
		MessageAutoDeleted:  true,
	}

	message := map[string]interface{}{
		"donationId": float64(293858),
		"memberId":   float64(1564),
		"mobile":     "01000000000",
	}

	eventType := "unit-test-group-id"

	err := SendTestMessage(eventType, message)

	if err != nil {
		t.Fail()
	}

	////when
	go consumer.Consume(c)

	//then
	result := <-c //메시지 수신 received message
	assert.Equal(t, "unit-test-group-id", result.EventType)
}

func SendTestMessage(eventType string, message map[string]interface{}) error {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := sqs.New(sess)

	queueURL := TestQueueURL
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
				StringValue: aws.String("notification-service"),
			},
		},
		MessageBody:            aws.String(string(messageBody)),
		QueueUrl:               &queueURL,
		MessageGroupId:         &messageGroupId,
		MessageDeduplicationId: &messageDeduplicationId,
	})

	if err != nil {
		return err
	}

	return nil
}

func TestGetMessages(t *testing.T) {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := sqs.New(sess)

	consumer := SqsConsumer{
		QueueUrl:            TestQueueURL,
		MaxNumberOfMessages: 10,
		VisibilityTimeout:   time.Second * 1,
		//MessageAutoDeleted:  true,
	}

	msgResult, err := svc.ReceiveMessage(&sqs.ReceiveMessageInput{
		AttributeNames: []*string{
			aws.String(sqs.MessageSystemAttributeNameSentTimestamp),
		},
		MessageAttributeNames: []*string{
			aws.String(sqs.QueueAttributeNameAll),
		},
		QueueUrl:            &consumer.QueueUrl,
		MaxNumberOfMessages: aws.Int64(int64(consumer.MaxNumberOfMessages)),
		VisibilityTimeout:   aws.Int64(int64(consumer.VisibilityTimeout.Seconds())),
	})

	if err != nil {
		log.Println(err)
	}

	fmt.Println("msgResult :msgResult ", msgResult)

}
