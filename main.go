package main

import (
	"fmt"
	"log"
	"receive-message-service/service"

	"github.com/aws/aws-sdk-go/aws/session"
)

func main() {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	timeout := int64(5)
	queueURL := "https://sqs.ap-northeast-2.amazonaws.com/400281609678/Queue_Test.fifo"

	msgResult, err := service.GetMessages(sess, &queueURL, &timeout)
	if err != nil {
		log.Println("Got an error receiving messages:")
	}

	if len(msgResult.Messages) != 0 {
		fmt.Println("Message ID:     " + *msgResult.Messages[0].MessageId)
		fmt.Println("바디값 : ", *msgResult.Messages[0].Body)
		messageHandle := *msgResult.Messages[0].ReceiptHandle
		if messageHandle == "" {
			log.Println("You must supply message receipt handle (-m MESSAGE-HANDLE)")
		}

	} else {
		log.Println("서비스2 -큐에 아무것도 없음!!")
	}
}
