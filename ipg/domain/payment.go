package domain

import (
	"time"
)

type PaymentLinkRequest struct {
	MerchantId          string    `json:"MerchantId"`
	OrderId             int64     `json:"OrderId"`
	UserType            int       `json:"UserType"`
	TerminalId          string    `json:"TerminalId"`
	Amount              int64     `json:"Amount"`
	LocalDateTime       time.Time `json:"LocalDateTime"`
	ReturnUrl           string    `json:"ReturnUrl"`
	SignData            string    `json:"SignData"`
	UserId              string    `json:"UserId"`
	PanAuthenticateType int8      `json:"PanAuthenticateType"`
	NationalCode        string    `json:"NationalCode"`
	ApplicationName     string    `json:"ApplicationName"`
	CardHolderIdentity  string    `json:"CardHolderIdentity"` // cart owner's mobile
	NationalCodeEnc     string    `json:"NationalCodeEnc"`    // encrypted national code
	SourcePanList       []uint64  `json:"SourcePanList"`
}

type PaymentLinkResp struct {
	ResCode     string `json:"ResCode"`
	ResCodeStr  string `json:"-"`
	Token       string `json:"Token"`
	Description string `json:"Description"`
	PaymentUrl  string `json:"-"`
}

func (p PaymentLinkResp) GetPaymentUrl() string {
	return p.PaymentUrl
}

func (p PaymentLinkResp) HasErrors() bool {
	return p.ResCode != "0"
}

func (p PaymentLinkResp) GetError() string {
	return p.ResCodeStr
}
