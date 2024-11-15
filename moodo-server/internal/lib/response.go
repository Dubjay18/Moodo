package lib

import (
	"github.com/gin-gonic/gin"
)

func BuildResponse[T any](c *gin.Context, data T, message string, status int) {
	resp := make(map[string]interface{})
	resp["data"] = data
	resp["message"] = message
	resp["status"] = status
	// TODO: Add pagination

	c.JSON(status, resp)
}

func BuildErrorResponse(c *gin.Context, err error, message string, status int) {
	resp := make(map[string]interface{})
	resp["error"] = true
	resp["message"] = message
	resp["status"] = status
	resp["details"] = err.Error()

	c.AbortWithStatusJSON(status, resp)
}
