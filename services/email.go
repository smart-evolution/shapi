package services

import (
    "os"
    "log"
    "net/smtp"
)

func SendEmail(body string) {
    from := os.Getenv("EMAILNAME")
    pass := os.Getenv("EMAILPASS")
    to := os.Getenv("EMAILNAME")
    smtpPort := os.Getenv("SMTPPORT")
    smtpAuthUrl := os.Getenv("SMTPAUTHURL")

    msg := "From: " + from + "\n" +
        "To: " + to + "\n" +
        "Subject: Home alert\n\n" +
        body

    smtpAuth := smtp.PlainAuth("", from, pass, smtpAuthUrl)

    err := smtp.SendMail(smtpPort, smtpAuth, from, []string{to}, []byte(msg))

    if err != nil {
        log.Printf("services: ", err)
        return
    }

    log.Print("services: alert sent")
}
