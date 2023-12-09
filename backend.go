package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joe5saia/cautious-memory/internal/models"
	"github.com/joe5saia/cautious-memory/internal/routers"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	db := setupDatabase()
	db.AutoMigrate(&models.Profile{}, &models.SubField{})

	router := gin.Default()

	router.GET("/profiles/:id", routers.GetProfile(db))

	router.Run(":8080")
}

func setupDatabase() *gorm.DB {
	dsn := "host=localhost user=postgres password=password dbname=postgres port=5432 sslmode=disable TimeZone=America/New_York"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	return db
}

func setupRouter() {
}
