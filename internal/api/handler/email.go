package handler

import (
	"net/http"

	"github.com/GoLembrar/goLembrar-emailService/internal/email"
	"github.com/gin-gonic/gin"
)

type EmailHandler struct {
	emailService *email.EmailService
}

func NewEmailHandler(emailService *email.EmailService) *EmailHandler {
	return &EmailHandler{emailService: emailService}
}

func (h *EmailHandler) SendEmail(c *gin.Context) {
	var params email.EmailParams

	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := h.emailService.SendEmail(&params)
	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Email sent successfully", "id": id})
}
