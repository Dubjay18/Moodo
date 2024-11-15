package server

import (
	"moodo-server/internal/lib"
	"moodo-server/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

const (
	GetMoviesSuccessMessage = "Successfully fetched movies"
)

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

	movies, _, _, err := services.RecommendMoviesByMood(mood, page)
	if err != nil {
		lib.BuildErrorResponse(c, err, err.Error(), http.StatusInternalServerError)
		return
	}
	lib.BuildResponse[[]services.Movie](c, movies, GetMoviesSuccessMessage, http.StatusOK)
}
