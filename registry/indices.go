package registry

import "blackedge-backend/models"

var Indices = []models.Instrument{
	{
		Symbol:   "US30",
		Type:     "index",
	},
	{
		Symbol:   "SPX500",
		Type:     "index",
	},
	{
		Symbol:   "NAS100",
		Type:     "index",
	},
	{
		Symbol:   "GER40",
		Type:     "index",
	},
	{
		Symbol:   "UK100",
		Type:     "index",
	},
}