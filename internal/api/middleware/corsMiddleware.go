package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CorsMiddleware() gin.HandlerFunc {
	allowedOrigins := map[string]struct{}{
		"https://api.golembrar.com": {},
	}

	excludeEndPoints := map[string]struct{}{
		"/check": {},
	}

	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")
		path := c.Request.URL.Path

		if _, ok := excludeEndPoints[path]; ok {
			c.Next()
			return
		}

		if _, ok := allowedOrigins[origin]; ok {
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
			c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
			c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
			c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, HEAD, OPTIONS")
		} else {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		c.Next()
	}
}
