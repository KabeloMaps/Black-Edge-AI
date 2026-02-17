package services

import (
	"blackedge-backend/database"
	"blackedge-backend/ingestion"
	"blackedge-backend/models"
)

func SaveToPostgres(data []ingestion.NormalizedManga) error {
	for _, m := range data {
		record := models.Manga{
			Title:       m.Title,
			Description: m.Description,
			Author:      m.Author,
			Status:      m.Status,
			CoverImage:  m.CoverImage,
		}
		database.DB.Create(&record)
	}
	return nil
}
