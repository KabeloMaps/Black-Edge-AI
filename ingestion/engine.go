package ingestion

import (
	"blackedge-backend/models"
	"fmt"
)

func RunIngestion(source SourceAdapter) ([]models.NormalizedManga, error) {

	fmt.Println("🔁 Ingesting from:", source.Name())

	raw, err := source.Scrape()
	if err != nil {
		return nil, err
	}

	fmt.Println("📥 Raw items scraped:", len(raw))

	normalized := NormalizeManga(raw, source.Name())

	return normalized, nil
}
