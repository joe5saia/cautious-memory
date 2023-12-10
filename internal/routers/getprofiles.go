package routers

import (
	"net/http"
	"strconv"

	"github.com/joe5saia/cautious-memory/internal/models"
	"github.com/joe5saia/cautious-memory/internal/repository"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

// getPageNumber extracts the page number from the query parameters of the given context.
// It returns the page number as an integer and any error encountered.
// If the page number is not valid or not present, it responds to the client with a 400 Bad Request status and an error message.
func getPageNumber(c *gin.Context) (int, error) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page number"})
		return 0, err
	}
	return page, nil
}

// getPageSize extracts the page size from the query parameters of the given context.
// It returns the page size as an integer and any error encountered.
// If the page size is not valid or not present, it responds to the client with a 400 Bad Request status and an error message.
func getPageSize(c *gin.Context) (int, error) {
	pageSize, err := strconv.Atoi(c.Query("pageSize"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page size"})
		return 0, err
	}
	return pageSize, nil
}

// fetchProfiles fetches the profiles from the database using the given repository.
// It returns the profiles and any error encountered.
// If there is an error, it responds to the client with a 500 Internal Server Error status and an error message.
func fetchProfiles(profileRepo repository.GenericRepository[models.Profile], c *gin.Context, page int, pageSize int) ([]models.Profile, error) {
	profiles, err := profileRepo.ReadAll(c, page, pageSize, "SubFields")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return nil, err
	}
	return profiles, nil
}

func GetProfiles(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		profileRepo := repository.NewGormRepository[models.Profile](db)

		// Get the page number from the query string
		page, err := getPageNumber(c)
		if err != nil {
			return
		}

		// Get the page size from the query string
		pageSize, err := getPageSize(c)
		if err != nil {
			return
		}

		// Get the profiles from the database
		profiles, err := fetchProfiles(profileRepo, c, page, pageSize)
		if err != nil {
			return
		}

		c.JSON(http.StatusOK, profiles)

	}
}
