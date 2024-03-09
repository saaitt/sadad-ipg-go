package usecase

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/saaitt/sadad-ipg-go/refund/domain"
	"github.com/saaitt/sadad-ipg-go/utils"
	"io"
	"net/http"
	"strings"
)

func (c *client) NewRefund(refReq *domain.RefundRequest) (*domain.RefundResponse, error) {
	req, httpClient, err := c.makeRefundRequestPayload(refReq)
	if err != nil {
		return nil, err
	}

	res, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return c.makeRefundResponse(err, res)
}

func (c *client) makeRefundRequestPayload(refReq *domain.RefundRequest) (*http.Request, *http.Client, error) {
	requestUrl := *c.baseUrl
	requestUrl.Path += "/api/v1/refund/new"
	var err error

	refReq.TerminalID = c.config.TerminalId
	refReq.Sign, err = c.MakeNewRefundSign(refReq)
	if err != nil {
		return nil, nil, nil
	}
	payload := strings.NewReader(fmt.Sprintf(`{ 
    "TerminalID": "%v",
      "ClientRequestID": "%v",
      "RRN": "%v",
      "TxnDateTime": "%v",
      "TxnAmount": %v,
      "RefundAmount": %v,
      "TransferMethodID": %v,
      "WagePayMethodID": %v,
      "Description": "string",
      "Sign": "%v"
}`, refReq.TerminalID, refReq.ClientRequestID, refReq.RNN, refReq.TxnDateTime, refReq.TxnAmount, refReq.RefundAmount,
		refReq.TransferMethodID, refReq.WagePayMethodID, refReq.Sign))

	httpClient := &http.Client{}
	req, err := http.NewRequest(
		"POST",
		requestUrl.String(),
		payload,
	)
	if err != nil {
		return nil, nil, nil
	}

	if c.accessToken == "" {
		err := c.GetToken()
		if err != nil {
			return nil, nil, nil
		}
	}

	req.Header.Add("accept", "*/*")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", c.accessToken))
	req.Header.Add("Content-Type", "application/json-patch+json")
	return req, httpClient, nil
}

func (c *client) makeRefundResponse(err error, res *http.Response) (*domain.RefundResponse, error) {
	b, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	resp := make(map[string]interface{})
	//_ = json.Unmarshal([]byte(str), &r)
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return nil, err
	}
	var response domain.RefundResponse
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return nil, err
	}
	d := resp["Data"]
	response.ResCode = resp["ResCode"].(float64)
	response.ResMsg = resp["ResMsg"].(string)
	if response.ResCode == float64(0) {
		data := d.(map[string]interface{})
		response.Data.RefundID = data["RefundID"].(float64)
		//response.Data.WageAmount = data["WageAmount"].(string)
		//response.Data.CardToIbanWageAmount = data["CardToIbanWageAmount"].(string)
		response.Data.RNN = data["RRN"].(string)
		//response.Data.TxnAmount = data["TxnAmount"].(string)
		//response.Data.RefundAmount = data["RefundAmount"].(string)
		return &response, nil
	} else {
		//data := d.([]interface{})

		if response.ResMsg == "Authentication failed." || response.ResCode == float64(-1) {
			c.accessToken = ""
		}
		return nil, errors.New("refund error")
	}
}

func (c *client) MakeNewRefundSign(refReq *domain.RefundRequest) (string, error) {
	baseStr := refReq.MakeSignBase(c.config.Username)
	return utils.SignPKCS1v15FromXml(baseStr, c.config.PrKey)
}
