package api

import (
	"blackedge-backend/database"
	"blackedge-backend/models"
	"github.com/gin-gonic/gin"
)

func GetMangaList(c *gin.Context) {
	var manga []models.Manga
	database.DB.Find(&manga)
	c.JSON(200, manga)
}