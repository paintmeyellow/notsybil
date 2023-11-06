package asset

import "fmt"

func APT(amt, addr string) *Asset {
	return &Asset{
		Amt:    amt,
		Dest:   "4",
		Ccy:    "APT",
		Chain:  "APT-Aptos",
		ToAddr: addr,
	}
}

func ToAPT(amt, chain, addr string) (*Asset, error) {
	switch chain {
	case "aptos":
		return APT(amt, addr), nil
	}
	return nil, fmt.Errorf("chain=%s: %w", chain, ErrChainIsNotSupported)
}
