package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joe5saia/cautious-memory/internal/model"
	"github.com/joe5saia/cautious-memory/internal/store"
)

type FieldHandler struct {
	store *store.Store
}

func NewFieldHandler(store *store.Store) *FieldHandler {
	return &FieldHandler{store: store}
}

// GetField godoc
//
//	@Summary		Get a field
//	@Description	Get a field by ID
//	@Tags			fields
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Field ID"
//	@Success		200	{object}	model.Field
//	@Failure		400	{object}	error
//	@Failure		404	{object}	error
//	@Failure		500	{object}	error
//	@Router			/fields/{id} [get]
func (h *FieldHandler) GetField(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	field, err := h.store.GetField(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, field)
}

// GetAllFields godoc
//
//	@Summary		Get all fields
//	@Description	Get all fields
//	@Tags			fields
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		model.Field
//	@Failure		500	{object}	error
//	@Router			/fields [get]
func (h *FieldHandler) GetAllFields(c *gin.Context) {
	fields, err := h.store.GetFields()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	c.JSON(http.StatusOK, fields)
}

// CreateField godoc
//
//	@Summary		Create a field
//	@Description	Create a field
//	@Tags			fields
//	@Accept			json
//	@Produce		json
//	@Param			field	body		model.Field	true	"Field"
//	@Success		201		{object}	model.Field
//	@Failure		400		{object}	error
//	@Failure		500		{object}	error
//	@Router			/fields [post]
func (h *FieldHandler) CreateField(c *gin.Context) {
	var field model.Field
	if err := c.ShouldBindJSON(&field); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.store.CreateField(&field); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusCreated, field)
}

// UpdateField godoc
//
//	@Summary		Update a field
//	@Description	Update a field
//	@Tags			fields
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int			true	"Field ID"
//	@Param			field	body		model.Field	true	"Field"
//	@Success		200		{object}	string
//	@Failure		400		{object}	error
//	@Failure		500		{object}	error
//	@Router			/fields/{id} [post]
func (h *FieldHandler) UpdateField(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var updatedField model.Field
	if err := c.ShouldBindJSON(&updatedField); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.store.UpdateField(uint(id), &updatedField); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Field updated successfully"})
}

// DeleteField godoc
//
//	@Summary		Delete a field
//	@Description	Delete a field
//	@Tags			fields
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Field ID"
//	@Success		200	{object}	string
//	@Failure		400	{object}	error
//	@Failure		500	{object}	error
//	@Router			/fields/{id} [delete]
func (h *FieldHandler) DeleteField(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	if err := h.store.DeleteField(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Field deleted successfully"})
}
