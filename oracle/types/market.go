package types

import (
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Asset struct {
	Symbol string
	Denom  string
}

type AssetMarket struct {
	Asset
	MarketCap         math.Int
	TargetUtilization sdk.Dec
	MaxLTV            sdk.Dec
	LiqLTV            sdk.Dec
}
