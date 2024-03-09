package usecase

import (
	"fmt"
	"github.com/saaitt/sadad-ipg-go/refund/domain"
	"net/url"
	"sync"
	"time"
)

type client struct {
	config      *domain.SadadRefundConfig
	locker      sync.Mutex
	baseUrl     *url.URL
	accessToken string
	expiresIn   time.Time
}

func Init(refundConfig *domain.SadadRefundConfig) domain.SadadRefundClient {
	c := &client{
		config: refundConfig,
		locker: sync.Mutex{},
	}
	var err error
	c.baseUrl, err = url.Parse(c.config.BaseUrl)
	if err != nil {
		panic(fmt.Sprintf("error on initiate sadad refund client, base url is invalid, err: %v", err))
	}

	err = c.GetToken()
	if err != nil {
		fmt.Println(fmt.Errorf("SADAD REFUND GET TOKEN ERR: %+v", err))
	}

	return c
}
