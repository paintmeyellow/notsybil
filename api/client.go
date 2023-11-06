package api

import (
	"errors"
	"fmt"
	"net/http"

	"notsybil/asset"
)

const (
	scheme = "https"
	host   = "www.okx.com"
)

var (
	ErrFeeIsEmpty   = errors.New("fee is empty")
	ErrAssetIsEmpty = errors.New("asset is empty")
)

type Client struct {
	ApiKey     string
	SecretKey  string
	Passphrase string
	HTTPClient *http.Client
}

func NewClient(apiKey, secretKey, passphrase string) *Client {
	return &Client{
		ApiKey:     apiKey,
		SecretKey:  secretKey,
		Passphrase: passphrase,
		HTTPClient: &http.Client{},
	}
}

type AssetBalance struct {
	Currency  string
	Available string
}

func (c *Client) AssetsBalance(assets []*asset.Asset) ([]*AssetBalance, error) {
	resp, err := c.balance()
	if err != nil {
		return nil, fmt.Errorf("get balance: %w", err)
	}
	unq := make(map[string]struct{}, 0)
	for _, a := range assets {
		unq[a.Ccy] = struct{}{}
	}
	var bal []*AssetBalance
	for _, dt := range resp.Data {
		if _, ok := unq[dt.Currency]; ok {
			bal = append(bal, &AssetBalance{Currency: dt.Currency, Available: dt.Available})
		}
	}
	return bal, nil
}

type Withdraw struct {
	Asset *asset.Asset
	Fee   string
}

func (c *Client) Withdraw(wd *Withdraw) (string, error) {
	if a := wd.Asset; a == nil {
		return "", ErrAssetIsEmpty
	}
	req := withdrawReq{
		Amt:    wd.Asset.Amt,
		Fee:    wd.Fee,
		Dest:   wd.Asset.Dest,
		Ccy:    wd.Asset.Ccy,
		Chain:  wd.Asset.Chain,
		ToAddr: wd.Asset.ToAddr,
	}
	resp, err := c.withdraw(&req)
	if err != nil {
		return "", err
	}
	if !resp.IsOK() || len(resp.Data) == 0 {
		return "", fmt.Errorf("code=%s, msg=%s", resp.Code, resp.Msg)
	}
	return resp.Data[0].WdId, nil
}

func (c *Client) ComposeWithdraw(assets []*asset.Asset) ([]*Withdraw, error) {
	var wds []*Withdraw
	for _, a := range assets {
		fee, err := c.Fee(a)
		if err != nil {
			return nil, fmt.Errorf("get fee: %w", err)
		}
		if fee == "" {
			return nil, ErrFeeIsEmpty
		}
		wd := Withdraw{
			Asset: a,
			Fee:   fee,
		}
		wds = append(wds, &wd)
	}
	return wds, nil
}

func (c *Client) Fee(a *asset.Asset) (string, error) {
	ccy, err := c.currency(a.Ccy)
	if err != nil {
		return "", err
	}
	for _, item := range ccy.Data {
		if item.Chain == a.Chain && item.CanWd {
			return item.MinFee, nil
		}
	}
	return "", nil
}
