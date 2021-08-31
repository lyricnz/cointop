package api

import (
	bin "github.com/miguelmota/cointop/pkg/api/impl/binance"
	cg "github.com/miguelmota/cointop/pkg/api/impl/coingecko"
	cmc "github.com/miguelmota/cointop/pkg/api/impl/coinmarketcap"
)

// NewCMC new CoinMarketCap API
func NewCMC(apiKey string) Interface {
	return cmc.NewCMC(apiKey)
}

// NewCC new CryptoCompare API
func NewCC() {
	// TODO
}

// NewCG new CoinGecko API
func NewCG(perPage, maxPages uint) Interface {
	return cg.NewCoinGecko(&cg.Config{
		PerPage:  perPage,
		MaxPages: maxPages,
	})
}

func NewBin(apiKey string, secretKey string) Interface {
	return bin.NewBinance(apiKey, secretKey)
}
