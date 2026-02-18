package scrapers

import (
	"blackedge-backend/models"
)

type MangaKakalotScraper struct{}

func (s MangaKakalotScraper) Name() string {
	return "MangaKakalot"
}

func (s MangaKakalotScraper) Scrape() ([]models.ScrapedManga, error) {

	// TEMP test data (real scraper next step)
	data := []models.ScrapedManga{
		{
			Title:       "Solo Leveling",
			Description: "Hunter world dungeon system",
			Author:      "Chugong",
			CoverImage:  "https://img.sololeveling.jpg",
			Genres:      []string{"action", "fantasy"},
			Language:    "en",
			Status:      "completed",
			SourceURL:   "https://mangakakalot.com/solo-leveling",
		},
	}

	return data, nil
}
