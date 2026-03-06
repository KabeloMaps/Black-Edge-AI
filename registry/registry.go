package registry

import "blackedge-backend/models"

func GetAllInstruments() []models.Instrument {

	var instruments []models.Instrument
	
	instruments = append(instruments, CryptoPairs...)
	instruments = append(instruments, Commodities...)
	instruments = append(instruments, ForexPairs...)
	instruments = append(instruments, Indices...)
	instruments = append(instruments, Synthetics...)
	instruments = append(instruments, ZarPairs...)
	return instruments
}