package main

import (
	"flag"
	"github.com/aws/aws-sdk-go/aws/session"
	"log"
	"receive-message-service/service"
	"time"

	"fmt"
)

func main() {
	// Initialize a session that the SDK will use to load
	// credentials from the shared credentials file. (~/.aws/credentials).
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	timeout := flag.Int64("t", 5, "How long, in seconds, that the message is hidden from others")
	queue := flag.String("q", "", "The name of the queue")
	flag.Parse()

	//큐 주소 조회하기
	queueURL, err := service.GetQueueURL(sess, queue)
	if err != nil {
		fmt.Println("Got an error getting the queue URL:")
	}

	if *queue == "" {
		fmt.Println("You must supply the name of a queue (-q QUEUE)")
		return
	}

	//메시지 받기
	if *timeout < 0 {
		*timeout = 0
	}

	if *timeout > 12*60*60 {
		*timeout = 12 * 60 * 60
	}


	for  {
		time.Sleep(time.Second*5)
		msgResult, err := service.GetMessages(sess, queueURL.QueueUrl, timeout)
		if err != nil {
			fmt.Println("Got an error receiving messages:")
			fmt.Println(err)
		}

		if len(msgResult.Messages) !=0 {
			fmt.Println("Message ID:     " + *msgResult.Messages[0].MessageId)
			fmt.Println("Message Handle: " + *msgResult.Messages[0].ReceiptHandle)
			fmt.Println("Message Body: " + *msgResult.Messages[0].Body)

			//messageHandle := flag.String("m", *msgResult.Messages[0].ReceiptHandle, "The receipt handle of the message")

			messageHandle:= *msgResult.Messages[0].ReceiptHandle
			flag.Parse()

			if messageHandle == "" {
				fmt.Println("You must supply message receipt handle (-m MESSAGE-HANDLE)")
			}

			err = service.DeleteMessage(sess, queueURL.QueueUrl, &messageHandle)
			if err != nil {
				fmt.Println("Got an error deleting the message:")
				fmt.Println(err)
			}

			fmt.Println("Deleted message from queue with URL " + *queueURL.QueueUrl)
		}else {
			log.Println("큐에 아무것도 없음!!")
		}
	}

}
