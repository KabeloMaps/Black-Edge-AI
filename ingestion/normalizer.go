package ingestion

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"regexp"
	"strings"
)

type NormalizedManga struct {
	ID          string   `json:"id"`          // Unique hash ID
	Title       string   `json:"title"`       // Normalized title
	Slug        string   `json:"slug"`        // URL-friendly
	Description string   `json:"description"` // Optional
	Author      string   `json:"author"`      // Optional
	CoverImage  string   `json:"cover_image"` // Optional
	Genres      []string `json:"genres"`      // Optional
	Language    string   `json:"language"`    // Optional
	Status      string   `json:"status"`      // Ongoing/Completed/Unknown
	Source      string   `json:"source"`      // Name of scraper
	SourceURL   string   `json:"source_url"`  // Original URL
}

// NormalizeManga takes raw scraped data and returns normalized entries
func NormalizeManga(scraped []ScrapedManga, sourceName string) []NormalizedManga {
	var normalized []NormalizedManga
	seen := make(map[string]bool) // deduplication

	// Regex for slug cleanup (remove non-alphanumeric except dash)
	re := regexp.MustCompile(`[^\w-]`)

	for _, m := range scraped {
		title := strings.TrimSpace(strings.Title(m.Title))
		slug := strings.ToLower(strings.ReplaceAll(title, " ", "-"))
		slug = re.ReplaceAllString(slug, "-")

		status := strings.ToLower(strings.TrimSpace(m.Status))
		if status != "ongoing" && status != "completed" {
			status = "unknown"
		}

		var genres []string
		for _, g := range m.Genres {
			if trimmed := strings.TrimSpace(g); trimmed != "" {
				genres = append(genres, trimmed)
			}
		}

		hash := md5.Sum([]byte(title + m.SourceURL))
		id := hex.EncodeToString(hash[:])

		if seen[id] {
			fmt.Println("⛔ Duplicate skipped:", title)
			continue
		}
		seen[id] = true

		normalized = append(normalized, NormalizedManga{
			ID:          id,
			Title:       title,
			Slug:        slug,
			Description: strings.TrimSpace(m.Description),
			Author:      strings.TrimSpace(m.Author),
			CoverImage:  strings.TrimSpace(m.CoverImage),
			Genres:      genres,
			Language:    strings.TrimSpace(m.Language),
			Status:      status,
			Source:      sourceName,
			SourceURL:   strings.TrimSpace(m.SourceURL),
		})
	}

	fmt.Printf("📦 Normalized items: %d\n", len(normalized))
	return normalized
}
