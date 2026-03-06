package registry

import "blackedge-backend/models"

var ZarPairs = []models.Instrument{
	{
		Symbol: "USDZAR",
		Type:  "forex",
		Base:   "USD",
		Quote:  "ZAR",
	},
	{
		Symbol: "EURZAR",
		Type:  "forex",
		Base:   "EUR",
		Quote:  "ZAR",
	},
	{
		Symbol: "GBPZAR",
		Type:  "forex",
		Base:   "GBP",
		Quote:  "ZAR",
	},
	{
		Symbol: "ZARJPY",
		Type:  "forex",
		Base:   "ZAR",
		Quote:  "JPY",
	},
}