package refund

import (
	"github.com/saaitt/sadad-ipg-go/refund/domain"
	"github.com/saaitt/sadad-ipg-go/refund/usecase"
)

func NewRefundClient(refundConfig *domain.SadadRefundConfig) domain.SadadRefundClient {
	return usecase.Init(refundConfig)
}
