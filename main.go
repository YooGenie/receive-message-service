package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws/session"
	"receive-message-service/queue"
	"sync"
)

func main() {
	// Initialize a session that the SDK will use to load
	// credentials from the shared credentials file. (~/.aws/credentials).
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		queue.TestYsh(sess)
		wg.Done()
	}()
	go func() {
		queue.TestYsun(sess)
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("끝")

}
