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
	"strconv"
	"strings"
	"time"
)

func (c *client) CreatePaymentLink(request *domain.PaymentLinkRequest) (*domain.PaymentLinkResp, error) {
	req, err := c.makePaymentRequestBody(request)
	if err != nil {
		return nil, err
	}
	httpClient := &http.Client{
		Timeout: c.config.Timeout,
	}
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	res := new(domain.PaymentLinkResp)
	err = json.Unmarshal(b, &res)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == 200 {
		return c.handlePaymentRequestResponse(res)
	}
	return nil, errors.New(fmt.Sprintf("payment request failed, status code: %v --- body is: %+v", resp.StatusCode, string(b)))
}

func (c *client) makePaymentRequestBody(request *domain.PaymentLinkRequest) (*http.Request, error) {
	request.MerchantId = c.config.MerchantId
	request.TerminalId = c.config.TerminalId
	signData, err := c.MakeSignData(request.TerminalId, request.OrderId, request.Amount)
	if err != nil {
		return nil, err
	}
	request.SignData = signData
	nationalCodeEnc, err := c.EncryptNationalCode(strconv.Itoa(request.UserType), request.NationalCode)
	if err != nil {
		return nil, err
	}
	request.NationalCodeEnc = nationalCodeEnc

	requestUrl := *c.baseUrl
	requestUrl.Path += "/v0/Request/PaymentRequest"
	values := url.Values{}
	values.Set("MerchantId", request.MerchantId)
	values.Set("TerminalId", request.TerminalId)
	values.Set("Amount", strconv.FormatInt(request.Amount, 10))
	values.Set("OrderId", strconv.FormatInt(request.OrderId, 10))
	values.Set("LocalDateTime", request.LocalDateTime.Format(time.RFC3339))
	values.Set("ReturnUrl", request.ReturnUrl)
	values.Set("SignData", request.SignData)
	values.Set("UserId", request.UserId)
	values.Set("NationalCode", request.NationalCode)
	values.Set("PanAuthenticationType", "4")
	values.Set("NationalCodeEnc", request.NationalCodeEnc)
	values.Set("ApplicationName", request.ApplicationName)

	req, err := http.NewRequest(
		http.MethodPost,
		requestUrl.String(),
		strings.NewReader(values.Encode()),
	)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	return req, nil
}

func (c *client) handlePaymentRequestResponse(res *domain.PaymentLinkResp) (*domain.PaymentLinkResp, error) {
	if res.ResCode != domain.PaymentReqResCodeSuccess {
		return nil, errors.New(fmt.Sprintf("payment request was not successful, the response was %v - %v", res.ResCode, res.ResCodeStr))
	}
	redirectUrl := *c.baseUrl
	redirectUrl.Path = fmt.Sprintf("/Purchase")
	ru := fmt.Sprintf("%v?token=%s", redirectUrl.String(), res.Token)
	res.PaymentUrl = ru
	return res, nil
}

func (c *client) MakeSignData(terminalId string, orderId int64, amount int64) (string, error) {
	// TerminalId;OrderId;Amount
	str := fmt.Sprintf("%v;%v;%v", terminalId, orderId, amount)
	signDataBytes := []byte(str)
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

func (c *client) EncryptNationalCode(userType, nationalCode string) (string, error) {
	natCode := fmt.Sprintf("%v|%v|%v", userType, nationalCode, utils.GenerateNewCode7Dig())
	sKey, err := base64.StdEncoding.DecodeString(c.config.ShaparakKey)
	if err != nil {
		return "", errors.New("decode_shaparak_key_error")
	}
	iv, err := base64.StdEncoding.DecodeString(c.config.ShaparakIv)
	if err != nil {
		return "", errors.New("decode_shaparak_key_error")
	}
	encrypt, err := utils.AesCbcPkcs7Encrypt(natCode, sKey, iv)
	if err != nil {
		return "", err
	}
	return encrypt, nil
}
