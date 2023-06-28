package main

import (
	"fmt"
	"log"
	"receive-message-service/service"
	"time"

	"github.com/aws/aws-sdk-go/aws/session"
)

func main() {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	timeout := int64(10)
	queueURL := ""

	for {
		time.Sleep(time.Second * 5)

		msgResult, err := service.GetMessages(sess, &queueURL, &timeout)
		if err != nil {
			log.Println("Got an error receiving messages:")
		}

		if len(msgResult.Messages) != 0 {
			for i, _ := range msgResult.Messages {
				fmt.Println(i, "번째 메시지 바디내용", *msgResult.Messages[i].Body)
				messageHandle := *msgResult.Messages[i].ReceiptHandle
				if messageHandle == "" {
					log.Println("You must supply message receipt handle (-m MESSAGE-HANDLE)")
				}
				err = service.DeleteMessage(sess, &queueURL, &messageHandle)
				if err != nil {
					log.Println("Got an error deleting the message:", err)
				}
				log.Println("Deleted message from queue with URL " + queueURL)
			}
		} else {
			log.Println("서비스2 -큐에 아무것도 없음!!")
		}
	}
}
