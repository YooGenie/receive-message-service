package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/aws/aws-sdk-go/aws/session"
	"log"
	"receive-message-service/dto"
	"receive-message-service/kakao"
	"receive-message-service/service"
	"time"
)

func main() {
	//Config.Bizmessage.Account.Password = properties["MESSAGING_ACCOUNT_PASSWORD"]
	// Initialize a session that the SDK will use to load
	// credentials from the shared credentials file. (~/.aws/credentials).
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	queue := flag.String("q", "ysun-test.fifo", "The name of the queue")
	flag.Parse()
	if *queue == "" {
		log.Println("You must supply the name of a queue (-q QUEUE)")
	}

	//큐 주소 조회하기
	queueURL, err := service.GetQueueURL(sess, queue)
	if err != nil {
		log.Println("Got an error getting the queue URL:", queueURL)
	}

	fmt.Println(queueURL)

	timeout := flag.Int64("tag", 5, "How long, in seconds, that the message is hidden from others")

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

			messageBody := *msgResult.Messages[0].Body
			//log.Println("Message Body11: " , messageBody	)


			fmt.Println(messageBody)
			data :=dto.DonationInfo{}
			err := json.Unmarshal([]byte(messageBody), &data)
			if err != nil {
				log.Println(err)
			}



			kakao.KakaoBizmessageAdapter().SendDonationReservedMessage(data.MemberId,data.Mobile, data.ReservationNo, data.CampaignName,data.NickName,data.PostPlace)
			messageHandle := *msgResult.Messages[0].ReceiptHandle
			flag.Parse()

			if messageHandle == "" {
				log.Println("You must supply message receipt handle (-m MESSAGE-HANDLE)")
			}

			//err = service.DeleteMessage(sess, queueURL.QueueUrl, &messageHandle)
			//if err != nil {
			//	log.Println("Got an error deleting the message:")
			//}
			//log.Println("Deleted message from queue with URL " + *queueURL.QueueUrl)
		} else {
			log.Println("서비스2 -큐에 아무것도 없음!!")
		}
	}

}
