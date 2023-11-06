package asset

import "fmt"

func USDTERC20(amt, addr string) *Asset {
	return &Asset{
		Amt:    amt,
		Dest:   "4",
		Ccy:    "USDT",
		Chain:  "USDT-ERC20",
		ToAddr: addr,
	}
}

func USDTTRC20(amt, addr string) *Asset {
	return &Asset{
		Amt:    amt,
		Dest:   "4",
		Ccy:    "USDT",
		Chain:  "USDT-TRC20",
		ToAddr: addr,
	}
}

func USDTOptimism(amt, addr string) *Asset {
	return &Asset{
		Amt:    amt,
		Dest:   "4",
		Ccy:    "USDT",
		Chain:  "USDT-Optimism",
		ToAddr: addr,
	}
}

func USDTArbitrumOne(amt, addr string) *Asset {
	return &Asset{
		Amt:    amt,
		Dest:   "4",
		Ccy:    "USDT",
		Chain:  "USDT-Arbitrum One",
		ToAddr: addr,
	}
}

func USDTPolygon(amt, addr string) *Asset {
	return &Asset{
		Amt:    amt,
		Dest:   "4",
		Ccy:    "USDT",
		Chain:  "USDT-Polygon",
		ToAddr: addr,
	}
}

func ToUSDT(amt, chain, addr string) (*Asset, error) {
	switch chain {
	case "ethereum":
		return USDTERC20(amt, addr), nil
	case "tron":
		return USDTTRC20(amt, addr), nil
	case "optimism":
		return USDTOptimism(amt, addr), nil
	case "arbitrum-one":
		return USDTArbitrumOne(amt, addr), nil
	case "polygon":
		return USDTPolygon(amt, addr), nil
	}
	return nil, fmt.Errorf("chain=%s: %w", chain, ErrChainIsNotSupported)
}
