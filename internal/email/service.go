package email

import (
	"fmt"

	"github.com/GoLembrar/goLembrar-emailService/internal/utils"
	"github.com/go-playground/validator/v10"
	"github.com/resend/resend-go/v2"
)

type EmailParams struct {
	To      []string `json:"to" validate:"required,dive,email,max=100"`
	Html    string   `json:"html" validate:"required,max=500"`
	Subject string   `json:"subject" validate:"required,max=120"`
	Cc      []string `json:"cc" validate:"omitempty,dive,email,max=100"`
	Bcc     []string `json:"bcc" validate:"omitempty,dive,email,max=100"`
	ReplyTo string   `json:"replyTo" validate:"omitempty,email,max=100"`
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
		To:      params.To,
		Html:    params.Html,
		Subject: params.Subject,
		Cc:      params.Cc,
		Bcc:     params.Bcc,
		ReplyTo: utils.GetEnvVar("SEND_EMAIL"),
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
		To:          params.To,
		Html:        params.Html,
		Subject:     params.Subject,
		Cc:          params.Cc,
		Bcc:         params.Bcc,
		ReplyTo:     utils.GetEnvVar("SEND_EMAIL"),
		ScheduledAt: params.ScheduledAt,
	}

	sent, err := s.client.Emails.Send(resendParams)
	if err != nil {
		return "", fmt.Errorf("failed to send email: %w", err)
	}

	return sent.Id, nil
}
