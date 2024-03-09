package usecase

import "github.com/saaitt/sadad-ipg-go/ipg/domain"

func (c *client) getVerifyError(resCode string) string {
	c.locker.Lock()
	res, ok := domain.VerifyResCodes[resCode]
	c.locker.Unlock()
	if !ok {
		return ""
	}
	return res
}

func (c *client) getVerifyResponseError(resCode string) string {
	c.locker.Lock()
	res, ok := domain.VerifyResponseResCodes[resCode]
	c.locker.Unlock()
	if !ok {
		return ""
	}
	return res

}

func (c *client) getPaymentRequestError(resCode string) string {
	c.locker.Lock()
	res, ok := domain.PaymentReqResCodes[resCode]
	c.locker.Unlock()
	if !ok {
		return ""
	}
	return res
}
