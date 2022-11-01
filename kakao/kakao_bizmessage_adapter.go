package kakao

import (
	"fmt"
	"receive-message-service/config"
	"strings"
	"sync"
	"time"

	"github.com/bettercode-oss/rest"
)

const KoreaCountryCode = "82"

var (
	kakaoBizmessageAdapterOnce     sync.Once
	kakaoBizmessageAdapterInstance *kakaoBizmessageAdapter
)

func KakaoBizmessageAdapter() *kakaoBizmessageAdapter {
	kakaoBizmessageAdapterOnce.Do(func() {
		kakaoBizmessageAdapterInstance = &kakaoBizmessageAdapter{}
	})

	return kakaoBizmessageAdapterInstance
}

type kakaoBizmessageAdapter struct {
}

func (adapter kakaoBizmessageAdapter) getApiToken() (token map[string]interface{}, err error) {
	client := rest.Client{
		ShowHttpLog: config.Config.Log.ShowHttpLog,
		RetryMax:    5,
		RetryDelay:  1 * time.Second,
	}

	err = client.
		Request().
		SetHeader("Accept", "application/json").
		SetHeader("X-IB-Client-Id","beautiful_talk").
		SetHeader("X-IB-Client-Passwd", "qfbpBuxqZRx4VOsvg8dw").
		SetResult(&token).
		Post("https://msggw.supersms.co:9440/auth/v1/token")
	if err != nil {
		return
	}

	return token, nil
}

func (adapter kakaoBizmessageAdapter) sendMessage(requestBody map[string]interface{}) (err error) {
	token, err := adapter.getApiToken()
	//authorization := "bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJGbTBXbEtNM0FsNW1uVkJhSURqT2VSODJtMGVhTm9ZOU1uQmY2UEhvb2xFIn0.eyJleHAiOjE2NjczNjQzNTYsImlhdCI6MTY2NzI3Nzk1NiwianRpIjoiYjk1MTE1MzMtOWMzZC00ZjM0LWFhOGUtMzcxOTJkZWMxMmQ3IiwidHlwIjoiQmVhcmVyIiwicmVzb3VyY2VfYWNjZXNzIjp7Im1zZ2d3Ijp7InJvbGVzIjpbIkFMIiwiVVNFUiIsIkZUIl19fSwidXNlcl9uYW1lIjoiYmVhdXRpZnVsX3RhbGsifQ.hWJ338nXObJVJhdaq9xvX_o8DWu3Uj0XnZ4arhC0RYiB0SAN9hE7hnhRFOMkKN_BUTwtkL90s2UECJ7ZFqMYDfURzNvBN9i5UzptPzH1Lnz8j9kaCLdpnjqmk5J-YF5OgSB9hChtlUga1FadlYJDGHFrMcIruVeJaLVmi83MdQHgilVo1scqqRzMGeUUdqNE74OZLlX5_akSxw3-K0bVjZcAxMUj5iASLs0yjrz07g2a8ogmjYi_zeqF5HZygGZWNrdSHrntUA4YyN4AQ8F6zaf9R5Yks1XzAxqylYeBLfBFuFQj8bnd3o3UcEJnlPuyZAS5p2jyONqw8DT3ahZzAA"

	authorization := fmt.Sprintf("%s %s", token["schema"].(string), token["accessToken"].(string))
	client := rest.Client{
		ShowHttpLog: config.Config.Log.ShowHttpLog,
		RetryMax:    5,
		RetryDelay:  1 * time.Second,
	}

	url := "https://msggw.supersms.co:9443/v1/send/kko"
	response := map[string]interface{}{}

	err = client.
		Request().
		SetHeader("Accept", "application/json").
		SetHeader("X-IB-Client-Id","beautiful_talk").
		SetHeader("X-IB-Client-Passwd", "qfbpBuxqZRx4VOsvg8dw").
		SetHeader("Authorization", authorization).
		SetBody(requestBody).
		SetResult(&response).
		Post(url)

	if err != nil {
		 err.Error()
	}

	return err
}

func (kakaoBizmessageAdapter) convertKakaoSendingNo(cellPhone string) string {
	if strings.HasPrefix(cellPhone, "0") {
		return fmt.Sprintf("%s%s", KoreaCountryCode, cellPhone[1:])
	} else {
		return fmt.Sprintf("%s%s", KoreaCountryCode, cellPhone)
	}
}

// 택배신청-멤버
func (adapter kakaoBizmessageAdapter) SendDonationReservedMessage(memberId int64, cellPhoneNumber string, reservedNo string, campaignName string, nickName string, postPlace string) error {

	domain := "https://localhost:3000"
	if nickName == "" {
		nickName = "기부자"
	}

	var templateCode, contents, title, utm string
	switch postPlace {
	case "CU":
		courierText := postPlace + " 승인번호"
		formattedReservedNO := fmt.Sprintf("%s-%s-%s", reservedNo[0:4], reservedNo[4:8], reservedNo[8:])
		title = formattedReservedNO
		templateCode = "SI.0002-4" // CU
		contents = fmt.Sprintf("%s \n<택배기부 신청 안내>\n\n%s님, 물품 기부를 신청해 주셔서 감사합니다.\n\n▶%s: %s\n\n▶CU 포스트박스 이용방법\n편의점 포스트박스 화면에서 '쇼핑몰 거래(사전예약/선결제)' 선택하신 후, 위 승인번호를 입력해 주세요.\n\n*유의사항*\n물품 포장 박스의 무게가 5kg를 초과하면 배송이 불가하오니 유의해 주세요.\n\n%s님의 참여가 나눔과 순환으로 연결되어 세상의 생명을 연장합니다.", campaignName, nickName, courierText, formattedReservedNO, nickName)
		utm = "utm_source=kakao&utm_medium=kakaoalarm&utm_campaign=%EB%AC%BC%ED%92%88%EA%B8%B0%EB%B6%8022&utm_content=GS%EC%8B%A0%EC%B2%AD%EC%99%84%EB%A3%8C"

}

	requestBody := map[string]interface{}{
		"msg_type":    "AL",
		"mt_failover": "N",
		"msg_data": map[string]interface{}{
			"to":      adapter.convertKakaoSendingNo(cellPhoneNumber),
			"content": contents,
			"title":   fmt.Sprintf("%s", title),
		},
		"msg_attr": map[string]interface{}{
			"sender_key":     "57f25d13243c7cc799dca491da72b97620f56c7f",
			"template_code":   templateCode,
			"response_method": "push",
			"content_type":    "V",
			"attachment": map[string]interface{}{
				"button": []interface{}{
					map[string]interface{}{
						"name":       "기부 신청내역 확인",
						"type":       "WL",
						"url_pc":     fmt.Sprintf("%s/mypage/donation?%s", domain, utm),
						"url_mobile": fmt.Sprintf("%s/mypage/donation?%s", domain, utm),
					},
				},
			},
		},
	}

	return adapter.sendMessage(requestBody)
}

