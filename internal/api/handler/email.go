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

// @Summary			Send email
// @Description	Take the info and send email for you
// @Tags				email
// @Param			Body body email.EmailParams true "Email parameters"
// @Produce			json
// @Success			200
// @Failure			400			{object}	swagger.APIError
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
