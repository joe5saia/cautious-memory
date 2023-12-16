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

// GetSubfield handles the GET /subfield/:id endpoint
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

// GetAllSubfields handles the GET /subfields endpoint
func (h *SubfieldHandler) GetAllSubfields(c *gin.Context) {
	subfields, err := h.store.GetSubfields()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	c.JSON(http.StatusOK, subfields)
}

// CreateSubfield handles the POST /subfield endpoint
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

// UpdateSubfield handles the POST /subfield/:id endpoint
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

// DeleteSubfield handles the DELETE /subfield/:id endpoint
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
