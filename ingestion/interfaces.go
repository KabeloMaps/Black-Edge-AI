package core

import "blackedge-backend/models"

// SourceAdapter represents a data source that emits scraped manga entries.
type SourceAdapter interface {
	Name() string
	Scrape() ([]models.ScrapedManga, error)
}
