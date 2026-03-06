package registry

import "blackedge-backend/models"

var CryptoPairs = []models.Instrument{
	{
		Symbol:   "BTCUSD",
		Type:     "crypto",
		Base:     "BTC",
		Quote:    "USD",
		Exchange: "Binance",
	},
	{
		Symbol:   "ETHUSD",
		Type:     "crypto",
		Base:     "ETH",
		Quote:    "USD",
		Exchange: "Binance",
	},
	{
		Symbol:   "LTCUSD",
		Type:     "crypto",
		Base:     "LTC",
		Quote:    "USD",
		Exchange: "Binance",
	},
	{
		Symbol:   "XRPUSD",
		Type:     "crypto",
		Base:     "XRP",
		Quote:    "USD",
		Exchange: "Binance",
	},
	{
		Symbol: "ETHUSD",
		Type:   "crypto",
		Base:   "ETH",
		Quote:  "USD",
		Exchange: "Binance",
	},
}