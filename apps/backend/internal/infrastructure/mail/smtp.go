package mail

import (
	"fmt"
	"net/smtp"
	"strings"

	"api/config"
)

type SMTPMailSender struct {
	host     string
	port     int
	username string
	password string
	from     string
	fromName string
}

func NewSMTPMailSender(cfg config.SMTPConfig) *SMTPMailSender {
	return &SMTPMailSender{
		host:     cfg.Host,
		port:     cfg.Port,
		username: cfg.Username,
		password: cfg.Password,
		from:     cfg.From,
		fromName: cfg.FromName,
	}
}

func (s *SMTPMailSender) Send(to string, subject string, htmlBody string) error {
	addr := fmt.Sprintf("%s:%d", s.host, s.port)
	auth := smtp.PlainAuth("", s.username, s.password, s.host)

	from := s.from
	if s.fromName != "" {
		from = fmt.Sprintf("%s <%s>", s.fromName, s.from)
	}

	headers := []string{
		"MIME-Version: 1.0",
		"Content-Type: text/html; charset=UTF-8",
		fmt.Sprintf("From: %s", from),
		fmt.Sprintf("To: %s", to),
		fmt.Sprintf("Subject: %s", subject),
	}

	message := []byte(strings.Join(headers, "\r\n") + "\r\n\r\n" + htmlBody)

	if err := smtp.SendMail(addr, auth, s.from, []string{to}, message); err != nil {
		return fmt.Errorf("smtp send: %w", err)
	}

	return nil
}
