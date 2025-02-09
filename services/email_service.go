package services

import (
    "log"
    "github.com/sendgrid/sendgrid-go"
    "github.com/sendgrid/sendgrid-go/helpers/mail"
)

type EmailService struct {
    client *sendgrid.Client
}

func NewEmailService(apiKey string) *EmailService {
    return &EmailService{
        client: sendgrid.NewSendClient(apiKey),
    }
}

func (s *EmailService) SendTaskNotification(to, subject, content string) error {
    from := mail.NewEmail("Task Manager", "rupesh.singh.8370412@gmail.com")
    recipient := mail.NewEmail("User", to)
    message := mail.NewSingleEmail(from, subject, recipient, content, content)
    
    response, err := s.client.Send(message)
    if err != nil {
        log.Printf("Failed to send email: %v\n", err)
        return err
    }

    // Log the SendGrid response
    log.Printf("Email sent to %s. Response: %v\n", to, response)
    return nil
}