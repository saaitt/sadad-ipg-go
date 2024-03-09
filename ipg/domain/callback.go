package domain

type SadadCallback struct {
	ResCode      string `json:"ResCode"`
	Description  string `json:"Description"`
	PrimaryAccNo string `json:"PrimaryAccNo"`
	Token        string `json:"token"`
	HashedCardNo string `json:"hashedCardNo"`
	OrderId      string `json:"OrderId"`
}

func (c SadadCallback) GetOrderID() string { return c.OrderId }
func (c SadadCallback) GetToken() string   { return c.Token }
func (c SadadCallback) GetStatus() string  { return c.ResCode }
