package api

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

func (c *Client) get(u *url.URL) ([]byte, error) {
	if u.Scheme == "" {
		u.Scheme = scheme
	}
	if u.Host == "" {
		u.Host = host
	}
	// Create a new GET request
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	values := u.RawQuery
	if values != "" {
		values = "?" + values
	}

	// Add the required headers
	for key, value := range c.headers(req.Method, u.Path+values, "") {
		req.Header.Set(key, value)
	}

	// Send the request and get the response
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	// Close the response body
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (c *Client) post(u *url.URL, body []byte) ([]byte, error) {
	if u.Scheme == "" {
		u.Scheme = scheme
	}
	if u.Host == "" {
		u.Host = host
	}
	// Create a new POST request
	req, err := http.NewRequest("POST", u.String(), bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	// Add the required headers
	for key, value := range c.headers(req.Method, u.Path, string(body)) {
		req.Header.Set(key, value)
	}
	req.Header.Set("Content-Type", "application/json")

	// Send the request and get the response
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	// Close the response body
	defer resp.Body.Close()

	// Read the response body
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (c *Client) headers(method, path, body string) map[string]string {
	// Generate the timestamp and prehash string
	timestamp := time.Now().UTC().Format("2006-01-02T15:04:05Z")
	prehashString := timestamp + method + path + body

	// Sign the prehash string with the secret key using HMAC SHA256
	h := hmac.New(sha256.New, []byte(c.SecretKey))
	h.Write([]byte(prehashString))
	signature := h.Sum(nil)

	// Encode the signature in Base64 format
	sig := base64.StdEncoding.EncodeToString(signature)

	// Add the required headers
	headers := map[string]string{
		"OK-ACCESS-KEY":        c.ApiKey,
		"OK-ACCESS-SIGN":       sig,
		"OK-ACCESS-TIMESTAMP":  timestamp,
		"OK-ACCESS-PASSPHRASE": c.Passphrase,
	}

	return headers
}

type balanceResp struct {
	Code string `json:"code"`
	Data []struct {
		Available string `json:"availBal"`
		Balance   string `json:"bal"`
		Currency  string `json:"ccy"`
		Frozen    string `json:"frozenBal"`
	} `json:"data"`
}

func (c *Client) balance() (*balanceResp, error) {
	u := url.URL{
		Path: "/api/v5/asset/balances",
	}
	body, err := c.get(&u)
	if err != nil {
		return nil, err
	}
	var res balanceResp
	if err := json.Unmarshal(body, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

type currencyResp struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Data []struct {
		CanDep               bool   `json:"canDep"`
		CanInternal          bool   `json:"canInternal"`
		CanWd                bool   `json:"canWd"`
		Ccy                  string `json:"ccy"`
		Chain                string `json:"chain"`
		DepQuotaFixed        string `json:"depQuotaFixed"`
		DepQuoteDailyLayer2  string `json:"depQuoteDailyLayer2"`
		LogoLink             string `json:"logoLink"`
		MainNet              bool   `json:"mainNet"`
		MaxFee               string `json:"maxFee"`
		MaxFeeForCtAddr      string `json:"maxFeeForCtAddr"`
		MaxWd                string `json:"maxWd"`
		MinDep               string `json:"minDep"`
		MinDepArrivalConfirm string `json:"minDepArrivalConfirm"`
		MinFee               string `json:"minFee"`
		MinFeeForCtAddr      string `json:"minFeeForCtAddr"`
		MinWd                string `json:"minWd"`
		MinWdUnlockConfirm   string `json:"minWdUnlockConfirm"`
		Name                 string `json:"name"`
		NeedTag              bool   `json:"needTag"`
		UsedDepQuotaFixed    string `json:"usedDepQuotaFixed"`
		UsedWdQuota          string `json:"usedWdQuota"`
		WdQuota              string `json:"wdQuota"`
		WdTickSz             string `json:"wdTickSz"`
	} `json:"data"`
}

func (c *Client) currency(ticker string) (*currencyResp, error) {
	u := url.URL{
		Path: "/api/v5/asset/currencies",
	}
	values := url.Values{}
	values.Add("ccy", ticker)
	u.RawQuery = values.Encode()

	body, err := c.get(&u)
	if err != nil {
		return nil, err
	}
	var res currencyResp
	if err := json.Unmarshal(body, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

type withdrawReq struct {
	Amt    string `json:"amt"`
	Fee    string `json:"fee"`
	Dest   string `json:"dest"`
	Ccy    string `json:"ccy"`
	Chain  string `json:"chain"`
	ToAddr string `json:"toAddr"`
}

type withdrawResp struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Data []struct {
		Amt      string `json:"amt"`
		WdId     string `json:"wdId"`
		Ccy      string `json:"ccy"`
		ClientId string `json:"clientId"`
		Chain    string `json:"chain"`
	} `json:"data"`
}

func (r withdrawResp) IsOK() bool {
	return r.Code == "0"
}

func (c *Client) withdraw(req *withdrawReq) (*withdrawResp, error) {
	u := url.URL{
		Path: "/api/v5/asset/withdrawal",
	}
	reqBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	respBody, err := c.post(&u, reqBody)
	if err != nil {
		return nil, err
	}
	var res withdrawResp
	if err := json.Unmarshal(respBody, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
