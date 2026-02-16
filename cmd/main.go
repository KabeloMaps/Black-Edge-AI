package main

import (
	"blackedge-backend/database"
	"blackedge-backend/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	database.ConnectPostgres()
	database.ConnectMongo()

	routes.RegisterRoutes(r)

	r.Run(":8080")
}