package main

import (
	"context"
	"log"

	"blackedge-backend/api"
	"blackedge-backend/ingestion"
	"blackedge-backend/storage"

	"github.com/gin-gonic/gin"
)

func main() {

	// Mongo
	storage.ConnectMongo("mongodb://localhost:27017", "blackedge")

	// Ingestion
	source := ingestion.PublicDirectorySource{}
	items, err := ingestion.RunIngestion(source)
	if err != nil {
		log.Fatal("Ingestion failed:", err)
	}

	// Store
	ctx := context.Background()
	for _, m := range items {
		err := storage.InsertManga(ctx, m)
		if err != nil {
			log.Println("Insert failed:", err)
		}
	}

	// API
	r := gin.Default()
	r.GET("/manga", api.GetMangaList)
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	log.Println("🚀 Server listening on :8080")
	r.Run(":8080")
}
