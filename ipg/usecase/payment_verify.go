package usecase

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/saaitt/sadad-ipg-go/ipg/domain"
	"github.com/saaitt/sadad-ipg-go/utils"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func (c *client) ValidatePayment(data interface{}) (*domain.VerifyPayment, error) {
	req, err := c.makePaymentVerifyRequest(data)
	if err != nil {
		return nil, err
	}

	httpClient := &http.Client{
		Timeout: c.config.Timeout,
	}
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, errors.New("verify_sadad_payment_do_req")
	}
	defer resp.Body.Close()

	res, err := c.makeVerifyPaymentResponse(err, resp)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == 200 {
		return res, nil
	}
	return nil, errors.New(fmt.Sprintf("payment verify failed, status code: %v", resp.StatusCode))
}

func (c *client) makeVerifyPaymentResponse(err error, resp *http.Response) (*domain.VerifyPayment, error) {
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("verify_sadad_payment_read_response")
	}

	res := new(domain.VerifyPayment)
	err = json.Unmarshal(b, &res)
	if err != nil {
		return nil, errors.New("error in unmarshal json into the designated object, the response was: " + string(b))
	}
	return res, nil
}

func (c *client) makePaymentVerifyRequest(data interface{}) (*http.Request, error) {
	request := data.(domain.ValidateReq)
	signData, err := c.MakeVerifySignData(request.Token)
	if err != nil {
		return nil, nil
	}
	request.SignData = signData

	requestUrl := *c.baseUrl
	requestUrl.Path += "/v0/Advice/Verify"

	values := url.Values{}
	values.Set("Token", request.Token)
	values.Set("SignData", request.SignData)

	req, err := http.NewRequest(
		http.MethodPost,
		requestUrl.String(),
		strings.NewReader(values.Encode()),
	)
	if err != nil {
		return nil, nil
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	return req, err
}

func (c *client) MakeVerifySignData(token string) (string, error) {
	signDataBytes := []byte(token)
	sKey, err := base64.StdEncoding.DecodeString(c.config.TerminalKey)
	if err != nil {
		return "", errors.New("decode_sadad_key_error")
	}

	signData, err := utils.TripleEcbDesEncrypt(signDataBytes, sKey)
	if err != nil {
		log.Fatal(err)
	}
	b64SignData := base64.StdEncoding.EncodeToString(signData)
	if err != nil {
		return "", errors.New("decode_sadad_key_error")
	}
	return b64SignData, nil
}
