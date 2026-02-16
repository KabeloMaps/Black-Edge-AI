package main

import (
	"blackedge-backend/config"
	"blackedge-backend/database"
	"blackedge-backend/models"
	"blackedge-backend/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()

	r := gin.Default()
	database.ConnectMongo(cfg)
	database.ConnectPostgres(cfg)
	database.DB.AutoMigrate(&models.Manga{}, &models.Chapter{},)

	routes.RegisterRoutes(r)
	r.Run(":" + cfg.ServerPort)
}