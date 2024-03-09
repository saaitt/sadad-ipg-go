package usecase

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/saaitt/sadad-ipg-go/refund/domain"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func (c *client) GetToken() error {
	requestUrl := *c.baseUrl
	requestUrl.Path += "/api/v1/users/token"
	values := url.Values{}
	values.Set("username", c.config.Username)
	values.Set("password", c.config.Password)
	values.Set("grant_type", "password")

	req, err := http.NewRequest(
		http.MethodPost,
		requestUrl.String(),
		strings.NewReader(values.Encode()),
	)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	httpClient := &http.Client{
		Timeout: c.config.Timeout,
	}
	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	res := new(domain.RefundTokenResponse)
	//resMap := make(map[string]string)
	err = json.Unmarshal(b, &res)
	//err = json.Unmarshal(b, &resMap)
	if err != nil {
		return errors.New("error in unmarshal json into the model")
	}

	if res.AccessToken != "" {
		c.accessToken = res.AccessToken
		//exp, _ := strconv.Atoi(res.ExpiresIn)
		//c.expiresIn = time.Now().Add(time.Duration(exp))
		return nil
	}

	if c.accessToken != "" {
		return nil
	}
	errStr := fmt.Sprintf("payment request failed, status code: %v ,,,, body is: %+v", resp.StatusCode, string(b))
	return errors.New(
		errStr,
	)
}
