package email

import (
	"github.com/smart-evolution/smarthome/utils"
	"net/smtp"
)

// IMailer - interface for mailer
type IMailer interface {
	AddRecipient(string)
	SendEmail(string, string)
	BulkEmail(string)
}

// Mailer - email notifier
type Mailer struct {
	Sender      string
	Password    string
	SMTPPort    string
	SMTPAuthURL string
	recipients  []string
}

// New - creates new instance of Mailer
func New(recipients []string, sender string, password string, SMTPPort string, SMTPAuthURL string) *Mailer {
	return &Mailer{
		Sender:      sender,
		Password:    password,
		SMTPPort:    SMTPPort,
		SMTPAuthURL: SMTPAuthURL,
		recipients:  recipients,
	}
}

// AddRecipient - adds new recipient of mailer
func (m *Mailer) AddRecipient(email string) {
	m.recipients = append(m.recipients, email)
}

func composeMessage(from string, to string, body string) string {
	return "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: Home alert\n\n" +
		body
}

// SendEmail - send email to subscriber
func (m *Mailer) SendEmail(body string, recipient string) {
	msg := composeMessage(m.Sender, recipient, body)
	smtpAuth := smtp.PlainAuth("", m.Sender, m.Password, m.SMTPAuthURL)

	err := smtp.SendMail(m.SMTPPort, smtpAuth, m.Sender, []string{recipient}, []byte(msg))

	if err != nil {
		utils.Log("error sending email to "+recipient, err)
		return
	}

	utils.Log("alert sent to " + recipient)
}

// BulkEmail - sends alerts to all home users
func (m *Mailer) BulkEmail(body string) {
	for _, r := range m.recipients {
		m.SendEmail(body, r)
	}
}
