package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	docs "github.com/joe5saia/cautious-memory/docs"
	"github.com/joe5saia/cautious-memory/internal/handler"
	"github.com/joe5saia/cautious-memory/internal/model"
	"github.com/joe5saia/cautious-memory/internal/store"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// gin-swagger middleware
// swagger embed files

//	@title			Swagger Example API
//	@version		1.0
//	@description	This is a sample server celler server.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:8080
//	@BasePath	/api/v1

//	@securityDefinitions.basic	BasicAuth

// @externalDocs.description	OpenAPI
// @externalDocs.url			https://swagger.io/resources/open-api/
func main() {
	dbHost := os.Getenv("POSTGRES_HOSTNAME")
	dbPort := os.Getenv("POSTGRES_PORT")
	dbUser := os.Getenv("POSTGRES_USER")
	dbPass := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")

	// Initialize Gin router
	router := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"

	// Set up database connection
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbHost, dbUser, dbPass, dbName, dbPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// array of models
	models := []interface{}{
		&model.Profile{},
		&model.Subfield{},
		&model.Field{},
	}

	// AutoMigrate models
	db.AutoMigrate(models...)

	// Initialize store
	store := store.NewStore(db)

	// Initialize handlers
	profileHandler := handler.NewProfileHandler(store)
	subfieldHandler := handler.NewSubfieldHandler(store)

	v1 := router.Group("/api/v1")
	{
		v1.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Welcome to the API!",
			})
		})
		// Group all the profile routes
		profileRoutes := v1.Group("profiles")
		{
			profileRoutes.GET("/", profileHandler.GetAllProfiles)
			profileRoutes.GET("/:id", profileHandler.GetProfile)
			profileRoutes.POST("/", profileHandler.CreateProfile)
			profileRoutes.POST("/:id", profileHandler.UpdateProfile)
			profileRoutes.DELETE("/:id", profileHandler.DeleteProfile)
		}
		// Group all the subfield routes
		subfieldRoutes := v1.Group("subfields")
		{
			subfieldRoutes.GET("/", subfieldHandler.GetAllSubfields)
			subfieldRoutes.GET("/:id", subfieldHandler.GetSubfield)
			subfieldRoutes.POST("/", subfieldHandler.CreateSubfield)
			subfieldRoutes.POST("/:id", subfieldHandler.UpdateSubfield)
			subfieldRoutes.DELETE("/:id", subfieldHandler.DeleteSubfield)
		}

	}

	// Set up Gin Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Run the server
	router.Run(":8080")
}
