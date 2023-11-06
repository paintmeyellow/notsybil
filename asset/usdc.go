package asset

import "fmt"

func USDCERC20(amt, addr string) *Asset {
	return &Asset{
		Amt:    amt,
		Dest:   "4",
		Ccy:    "USDC",
		Chain:  "USDC-ERC20",
		ToAddr: addr,
	}
}

func USDCTRC20(amt, addr string) *Asset {
	return &Asset{
		Amt:    amt,
		Dest:   "4",
		Ccy:    "USDC",
		Chain:  "USDC-TRC20",
		ToAddr: addr,
	}
}

func USDCOptimism(amt, addr string) *Asset {
	return &Asset{
		Amt:    amt,
		Dest:   "4",
		Ccy:    "USDC",
		Chain:  "USDC-Optimism",
		ToAddr: addr,
	}
}

func USDCArbitrumOne(amt, addr string) *Asset {
	return &Asset{
		Amt:    amt,
		Dest:   "4",
		Ccy:    "USDC",
		Chain:  "USDC-Arbitrum One",
		ToAddr: addr,
	}
}

func USDCPolygon(amt, addr string) *Asset {
	return &Asset{
		Amt:    amt,
		Dest:   "4",
		Ccy:    "USDC",
		Chain:  "USDC-Polygon",
		ToAddr: addr,
	}
}

func ToUSDC(amt, chain, addr string) (*Asset, error) {
	switch chain {
	case "ethereum":
		return USDCERC20(amt, addr), nil
	case "tron":
		return USDCTRC20(amt, addr), nil
	case "optimism":
		return USDCOptimism(amt, addr), nil
	case "arbitrum-one":
		return USDCArbitrumOne(amt, addr), nil
	case "polygon":
		return USDCPolygon(amt, addr), nil
	}
	return nil, fmt.Errorf("chain=%s: %w", chain, ErrChainIsNotSupported)
}
