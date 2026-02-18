package ingestion

import (
	"context"
	"fmt"

	"blackedge-backend/storage"
)

func RunPipeline(source SourceAdapter) error {
	fmt.Println("🔁 Ingesting from:", source.Name())

	// Scrape
	raw, err := source.Scrape()
	if err != nil {
		return err
	}
	fmt.Println("📥 Raw items scraped:", len(raw))

	// Normalize
	normalized := NormalizeManga(raw, source.Name())
	fmt.Println("📦 Normalized items:", len(normalized))

	// Store
	ctx := context.Background()
	err = storage.InsertMangaBatch(ctx, normalized)
	if err != nil {
		return err
	}

	fmt.Println("✅ Pipeline complete for:", source.Name())
	return nil
}
