package usecase

import (
	"github.com/saaitt/sadad-ipg-go/ipg/domain"
	"github.com/saaitt/sadad-ipg-go/ipg/usecase"
)

type SadadIpg interface {
	CreatePaymentLink(request *domain.PaymentLinkRequest) (*domain.PaymentLinkResp, error)
	ValidatePayment(data interface{}) (*domain.VerifyPayment, error)
}

func NewIpgClient(config *domain.SadadIpgConfig) SadadIpg {
	return usecase.Init(config)
}
