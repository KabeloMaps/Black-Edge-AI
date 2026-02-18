package ingestion

import "blackedge-backend/models"

type PublicDirectorySource struct{}

func (p PublicDirectorySource) Name() string {
	return "PublicDirectorySource"
}

func (p PublicDirectorySource) Scrape() ([]models.ScrapedManga, error) {

	data := []models.ScrapedManga{
		{
			Title:       "Test Manga",
			Description: "Public domain test",
			Author:      "Unknown",
			CoverImage:  "",
			Genres:      []string{"classic"},
			Language:    "en",
			Status:      "unknown",
			SourceURL:   "https://publicsource.org/test",
		},
	}

	return data, nil
}
