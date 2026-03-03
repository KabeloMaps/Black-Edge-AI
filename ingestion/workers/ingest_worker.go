package workers

import (
	"encoding/json"
	"log"

	"blackedge-backend/ingestion/normalize"
	"blackedge-backend/models"
)

// Process handles incoming queue data
func Process(data []byte, sourceName string) {

	// Step 1: Decode raw JSON into scraped model
	var scraped []models.ScrapedManga
	err := json.Unmarshal(data, &scraped)
	if err != nil {
		log.Println("failed to unmarshal scraped data:", err)
		return
	}

	// Step 2: Normalize scraped data
	normalized := normalize.NormalizeManga(scraped, sourceName)

	log.Println("Normalized items:", len(normalized))

	// Step 3: (Next step later: save to DB)
}