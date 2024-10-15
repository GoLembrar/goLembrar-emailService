package email

import (
	"fmt"

	"github.com/GoLembrar/goLembrar-emailService/internal/utils"
	"github.com/go-playground/validator/v10"
	"github.com/resend/resend-go/v2"
)

type EmailParams struct {
	Emails     []string `json:"emails" validate:"required,dive,email,max=100" example:"joaozinho@gmail.com"`
	Title      string   `json:"title" validate:"required,max=120" example:"The message title"`
	DescHtml   string   `json:"descHtml" validate:"required,max=500" example:"<p>Message description</p>"`
	OwnerEmail string   `json:"ownerEmail" validate:"required,email,max=100" example:"maria@gmail.com"`
}

type ScheduleEmailParams struct {
	EmailParams
	ScheduledAt string `json:"scheduledAt" validate:"required,datetime"`
}

type EmailService struct {
	client   *resend.Client
	validate *validator.Validate
}

func NewEmailService() (*EmailService, error) {
	apiKey := utils.GetEnvVar("APIKEY_RESEND")
	client := resend.NewClient(apiKey)

	validate := validator.New()

	return &EmailService{
		client:   client,
		validate: validate,
	}, nil
}

func (s *EmailService) SendEmail(params *EmailParams) (string, error) {

	if err := s.validate.Struct(params); err != nil {
		validationErrs := err.(validator.ValidationErrors)
		return "", fmt.Errorf("validation error: %v", validationErrs)
	}

	resendParams := &resend.SendEmailRequest{
		From:    utils.GetEnvVar("SEND_EMAIL"),
		To:      params.Emails,
		Subject: params.Title,
		Html:    params.DescHtml,
		ReplyTo: params.OwnerEmail,
	}

	sent, err := s.client.Emails.Send(resendParams)
	if err != nil {
		return "", fmt.Errorf("failed to send email: %w", err)
	}

	return sent.Id, nil
}

func (s *EmailService) ScheduleEmail(params *ScheduleEmailParams) (string, error) {
	if err := s.validate.Struct(params); err != nil {
		validationErrs := err.(validator.ValidationErrors)
		return "", fmt.Errorf("validation error: %v", validationErrs)
	}

	resendParams := &resend.SendEmailRequest{
		From:        utils.GetEnvVar("SEND_EMAIL"),
		ScheduledAt: params.ScheduledAt,
		To:          params.Emails,
		Subject:     params.Title,
		Html:        params.DescHtml,
		ReplyTo:     params.OwnerEmail,
	}

	sent, err := s.client.Emails.Send(resendParams)
	if err != nil {
		return "", fmt.Errorf("failed to send email: %w", err)
	}

	return sent.Id, nil
}
