package services

import (
	"fmt"
	em "github.com/carlosclavijo/Pinterest-Services/internal/domain/email"
	"github.com/jordan-wright/email"
	"net/smtp"
	"net/url"
)

type EmailService struct {
	Host     string
	Port     string
	Username string
	Password string
	From     string
	AppUrl   string
}

func (e *EmailService) SendVerificationEmail(toEmail, token string) error {
	// URL-encode token for safety
	escapedToken := url.QueryEscape(token)
	link := fmt.Sprintf("%s/verify-email?token=%s", e.AppUrl, escapedToken)

	// Create email using jordan-wright/email
	em := email.NewEmail()
	em.From = fmt.Sprintf("%s <%s>", "Pinterest-Clone", e.Username) // must match Gmail account
	em.To = []string{toEmail}
	em.Subject = "Verify your email"
	em.HTML = []byte(fmt.Sprintf("<p>Click <a href='%s'>here</a> to verify your email.</p>", link))

	// SMTP authentication
	auth := smtp.PlainAuth("", e.Username, e.Password, e.Host)
	addr := fmt.Sprintf("%s:%s", e.Host, e.Port)

	// Send asynchronously
	go func() {
		if err := em.Send(addr, auth); err != nil {
			fmt.Printf("failed to send verification email to %s: %v\n", toEmail, err)
		}
	}()

	return nil
}

var _ em.Sender = (*EmailService)(nil)
