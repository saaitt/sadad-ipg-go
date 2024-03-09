package domain

var PaymentReqResCodes = map[string]string{
	"0":    "PaymentReqResCodeSuccess",
	"3":    "PaymentReqResCodeInvalidMerchant",
	"23":   "PaymentReqResCodeInactiveMerchant",
	"58":   "PaymentReqResCodeTransactionNotPermitted",
	"61":   "PaymentReqResCodeInvalidAmount",
	"1000": "PaymentReqResCodeInvalidParamPriority",
	"1001": "PaymentReqResCodeInvalidParam",
	"1002": "PaymentReqResCodeSystemError",
	"1003": "PaymentReqResCodeInvalidMerchantIP",
	"1004": "PaymentReqResCodeInvalidMerchantID",
	"1006": "PaymentReqResCodeSystemFailure",
	"1011": "PaymentReqResCodeDuplicateRequest",
	"1012": "PaymentReqResCodeInvalidMerchantInfo",
	"1017": "PaymentReqResCodePaymentMoreThanMerchantLimit",
	"1018": "PaymentReqResCodeInvalidDate",
}

var VerifyResCodes = map[string]string{
	"0":   "VerifyResCodeSuccess",
	"100": "VerifyResCodeDuplicate",
	"-1":  "VerifyResCodeInvalidParamOrTransactionNotFound",
	"101": "VerifyResCodeExpired",
}

var VerifyResponseResCodes = map[string]string{
	"0":  "VerifyResponseResCodeSuccess",
	"-1": "VerifyResponseResCodeFailed",
}

const (
	PaymentReqResCodeSuccess                      string = "0"
	PaymentReqResCodeInvalidMerchant              string = "3"
	PaymentReqResCodeInactiveMerchant             string = "23"
	PaymentReqResCodeTransactionNotPermitted      string = "58"
	PaymentReqResCodeInvalidAmount                string = "61"
	PaymentReqResCodeInvalidParamPriority         string = "1000"
	PaymentReqResCodeInvalidParam                 string = "1001"
	PaymentReqResCodeSystemError                  string = "1002"
	PaymentReqResCodeInvalidMerchantIP            string = "1003"
	PaymentReqResCodeInvalidMerchantID            string = "1004"
	PaymentReqResCodeAccessDenied                 string = "1005"
	PaymentReqResCodeSystemFailure                string = "1006"
	PaymentReqResCodeDuplicateRequest             string = "1011"
	PaymentReqResCodeInvalidMerchantInfo          string = "1012"
	PaymentReqResCodePaymentMoreThanMerchantLimit string = "1017"
	PaymentReqResCodeInvalidDate                  string = "1018"
	PaymentReqResCode
)

const (
	VerifyResCodeSuccess                           string = "0"
	VerifyResCodeDuplicate                         string = "100"
	VerifyResCodeInvalidParamOrTransactionNotFound string = "-1"
	VerifyResCodeExpired                           string = "101"
)

const (
	VerifyResponseResCodeSuccess string = "0"
	VerifyResponseResCodeFailed  string = "-1"
)
