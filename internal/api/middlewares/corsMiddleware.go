package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func CorsMiddleware() gin.HandlerFunc {

	allowedOrigins := []string{"https://api.golembrar.com"}

	excludeEndPoints := []string{"/check"}

	return func(c *gin.Context) {

		origin := c.Request.Header.Get("Origin")
		path := c.Request.URL.Path

		fmt.Println(path)

		isOriginAllowed := func(origin string, allowedOrigins []string) bool {
			for _, allowedOrigin := range allowedOrigins {
				for _, excludeEndPoint := range excludeEndPoints {
					if origin != allowedOrigin && path != excludeEndPoint {
						c.AbortWithStatus(403)
					}

				}

			}
			return false
		}

		if isOriginAllowed(origin, allowedOrigins) {
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
			c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
			c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
			c.Writer.Header().Set("Access-Control-Allow-Methods", "GET POST PUT PATCH DELETE HEAD OPTIONS")
		}

		c.Next()
	}
}
