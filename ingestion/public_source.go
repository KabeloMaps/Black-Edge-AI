package core

import "blackedge-backend/models"

type PublicDirectorySource struct{}

func (p PublicDirectorySource) Name() string {
	return "public_directory"
}

func (p PublicDirectorySource) Scrape() ([]models.ScrapedManga, error) {

	// Mock source for now
	data := []models.ScrapedManga{
		{
			Title:       "One Piece",
			Description: "Pirate adventure",
			CoverURL:    "https://example.com/op.jpg",
			URL:         "https://example.com/onepiece",
			Genres:      []string{"Action", "Adventure"},
		},
	}

	return data, nil
}