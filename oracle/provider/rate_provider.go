package provider

import "price-feeder/oracle/types"

const (
	ProviderGhost Name = "ghost"
)

type rate_provider struct {
	provider
	markets      map[string]types.AssetMarket
	rate_tickers map[string]types.TickerRate
}
