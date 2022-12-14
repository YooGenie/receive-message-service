package main

import (
	"flag"
	"fmt"
	"github.com/aws/aws-sdk-go/aws/session"
	"log"
	"receive-message-service/service"
	"time"
)

func main() {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	timeout := flag.Int64("t", 5, "How long, in seconds, that the message is hidden from others")
	queue := flag.String("q", "sqs-test.fifo", "The name of the queue")
	flag.Parse()

	//큐 주소 조회하기
	queueURL, err := service.GetQueueURL(sess, queue)
	if err != nil {
		fmt.Println("Got an error getting the queue URL:")
		return
	}

	fmt.Println(queueURL)

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

	for {
		time.Sleep(time.Second * 5)

		msgResult, err := service.GetMessages(sess, queueURL.QueueUrl, timeout)
		if err != nil {
			log.Println("Got an error receiving messages:")
		}

		if len(msgResult.Messages) != 0 {
			fmt.Println("Message ID:     " + *msgResult.Messages[0].MessageId)
			fmt.Println(*msgResult.Messages[0].Body)
			messageHandle := *msgResult.Messages[0].ReceiptHandle
			if messageHandle == "" {
				log.Println("You must supply message receipt handle (-m MESSAGE-HANDLE)")
			}

			err = service.DeleteMessage(sess, queueURL.QueueUrl, &messageHandle)
			if err != nil {
				log.Println("Got an error deleting the message:")
			}
			log.Println("Deleted message from queue with URL " + *queueURL.QueueUrl)

		} else {
			log.Println("서비스2 -큐에 아무것도 없음!!")
		}
	}
}
