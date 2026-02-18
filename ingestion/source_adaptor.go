package ingestion

import "blackedge-backend/models"

type SourceAdapter interface {
	Name() string
	Scrape() ([]models.ScrapedManga, error)
}
