package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"internal/schemas"
)

func postAlbums(c *gin.Context) {
	var newProfile schemas.Profile

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newProfile); err != nil {
		return
	}

	// Add the new album to the slice.
	albums = append(albums, newProfile)
	c.IndentedJSON(http.StatusCreated, newProfile)
}
