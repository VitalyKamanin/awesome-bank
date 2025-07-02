package integrations

import (
	"github.com/go-mail/mail/v2"
	"log"
)

const (
	smtpHost = "smtp.example.com"
	smtpPort = 587
	smtpUser = "noreply@example.com"
	smtpPass = "yourpassword"
)

func SendEmail(to, subject, body string) error {
	m := mail.NewMessage()
	m.SetHeader("From", smtpUser)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	d := mail.NewDialer(smtpHost, smtpPort, smtpUser, smtpPass)
	d.TLSConfig = nil

	if err := d.DialAndSend(m); err != nil {
		log.Printf("SMTP error: %v", err)
		return err
	}
	return nil
}
