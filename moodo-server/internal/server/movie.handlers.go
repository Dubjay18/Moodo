package server

import (
	"moodo-server/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Gin handler for mood-based movie recommendations with pagination
func (s *Server) GetMoviesByMoodHandler(c *gin.Context) {
	mood := c.Param("mood")
	pageQuery := c.Query("page")

	// Default to the first page if no page query parameter is specified
	page := 1
	if pageQuery != "" {
		var err error
		page, err = strconv.Atoi(pageQuery)
		if err != nil || page < 1 {
			page = 1
		}
	}

	movies, currentPage, totalPages, err := services.RecommendMoviesByMood(mood, page)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"movies":      movies,
		"page":        currentPage,
		"total_pages": totalPages,
	})
}
