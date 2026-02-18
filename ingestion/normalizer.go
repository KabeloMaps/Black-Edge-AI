package ingestion

import (
	"blackedge-backend/models"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"regexp"
	"strings"
	"time"
)

func NormalizeManga(scraped []models.ScrapedManga, sourceName string) []models.NormalizedManga {

	var normalized []models.NormalizedManga
	seen := make(map[string]bool)

	re := regexp.MustCompile(`[^\w-]`)

	for _, m := range scraped {

		title := strings.TrimSpace(strings.Title(m.Title))
		slug := strings.ToLower(strings.ReplaceAll(title, " ", "-"))
		slug = re.ReplaceAllString(slug, "-")

		status := strings.ToLower(strings.TrimSpace(m.Status))
		if status != "ongoing" && status != "completed" {
			status = "unknown"
		}

		cleanGenres := []string{}
		for _, g := range m.Genres {
			gg := strings.TrimSpace(strings.ToLower(g))
			if gg != "" {
				cleanGenres = append(cleanGenres, gg)
			}
		}

		hash := md5.Sum([]byte(title + m.SourceURL))
		id := hex.EncodeToString(hash[:])

		if seen[id] {
			fmt.Println("⛔ Duplicate skipped:", title)
			continue
		}
		seen[id] = true

		now := time.Now().Unix()

		normalized = append(normalized, models.NormalizedManga{
			ID:          id,
			Title:       title,
			Slug:        slug,
			Description: m.Description,
			Author:      m.Author,
			CoverImage:  m.CoverImage,
			Genres:      cleanGenres,
			Language:    m.Language,
			Status:      status,
			Source:      sourceName,
			SourceURL:   m.SourceURL,
			CreatedAt:   now,
			UpdatedAt:   now,
			Version:     1,
		})
	}

	fmt.Printf("📦 Normalized items: %d\n", len(normalized))
	return normalized
}
