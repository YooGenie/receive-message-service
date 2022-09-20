package main

import (
	"flag"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"receive-message-service/service"

	"fmt"
	"os"
)

func main() {
	// Initialize a session that the SDK will use to load
	// credentials from the shared credentials file. (~/.aws/credentials).
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := sns.New(sess)

	sqsPtr := "" //ARN주소
	topicPtr := "" //ARN주소

	if sqsPtr == "" || topicPtr == "" {
		fmt.Println("You must supply an email address and topic ARN")
		fmt.Println("Usage: go run SnsSubscribe.go -e EMAIL -t TOPIC-ARN")
		os.Exit(1)
	}

	//SNS 구독하는 방법
	result, err := svc.Subscribe(&sns.SubscribeInput{
		Endpoint:              &sqsPtr,
		Protocol:              aws.String("sqs"),
		ReturnSubscriptionArn: aws.Bool(true), // Return the ARN, even if user has yet to confirm
		TopicArn:              &topicPtr,
	})
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println(*result)

	timeout := flag.Int64("t", 5, "How long, in seconds, that the message is hidden from others")
	queue := flag.String("q", "", "The name of the queue")
	flag.Parse()

	//큐 주소 조회하기
	queueURL, err := service.GetQueueURL(sess, queue)
	if err != nil {
		fmt.Println("Got an error getting the queue URL:")
		return
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

	msgResult, err := service.GetMessages(sess, queueURL.QueueUrl, timeout)
	if err != nil {
		fmt.Println("Got an error receiving messages:")
		fmt.Println(err)
		return
	}

	fmt.Println("Message ID:     " + *msgResult.Messages[0].MessageId)
	fmt.Println("Message Handle: " + *msgResult.Messages[0].ReceiptHandle)

	messageHandle := flag.String("m", *msgResult.Messages[0].ReceiptHandle, "The receipt handle of the message")
	flag.Parse()

	if *messageHandle == "" {
		fmt.Println("You must supply message receipt handle (-m MESSAGE-HANDLE)")
		return
	}

	err = service.DeleteMessage(sess, queueURL.QueueUrl, messageHandle)
	if err != nil {
		fmt.Println("Got an error deleting the message:")
		fmt.Println(err)
		return
	}

	fmt.Println("Deleted message from queue with URL " + *queueURL.QueueUrl)
}
