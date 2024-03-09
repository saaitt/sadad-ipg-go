package domain

import (
	"time"
)

type ValidateReq struct {
	Token    string `json:"Token"`
	SignData string `json:"SignData"`
}

type VerifyPayment struct {
	ResCode            string `json:"ResCode"`
	ResCodeStr         string `json:"-"`
	Amount             string `json:"Amount"`
	Description        string `json:"Description"`
	RetrievalRefNo     string `json:"RetrivalRefNo"`
	SystemTraceNo      string `json:"SystemTraceNo"`
	OrderId            string `json:"OrderId"`
	TransactionDate    string `json:"TransactionDate"`
	CardHolderFullName string `json:"CardHolderFullName"`
}

func (v VerifyPayment) IsSuccess() bool {
	return v.ResCode == VerifyResponseResCodeSuccess
}

func (v VerifyPayment) GetError() string {
	return v.ResCodeStr
}

func (v VerifyPayment) GetPaymentDate() time.Time {
	// TODO: payment date layout ?
	pd, _ := time.Parse("", v.TransactionDate)
	return pd.UTC()
}

func (v VerifyPayment) GetTransactionNumber() string {
	return v.RetrievalRefNo
}

func (v VerifyPayment) GetTraceNumber() string {
	return v.SystemTraceNo
}

func (v VerifyPayment) GetFactorNumber() string {
	return v.OrderId
}
