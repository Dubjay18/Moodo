package server

import (
	"moodo-server/internal/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()
	apiVersion := r.Group("/api/v1")
	protected := apiVersion.Group("/")
	protected.Use(middleware.AuthMiddleware(s.authClient))
	protected.GET("/", s.HelloWorldHandler)
	movies := protected.Group("/movies")
	{
		movies.GET("/mood/:mood", s.GetMoviesByMoodHandler)
	}
	r.GET("/health", s.healthHandler)

	return r
}

func (s *Server) HelloWorldHandler(c *gin.Context) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	c.JSON(http.StatusOK, resp)
}

func (s *Server) healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, s.db.Health())
}
