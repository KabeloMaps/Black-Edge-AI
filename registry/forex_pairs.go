package registry

import "blackedge-backend/models"

var ForexPairs = []models.Instrument{
	{
		Symbol: "EURUSD",
		Type:  "forex",
		Base:   "EUR",
		Quote:  "USD",
	},
	{
		Symbol: "GBPUSD",
		Type:  "forex",
		Base:   "GBP",
		Quote:  "USD",
	},
	{
		Symbol: "USDJPY",
		Type:  "forex",
		Base:   "USD",
		Quote:  "JPY",
	},
	{
		Symbol: "USDCHF",
		Type:  "forex",
		Base:   "USD",
		Quote:  "CHF",
	},
	{
		Symbol: "AUDUSD",
		Type:  "forex",
		Base:   "AUD",
		Quote:  "USD",
	},
	{
		Symbol: "USDCAD",
		Type:  "forex",
		Base:   "USD",
		Quote:  "CAD",
	},
	{
		Symbol: "NZDUSD",
		Type:  "forex",
		Base:   "NZD",
		Quote:  "USD",
	},
}