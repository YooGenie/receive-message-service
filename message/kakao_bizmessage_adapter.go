package message

import (
	"log"
)

// 카카오 메시지 보내는 함수
func SendMessage(requestBody map[string]interface{}, memberId int64, donationId int64) (code int64, err error) {
	log.Println("Start")
	// 성공
	code = 200

	return
}
