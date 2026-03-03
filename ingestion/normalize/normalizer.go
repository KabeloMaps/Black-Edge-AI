package normalize

import (
	"strings"
	"time"

	"blackedge-backend/models"
)

// NormalizeManga converts raw scraped entries into normalized models, deduplicating by hash.
func NormalizeManga(scraped []models.ScrapedManga, sourceName string) []models.NormalizedManga {
	normalized := make([]models.NormalizedManga, 0, len(scraped))
	seen := make(map[string]struct{})
	now := time.Now()

	for _, raw := range scraped {
		if raw.Title == "" || raw.URL == "" {
			continue
		}

		entry := models.NormalizedManga{
			Title:       strings.TrimSpace(raw.Title),
			Description: strings.TrimSpace(raw.Description),
			CoverURL:    strings.TrimSpace(raw.CoverURL),
			URL:         strings.TrimSpace(raw.URL),
			Genres:      raw.Genres,
			Source:      sourceName,
			CreatedAt:   now,
			UpdatedAt:   now,
		}

		entry.Hash = HashIdentity(entry.URL)

		if _, exists := seen[entry.Hash]; exists {
			continue
		}
		seen[entry.Hash] = struct{}{}
		normalized = append(normalized, entry)
	}

	return normalized
}
