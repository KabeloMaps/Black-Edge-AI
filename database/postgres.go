package database

import (
	"blackedge-backend/config"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

	func ConnectPostgres(cfg *config.Config) {
		fmt.Println("🔗 Postgres DSN:", cfg.PostgresDSN)
	
		db, err := gorm.Open(postgres.Open(cfg.PostgresDSN), &gorm.Config{})
		if err != nil {
			panic(err)
		}
	
		DB = db
		fmt.Println("✅ Postgres connected")
	}
