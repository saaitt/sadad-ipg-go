package usecase

import (
	"fmt"
	"github.com/saaitt/sadad-ipg-go/ipg/domain"
	"net/url"
	"sync"
)

type client struct {
	config  *domain.SadadIpgConfig
	locker  sync.Mutex
	baseUrl *url.URL
}

func Init(config *domain.SadadIpgConfig) *client {
	c := &client{
		config: config,
		locker: sync.Mutex{},
	}
	var err error
	c.baseUrl, err = url.Parse(c.config.BaseUrl)
	if err != nil {
		panic(fmt.Sprintf("error on initiate sadad client, base url is invalid, err: %v", err))
	}
	return c
}
