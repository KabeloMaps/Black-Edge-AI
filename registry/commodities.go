package registry

import "blackedge-backend/models"

var Commodities = []models.Instrument{
	{
		Symbol:   "XAUUSD",
		Type:     "commodity",
		Base:     "XAU",
		Quote:    "USD",
	},
	{
		Symbol:   "XAGUSD",
		Type:     "commodity",
		Base:     "XAG",
		Quote:    "USD",
	},
	{
		Symbol:   "WTI",
		Type:     "commodity",
		Base:     "OIL",
		Quote:    "USD",
	},
}