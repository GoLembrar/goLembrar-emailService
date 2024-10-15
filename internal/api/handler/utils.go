package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary		Health checker
// @Description	check if API is running
// @Tags			main
// @Success 200 {object} HealthCheckerResponse
// @Router			/check	[get]
func HealthChecker(c *gin.Context) {
	response := HealthCheckerResponse{Status: "ok"}
	c.JSON(200, response)
}

type HealthCheckerResponse struct {
	Status string `json:"status" example:"ok"`
}

func RedirectToDocs(c *gin.Context) {
	c.Redirect(http.StatusFound, "/docs/index.html")
}
