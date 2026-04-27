package mail

import (
	"fmt"
	"log"
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
	var auth smtp.Auth
	if s.username != "" || s.password != "" {
		if s.username == "" || s.password == "" {
			return fmt.Errorf("smtp auth is incomplete")
		}
		auth = smtp.PlainAuth("", s.username, s.password, s.host)
	}
	if isGmailHost(s.host) && auth == nil {
		return fmt.Errorf("gmail smtp requires username and app password")
	}

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

	log.Printf("smtp: attempting to send %q email to %s via %s", subject, to, addr)
	if err := smtp.SendMail(addr, auth, s.from, []string{to}, message); err != nil {
		log.Printf("smtp: failed to send email to %s via %s: %v", to, addr, err)
		return fmt.Errorf("smtp send: %w", err)
	}

	log.Printf("smtp: sent %q email to %s via %s", subject, to, addr)
	return nil
}

func isGmailHost(host string) bool {
	host = strings.ToLower(strings.TrimSpace(host))
	return host == "smtp.gmail.com" || strings.HasSuffix(host, ".smtp.gmail.com")
}
