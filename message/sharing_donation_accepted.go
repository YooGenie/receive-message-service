package message

import (
	"github.com/sirupsen/logrus"
)

type Content struct {
	Mobile     string
	DonationId int64
	MemberId   int64
}

func SendDonationRegisteredMessage(content Content) (code int64, err error) {
	// 물품등록-매장관리자
	logrus.Trace("SendDonationCancelledMessage")

	// 보통 카카오 메시지 템플릿 작성
	requestBody := map[string]interface{}{
		"memberId": content.MemberId,
	}

	return SendMessage(requestBody, content.MemberId, content.DonationId)

}
