package ingestion

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

// NormalizeChapters converts scraped chapters into canonical format
func NormalizeChapters(scraped []ScrapedChapter, mangaTitle string) []NormalizedChapter {
	var normalized []NormalizedChapter
	seen := make(map[string]bool)

	for _, c := range scraped {
		// Clean title
		title := strings.TrimSpace(c.Title)

		// Generate unique hash
		hash := md5.Sum([]byte(mangaTitle + title + c.URL))
		id := hex.EncodeToString(hash[:])

		// Skip duplicates
		if seen[id] {
			continue
		}
		seen[id] = true

		normalized = append(normalized, NormalizedChapter{
			ID:      id,
			Title:   title,
			URL:     c.URL,
			Release: c.Release,
			Number:  c.Number,
		})
	}

	return normalized
}
