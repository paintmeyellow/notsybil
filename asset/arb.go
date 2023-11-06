package asset

import "fmt"

func ARBArbitrumOne(amt, addr string) *Asset {
	return &Asset{
		Amt:    amt,
		Dest:   "4",
		Ccy:    "ARB",
		Chain:  "ARB-Arbitrum One",
		ToAddr: addr,
	}
}

func ToARB(amt, chain, addr string) (*Asset, error) {
	switch chain {
	case "arbitrum-one":
		return ARBArbitrumOne(amt, addr), nil
	}
	return nil, fmt.Errorf("chain=%s: %w", chain, ErrChainIsNotSupported)
}
