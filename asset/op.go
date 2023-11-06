package asset

import "fmt"

func OP(amt, addr string) *Asset {
	return &Asset{
		Amt:    amt,
		Dest:   "4",
		Ccy:    "OP",
		Chain:  "OP-Optimism",
		ToAddr: addr,
	}
}

func ToOP(amt, chain, addr string) (*Asset, error) {
	switch chain {
	case "optimism":
		return OP(amt, addr), nil
	}
	return nil, fmt.Errorf("chain=%s: %w", chain, ErrChainIsNotSupported)
}
