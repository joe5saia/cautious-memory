package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joe5saia/cautious-memory/internal/model"
	"github.com/joe5saia/cautious-memory/internal/store"
)

type SubfieldHandler struct {
	store *store.Store
}

func NewSubfieldHandler(store *store.Store) *SubfieldHandler {
	return &SubfieldHandler{store: store}
}

// GetSubfield godoc
//
//	@Summary		Get a subfield
//	@Description	Get a subfield by ID
//	@Tags			subfields
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Subfield ID"
//	@Success		200	{object}	model.Subfield
//	@Failure		400	{object}	error
//	@Failure		404	{object}	error
//	@Failure		500	{object}	error
//	@Router			/subfields/{id} [get]
func (h *SubfieldHandler) GetSubfield(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	subfield, err := h.store.GetSubfield(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, subfield)
}

// GetAllSubfields godoc
//
//	@Summary		Get all subfields
//	@Description	Get all subfields
//	@Tags			subfields
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		model.Subfield
//	@Failure		500	{object}	error
//	@Router			/subfields [get]
func (h *SubfieldHandler) GetAllSubfields(c *gin.Context) {
	subfields, err := h.store.GetSubfields()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	c.JSON(http.StatusOK, subfields)
}

// CreateSubfield godoc
//
//	@Summary		Create a subfield
//	@Description	Create a subfield
//	@Tags			subfields
//	@Accept			json
//	@Produce		json
//	@Param			subfield	body		model.Subfield	true	"Subfield"
//	@Success		201			{object}	model.Subfield
//	@Failure		400			{object}	error
//	@Failure		500			{object}	error
//	@Router			/subfields [post]
func (h *SubfieldHandler) CreateSubfield(c *gin.Context) {
	var subfield model.Subfield
	if err := c.ShouldBindJSON(&subfield); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.store.CreateSubfield(&subfield); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusCreated, subfield)
}

// UpdateSubfield godoc
//
//	@Summary		Update a subfield
//	@Description	Update a subfield
//	@Tags			subfields
//	@Accept			json
//	@Produce		json
//	@Param			id			path		int				true	"Subfield ID"
//	@Param			subfield	body		model.Subfield	true	"Subfield"
//	@Success		200			{object}	string
//	@Failure		400			{object}	error
//	@Failure		500			{object}	error
//	@Router			/subfields/{id} [post]
func (h *SubfieldHandler) UpdateSubfield(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var updatedSubfield model.Subfield
	if err := c.ShouldBindJSON(&updatedSubfield); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.store.UpdateSubfield(uint(id), &updatedSubfield); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Subfield updated successfully"})
}

// DeleteSubfield godoc
//
//	@Summary		Delete a subfield
//	@Description	Delete a subfield
//	@Tags			subfields
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Subfield ID"
//	@Success		200	{object}	string
//	@Failure		400	{object}	error
//	@Failure		500	{object}	error
//	@Router			/subfields/{id} [delete]
func (h *SubfieldHandler) DeleteSubfield(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	if err := h.store.DeleteSubfield(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Subfield deleted successfully"})
}
