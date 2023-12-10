package routers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetPageNumber(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("valid page number", func(t *testing.T) {
		r := gin.Default()
		r.GET("/test", func(c *gin.Context) {
			page, err := getPageNumber(c)
			assert.NoError(t, err)
			assert.Equal(t, 1, page)
		})
		req, _ := http.NewRequest(http.MethodGet, "/test?page=1", nil)
		resp := httptest.NewRecorder()
		r.ServeHTTP(resp, req)
	})

	t.Run("no page number", func(t *testing.T) {
		r := gin.Default()
		r.GET("/test", func(c *gin.Context) {
			_, err := getPageNumber(c)
			assert.Error(t, err)
		})
		req, _ := http.NewRequest(http.MethodGet, "/test", nil)
		resp := httptest.NewRecorder()
		r.ServeHTTP(resp, req)
	})

	t.Run("multiple page numbers", func(t *testing.T) {
		r := gin.Default()
		r.GET("/test", func(c *gin.Context) {
			page, err := getPageNumber(c)
			assert.NoError(t, err)
			assert.Equal(t, 1, page)
		})
		req, _ := http.NewRequest(http.MethodGet, "/test?page=1&page=2", nil)
		resp := httptest.NewRecorder()
		r.ServeHTTP(resp, req)
	})
}

func TestGetPageSize(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("valid page size", func(t *testing.T) {
		r := gin.Default()
		r.GET("/test", func(c *gin.Context) {
			pageSize, err := getPageSize(c)
			assert.NoError(t, err)
			assert.Equal(t, 10, pageSize)
		})
		req, _ := http.NewRequest(http.MethodGet, "/test?pageSize=10", nil)
		resp := httptest.NewRecorder()
		r.ServeHTTP(resp, req)
	})

	t.Run("no page size", func(t *testing.T) {
		r := gin.Default()
		r.GET("/test", func(c *gin.Context) {
			_, err := getPageSize(c)
			assert.Error(t, err)
		})
		req, _ := http.NewRequest(http.MethodGet, "/test", nil)
		resp := httptest.NewRecorder()
		r.ServeHTTP(resp, req)
	})

	t.Run("multiple page sizes", func(t *testing.T) {
		r := gin.Default()
		r.GET("/test", func(c *gin.Context) {
			pageSize, err := getPageSize(c)
			assert.NoError(t, err)
			assert.Equal(t, 10, pageSize)
		})
		req, _ := http.NewRequest(http.MethodGet, "/test?pageSize=10&pageSize=20", nil)
		resp := httptest.NewRecorder()
		r.ServeHTTP(resp, req)
	})
}
