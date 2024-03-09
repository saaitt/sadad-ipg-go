package domain

import (
	"fmt"
	"strings"
	"time"
)

type RefundTokenRequest struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	GrantType string `json:"grant_type"`
}

type RefundTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	//ExpiresIn   string `json:"expires_in"`
}

type RefundRequest struct {
	TerminalID       string `json:"TerminalID"`
	ClientRequestID  string `json:"ClientRequestID"`
	RNN              string `json:"Rnn"`
	TxnDateTime      string `json:"TxnDateTime"`
	TxnAmount        int64  `json:"TxnAmount"`
	RefundAmount     int64  `json:"RefundAmount"`
	TransferMethodID string `json:"TransferMethodID"`
	WagePayMethodID  string `json:"WagePayMethodID"`
	Description      string `json:"Description"`
	Sign             string `json:"Sign"`
}

func (r *RefundRequest) MakeSignBase(username string) string {
	return fmt.Sprintf("%v,%v,%v,%v,%v", strings.ToLower(username), r.TerminalID, r.RNN, r.TxnAmount, r.RefundAmount)
}

type RefundResponse struct {
	ResCode float64            `json:"ResCode"`
	ResMsg  string             `json:"ResMsg"`
	Data    RefundResponseData `json:"Data"`
}

type RefundResponseData struct {
	RefundID               float64 `json:"RefundID"`
	MoneyTransferMethodID  string  `json:"MoneyTransferMethodID"`
	WageAmount             string  `json:"WageAmount"`
	WagePercent            string  `json:"WagePercent"`
	CardToIbanWageAmount   string  `json:"CardToIbanWageAmount"`
	CardNoMask             string  `json:"CardNoMasked"`
	PayableAmountToEndUser string  `json:"PayableAmountToEndUser"`
	TerminalID             string  `json:"TerminalID"`
	ClientRequestID        string  `json:"ClientRequestID"`
	RNN                    string  `json:"RNN"`
	TxnAmount              string  `json:"TxnAmount"`
	RefundAmount           string  `json:"RefundAmount"`
	TransferMethodID       string  `json:"TransferMethodID"`
	WagePayMethodID        string  `json:"WagePayMethodID"`
	Description            string  `json:"Description"`
}

type InquiryResponse struct {
	ResCode float64             `json:"ResCode"`
	ResMsg  string              `json:"ResMsg"`
	Data    InquiryResponseData `json:"Data"`
}

type InquiryResponseData struct {
	RefundID               float64 `json:"RefundID"`
	WageAmount             string  `json:"WageAmount"`
	CardToIbanWageAmount   string  `json:"CardToIbanWageAmount"`
	CardNoMasked           string  `json:"CardNoMasked"`
	PayableAmountToEndUser string  `json:"PayableAmountToEndUser"`
	ParentID               string  `json:"ParentID"`
	SourceAccountNo        string  `json:"SourceAccountNo"`
	DestAccountNo          string  `json:"DestAccountNo"`
	TxnAmount              string  `json:"TxnAmount"`
	RefundAmount           string  `json:"RefundAmount"`
	TransferMethod         string  `json:"TransferMethod"`
	TransferMethodID       string  `json:"TransferMethodID"`
	WagePayMethodID        string  `json:"WagePayMethodID"`
	Msg                    string  `json:"Msg"`
	TxnDateTime            string  `json:"TxnDateTime"`
	MoneyTransferMethod    string  `json:"MoneyTransferMethod"`
	Description            string  `json:"Description"`
}

type SadadRefundConfig struct {
	Timeout    time.Duration
	BaseUrl    string
	MerchantId string
	TerminalId string
	Username   string
	Password   string
	PrKey      string
}
