package domain

import "time"

type SadadIpgConfig struct {
	Timeout     time.Duration `json:"timeout"`
	BaseUrl     string        `json:"baseUrl"`
	MerchantId  string        `json:"merchantId"`
	TerminalId  string        `json:"terminalId"`
	TerminalKey string        `json:"terminalKey"`
	ShaparakIv  string        `json:"shaparakIv"`
	ShaparakKey string        `json:"shaparakKey"`
}
