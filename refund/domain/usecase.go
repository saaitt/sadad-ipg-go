package domain

type SadadRefundClient interface {
	NewRefund(refReq *RefundRequest) (*RefundResponse, error)
}
