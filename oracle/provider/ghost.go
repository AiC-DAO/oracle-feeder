package provider

import (
	"time"
)

var (
	_                     Provider = (*FinProvider)(nil)
	ghostDefaultEndpoints          = Endpoint{
		Name: ProviderGhost,
		Urls: []string{
			"https://cosmos.directory/kujira",
			"https://lcd.kaiyo.kujira.setten.io",
			"https://lcd-kujira.mintthemoon.xyz",
		},
		PollInterval: 3 * time.Second,
	}
)

type GhostProvider struct {
	rate_provider
}
