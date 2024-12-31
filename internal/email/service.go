package email

import (
	"fmt"
	"net/smtp"
)

type Service interface {
	SendEmail(to, subject, body string) error
}

type service struct {
	smtpHost     string
	smtpPort     string
	smtpUsername string
	smtpPassword string
}

func NewService(host, port, username, password string) Service {
	return &service{
		smtpHost:     host,
		smtpPort:     port,
		smtpUsername: username,
		smtpPassword: password,
	}
}

func (s *service) SendEmail(to, subject, body string) error {
	msg := fmt.Sprintf("To: %s\r\nSubject: %s\r\n\r\n%s", to, subject, body)
	addr := fmt.Sprintf("%s:%s", s.smtpHost, s.smtpPort)
	auth := smtp.PlainAuth("", s.smtpUsername, s.smtpPassword, s.smtpHost)

	return smtp.SendMail(addr, auth, s.smtpUsername, []string{to}, []byte(msg))
}
