package api

import (
	"context"
	"net/http"

	"blackedge-backend/models"
	"blackedge-backend/storage"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func GetMangaList(c *gin.Context) {
	ctx := context.Background()

	cursor, err := storage.MangaCollection.Find(ctx, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "database query failed",
			"details": err.Error(),
		})
		return
	}
	defer cursor.Close(ctx)

	var manga []models.NormalizedManga

	if err := cursor.All(ctx, &manga); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to decode data",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"count": len(manga),
		"data":  manga,
	})
}
