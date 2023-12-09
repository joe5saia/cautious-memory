package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/joe5saia/cautious-memory/internal/schemas"
)

func postProfile(c *gin.Context, db *gorm.DB) {
	var newProfile schemas.Profile

	// Call BindJSON to bind the received JSON to
	if err := c.BindJSON(&newProfile); err != nil {
		return
	}

	// Write the new profile to the database
	if err := db.Create(&newProfile).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

}
