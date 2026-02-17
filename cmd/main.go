package main

import (
	"blackedge-backend/config"
	"blackedge-backend/database"
	"blackedge-backend/ingestion"
	"blackedge-backend/models"
	"blackedge-backend/routes"
	"blackedge-backend/services"
	"blackedge-backend/storage"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()

	r := gin.Default()

	// DB Connections
	database.ConnectPostgres(cfg)
	database.ConnectMongo(cfg)

	// Auto-migrate
	database.DB.AutoMigrate(
		&models.Manga{},
		&models.Chapter{},
	)

	// 🔥 SCRAPER PIPELINE (THIS IS THE MISSING LINK)
	source := ingestion.PublicSource{}
	normalizedData, err := ingestion.RunIngestion(source)
	if err != nil {
		panic(err)
	}

	_ = storage.SaveNormalizedManga(normalizedData)
	_ = services.SaveToPostgres(normalizedData) // optional: map NormalizedManga → Manga struct

	routes.RegisterRoutes(r)
	_ = r.Run(cfg.ServerPort)
}
