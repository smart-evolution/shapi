package services

import (
    "os"
    "log"
    "net/smtp"
)

func composeMessage(from string, to string, body string) string {
   return "From: " + from + "\n" +
    "To: " + to + "\n" +
    "Subject: Home alert\n\n" +
    body
}

// SendEmail - send email to subscriber
func SendEmail(body string) {
    sender := os.Getenv("EMAILNAME")
    pass := os.Getenv("EMAILPASS")
    recipient := os.Getenv("EMAILNAME")
    smtpPort := os.Getenv("SMTPPORT")
    smtpAuthURL := os.Getenv("SMTPAUTHURL")

    msg := composeMessage(sender, recipient, body)
    smtpAuth := smtp.PlainAuth("", sender, pass, smtpAuthURL)

    err := smtp.SendMail(smtpPort, smtpAuth, sender, []string{recipient}, []byte(msg))

    if err != nil {
        log.Println("services: ", err)
        return
    }

    log.Println("services: alert sent")
}
