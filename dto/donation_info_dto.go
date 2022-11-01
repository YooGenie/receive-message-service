package dto

type DonationInfo struct {
	DonationId    int64  `json:"donationId" `
	MemberId      int64 `json:"memberId" `
	Mobile        string `json:"mobile" `
	ReservationNo string `json:"reservationNo" `
	CampaignName  string `json:"campaignName" `
	NickName      string `json:"nickName" `
	PostPlace     string `json:"postPlace" `
}
