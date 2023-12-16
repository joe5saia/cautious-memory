package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/joe5saia/cautious-memory/internal/model"
	"github.com/joe5saia/cautious-memory/internal/store"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

type ProfileHandler struct {
	store *store.Store
}

func NewProfileHandler(store *store.Store) *ProfileHandler {
	return &ProfileHandler{store: store}
}

// GetProfile godoc
//
//	@Summary		Get a profile
//	@Description	Get a profile by ID
//	@Tags			profiles
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Profile ID"
//	@Success		200	{object}	model.Profile
//	@Failure		400	{object}	error
//	@Failure		404	{object}	error
//	@Failure		500	{object}	error
//	@Router			/profiles/{id} [get]
func (h *ProfileHandler) GetProfile(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	profile, err := h.store.GetProfile(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Profile not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}

	c.JSON(http.StatusOK, profile)
}

// GetAllProfiles godoc
//
//	@Summary		Get all profiles
//	@Description	Get all profiles
//	@Tags			profiles
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		[]model.Profile
//	@Failure		500	{object}	error
//	@Router			/profiles [get]
func (h *ProfileHandler) GetAllProfiles(c *gin.Context) {
	profiles, err := h.store.GetProfiles()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	c.JSON(http.StatusOK, profiles)
}

// CreateProfile godoc
//
//	@Summary		Create a profile
//	@Description	Create a profile
//	@Tags			profiles
//	@Accept			json
//	@Produce		json
//	@Param			profile	body		model.Profile	true	"Profile"
//	@Success		201		{object}	string
//	@Failure		400		{object}	error
//	@Failure		500		{object}	error
//	@Router			/profiles [post]
func (h *ProfileHandler) CreateProfile(c *gin.Context) {
	var profile model.Profile
	if err := c.ShouldBindJSON(&profile); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.store.CreateProfile(&profile); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusCreated, profile)
}

// UpdateProfile godoc
//
//	@Summary		Update a profile
//	@Description	Update a profile
//	@Tags			profiles
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int				true	"Profile ID"
//	@Param			profile	body		model.Profile	true	"Profile"
//	@Success		200		{object}	string
//	@Failure		400		{object}	error
//	@Failure		500		{object}	error
//	@Router			/profiles/{id} [post]
func (h *ProfileHandler) UpdateProfile(c *gin.Context) {
	var updatedProfile model.Profile
	if err := c.ShouldBindJSON(&updatedProfile); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	if err := h.store.UpdateProfile(uint(id), &updatedProfile); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Profile updated successfully"})
}

// DeleteProfile godoc
//
//	@Summary		Delete a profile
//	@Description	Delete a profile
//	@Tags			profiles
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Profile ID"
//	@Success		200	{object}	string
//	@Failure		400	{object}	error
//	@Failure		500	{object}	error
//	@Router			/profiles/{id} [delete]
func (h *ProfileHandler) DeleteProfile(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	if err := h.store.DeleteProfile(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Profile deleted successfully"})
}
