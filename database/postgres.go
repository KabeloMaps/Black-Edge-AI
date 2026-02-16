package database

import(
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectPostgres(){
	dsn := "host=localhost user=postgres password=postgres dbname=blackedge port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to PostgreSQL database!")
	}
	DB = db
	fmt.Println("Connected to PostgreSQL database successfully!")
}