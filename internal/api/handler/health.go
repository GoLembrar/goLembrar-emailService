package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealthChecker(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
	})
}
