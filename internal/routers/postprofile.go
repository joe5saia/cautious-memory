package routers

import (
	"encoding/json"
	"fmt"

	"net/http"

	"github.com/joe5saia/cautious-memory/internal/models"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func PostProfile(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		var newProfile models.Profile

		// print the request headers
		fmt.Println(c.Request.Header)

		if err := c.BindJSON(&newProfile); err != nil {
			return
		}

		jsonified_newProfile, err := json.Marshal(newProfile)

		if err != nil {
			fmt.Println("Error marshalling newProfile")
			fmt.Println(err)
			return
		}

		// Print "New Profile is {newProfile}"
		fmt.Println("New Profile is", string(jsonified_newProfile))

		// Save the new profile to the database
		if result := db.Create(&newProfile); result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}

		c.JSON(http.StatusCreated, newProfile)

	}
}
