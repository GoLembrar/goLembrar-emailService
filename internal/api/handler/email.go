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

// @Summary		Send email
// @Description	take request and send the email for you
// @Tags			email
// @Param			Body	body	email.EmailParams	true	"Email parameters"
// @Produce		json
// @Success		200		{object} EmailResponse
// @Failure		400				{object}	swagger.APIError
// @Router			/send-email	[post]
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

type EmailResponse struct {
	ID      string `json:"id" example:"cfb2011b..."`
	Message string `json:"message" example:"Email sent successfully"`
}

func (h *EmailHandler) ScheduleEmail(c *gin.Context) {
	var params email.ScheduleEmailParams

	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := h.emailService.ScheduleEmail(&params)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Email schedule successfully", "id": id})
}
