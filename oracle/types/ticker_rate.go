package types

import (
	"fmt"
	"time"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// TickerRate defines deposit rate and borrow rate information for an asset market.

type TickerRate struct {
	Supply             math.Int
	Available          math.Int
	CurrentUtilization sdk.Dec
	DepositRate        sdk.Dec   `json:"deposit_rate"`
	BorrowRate         sdk.Dec   `json:"borrow_rate"`
	Time               time.Time `json:"time"`
}

func NewTickerRate(deposit_rate string, borrow_rate string, timestamp time.Time) (TickerRate, error) {
	depositRateDec, err := sdk.NewDecFromStr(deposit_rate)
	if err != nil {
		return TickerRate{}, fmt.Errorf("failed to convert ticker deposit rate: %v", err)
	}
	borrowRateDec, err := sdk.NewDecFromStr(borrow_rate)
	if err != nil {
		return TickerRate{}, fmt.Errorf("failed to convert ticker borrow rate: %v", err)
	}

	ticker := TickerRate{
		DepositRate: depositRateDec,
		BorrowRate:  borrowRateDec,
		Time:        timestamp,
	}
	return ticker, nil
}
