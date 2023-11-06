package asset

import (
	"errors"
	"fmt"
)

var ErrCurrencyIsNotSupported = errors.New("currency is not supported")
var ErrChainIsNotSupported = errors.New("chain is not supported")

type Asset struct {
	Amt    string
	Dest   string
	Ccy    string
	Chain  string
	ToAddr string
}

func ToAsset(amt, ccy, chain, addr string) (*Asset, error) {
	switch ccy {
	case "eth":
		return ToETH(amt, chain, addr)
	case "op":
		return ToOP(amt, chain, addr)
	case "arb":
		return ToARB(amt, chain, addr)
	case "apt":
		return ToAPT(amt, chain, addr)
	case "usdt":
		return ToUSDT(amt, chain, addr)
	}
	return nil, fmt.Errorf("currency=%s: %w", ccy, ErrCurrencyIsNotSupported)
}
