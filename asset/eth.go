package asset

import "fmt"

func ETHERC20(amt, addr string) *Asset {
	return &Asset{
		Amt:    amt,
		Dest:   "4",
		Ccy:    "ETH",
		Chain:  "ETH-ERC20",
		ToAddr: addr,
	}
}

func ETHArbitrumOne(amt, addr string) *Asset {
	return &Asset{
		Amt:    amt,
		Dest:   "4",
		Ccy:    "ETH",
		Chain:  "ETH-ArbitrumOne",
		ToAddr: addr,
	}
}

func ETHzkSyncEra(amt, addr string) *Asset {
	return &Asset{
		Amt:    amt,
		Dest:   "4",
		Ccy:    "ETH",
		Chain:  "ETH-zkSync Era",
		ToAddr: addr,
	}
}

func ETHStarknet(amt, addr string) *Asset {
	return &Asset{
		Amt:    amt,
		Dest:   "4",
		Ccy:    "ETH",
		Chain:  "ETH-Starknet",
		ToAddr: addr,
	}
}

func ETHOptimism(amt, addr string) *Asset {
	return &Asset{
		Amt:    amt,
		Dest:   "4",
		Ccy:    "ETH",
		Chain:  "ETH-Optimism",
		ToAddr: addr,
	}
}

func ETHLinea(amt, addr string) *Asset {
	return &Asset{
		Amt:    amt,
		Dest:   "4",
		Ccy:    "ETH",
		Chain:  "ETH-Linea",
		ToAddr: addr,
	}
}

func ToETH(amt, chain, addr string) (*Asset, error) {
	switch chain {
	case "ethereum":
		return ETHERC20(amt, addr), nil
	case "starknet":
		return ETHStarknet(amt, addr), nil
	case "optimism":
		return ETHOptimism(amt, addr), nil
	case "arbitrum-one":
		return ETHArbitrumOne(amt, addr), nil
	case "zksync-era":
		return ETHzkSyncEra(amt, addr), nil
	case "linea":
		return ETHLinea(amt, addr), nil
	}
	return nil, fmt.Errorf("chain=%s: %w", chain, ErrChainIsNotSupported)
}
