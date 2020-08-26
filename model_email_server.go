package wrpgomail

import (
	"strings"

	"github.com/pkg/errors"
	gomail "gopkg.in/gomail.v2"
)

// EmailServerConfig Concentrates Email Server configuration.
type EmailServerConfig struct {
	URI      string // URI for SMTP server
	Port     int    // Server Port
	User     string // Authentication - user name
	Password string // Authentication - password
	UseTLS   bool
}

// EmailServer Model for email server.
type EmailServer struct {
	EmailServerConfig
}

// NewEmailServer Constructor for email server
func NewEmailServer(cfg EmailServerConfig) (*EmailServer, error) {
	// validate config

	return &EmailServer{
		EmailServerConfig: cfg,
	}, nil
}

// EmailData Concentrates data needed to send an email.
type EmailData struct {
	Important   bool
	Subject     string
	MessageHTML string
	From        string
	To          []string
	CC          []string
	Embedded    []string // Contains paths for files to be embedded
	Attachments []string // Contains paths for files to be attached
}

// SendEmail Method sends email.
func (e *EmailServer) SendEmail(emData EmailData) error {
	// validate data

	m := gomail.NewMessage()
	m.SetHeader("From", emData.From)
	m.SetHeader("To", strings.Join(emData.To, ","))
	m.SetHeader("CC", strings.Join(emData.CC, ","))
	m.SetHeader("Subject", emData.Subject)
	m.SetBody("text/html", emData.MessageHTML)

	for _, a := range emData.Attachments {
		m.Attach(a)
	}

	for _, e := range emData.Embedded {
		m.Embed(e)
	}

	d := gomail.NewDialer(e.EmailServerConfig.URI, e.EmailServerConfig.Port, e.EmailServerConfig.User, e.EmailServerConfig.Password)
	errDial := d.DialAndSend(m)
	if errDial != nil {
		return errors.WithMessage(errDial, "could not dial and send")
	}
	return nil
}
