package middleware

import (
	"context"
	"firebase.google.com/go/auth"
	_ "firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
	"net/http"
)

// AuthMiddleware is a middleware function for the Gin framework that handles
// authentication using an auth.Client. It checks for the presence of an
// Authorization token in the request header and verifies its validity.
//
// If the token is missing or invalid, the middleware aborts the request and
// responds with an HTTP 401 Unauthorized status. If the token is valid, the
// middleware allows the request to proceed to the next handler.
//
// Parameters:
// - authClient: A pointer to an auth.Client used to verify the ID token.
//
// Returns:
// - A gin.HandlerFunc that can be used as middleware in a Gin router.
func AuthMiddleware(authClient *auth.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, "Authorization token missing")
			return
		}

		_, err := authClient.VerifyIDToken(context.Background(), token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, "Invalid token")
			return
		}
		c.Next()
	}
}
