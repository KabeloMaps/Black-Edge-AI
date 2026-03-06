package core

import (
	"context"
	"encoding/json"
	"log"

	corepkg "blackedge-backend/core"
	"blackedge-backend/ingestion/normalize"
	"blackedge-backend/models"
	"blackedge-backend/storage"
)

// Orchestrates full ingestion pipeline
func RunPipeline(source corepkg.Source) error {

	// 1) Scrape
	raw, err := source.Fetch()
	if err != nil {
		return err
	}

	// 2) Decode + normalize
	var scraped []models.ScrapedManga
	if err := json.Unmarshal(raw, &scraped); err != nil {
		return err
	}

	normalized := normalize.NormalizeManga(scraped, source.Name())

	// 3) Store
	ctx := context.Background()
	for _, m := range normalized {
		if err := storage.InsertManga(ctx, m); err != nil {
			log.Println("Insert failed:", err)
		}
	}

	log.Println("✅ Pipeline completed:", source.Name(), "items:", len(normalized))
	return nil
}