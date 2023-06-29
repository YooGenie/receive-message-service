package consumer

import (
	"encoding/json"
	"receive-message-service/dto"
	"receive-message-service/service"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/sirupsen/logrus"

	"time"
)

type SqsConsumer struct {
	QueueUrl            string
	MaxNumberOfMessages uint
	VisibilityTimeout   time.Duration
	MessageAutoDeleted  bool
}

func (s SqsConsumer) Consume(c chan dto.ReceivedEventMessage) {

	sess := session.Must(session.NewSessionWithOptions(
		session.Options{
			SharedConfigState: session.SharedConfigEnable,
		}))

	timeout := int64(10)
	for {
		messages, err := service.GetMessages(sess, &s.QueueUrl, &timeout)
		if err != nil {
			logrus.Println("Got an error receiving messages:", err)
		}

		if len(messages.Messages) != 0 {
			for i := 0; i < len(messages.Messages); i++ {
				messagesBody := *messages.Messages[i].Body
				messageHandle := *messages.Messages[i].ReceiptHandle
				if messageHandle == "" {
					logrus.Println("You must supply message receipt handle (-m MESSAGE-HANDLE)")
				}

				eventType := *messages.Messages[i].MessageAttributes["EventType"].StringValue
				env := *messages.Messages[i].MessageAttributes["Env"].StringValue
				module := *messages.Messages[i].MessageAttributes["Module"].StringValue

				eventMessage := make(map[string]interface{}, 1)
				err = json.Unmarshal([]byte(messagesBody), &eventMessage)
				if err != nil {
					logrus.Println("messagesBody Unmarshal error : ", err)
				}

				ReceivedEventMessage := dto.ReceivedEventMessage{
					MessageHandle: messageHandle,
					EventType:     eventType,
					Env:           env,
					Module:        module,
					Payload:       eventMessage,
				}
				err = service.DeleteMessage(sess, &s.QueueUrl, &messageHandle)
				if err != nil {
					logrus.Println("Got an error deleting the message:")
				}
				c <- ReceivedEventMessage
			}
		} else {
			continue
		}
	}

}
