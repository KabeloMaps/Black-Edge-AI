package models

type Instrument struct {
	Symbol string `json:"symbol"`
	Type  string `json:"type"`
	Base   string `json:"base"`
	Quote  string `json:"quote"`
	Exchange string `json:"exchange"`
}