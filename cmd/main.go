package main

import (
	"fmt"

	"github.com/joe5saia/cautious-memory/internal/handler"
	"github.com/joe5saia/cautious-memory/internal/model"
	"github.com/joe5saia/cautious-memory/internal/store"

	"log"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dbHost := os.Getenv("POSTGRES_HOSTNAME")
	dbPort := os.Getenv("POSTGRES_PORT")
	dbUser := os.Getenv("POSTGRES_USER")
	dbPass := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")

	// Initialize Gin router
	router := gin.Default()

	// Set up database connection
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbHost, dbUser, dbPass, dbName, dbPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// AutoMigrate models
	db.AutoMigrate(&model.Profile{}, &model.Subfield{})

	// Initialize store
	store := store.NewStore(db)

	// Initialize handlers
	profileHandler := handler.NewProfileHandler(store)
	subfieldHandler := handler.NewSubfieldHandler(store)

	// Set up routes
	router.GET("/profiles", profileHandler.GetAllProfiles)
	router.GET("/profile/:id", profileHandler.GetProfile)
	router.POST("/profile", profileHandler.CreateProfile)
	router.POST("/profile/:id", profileHandler.UpdateProfile)
	router.GET("/subfields", subfieldHandler.GetAllSubfields)
	router.GET("/subfield/:id", subfieldHandler.GetSubfield)
	router.DELETE("/subfield/:id", subfieldHandler.DeleteSubfield)
	router.POST("/subfield", subfieldHandler.CreateSubfield)
	router.POST("/subfield/:id", subfieldHandler.UpdateSubfield)

	// Run the server
	router.Run(":8080")
}
