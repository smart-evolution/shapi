package services

import (
    "os"
    "log"
    "net/smtp"
    "github.com/smart-evolution/smarthome/utils"
    "github.com/smart-evolution/smarthome/models"
    "gopkg.in/mgo.v2/bson"
)

func composeMessage(from string, to string, body string) string {
   return "From: " + from + "\n" +
    "To: " + to + "\n" +
    "Subject: Home alert\n\n" +
    body
}

// SendEmail - send email to subscriber
func SendEmail(body string, recipient string) {
    sender := os.Getenv("EMAILNAME")
    pass := os.Getenv("EMAILPASS")
    smtpPort := os.Getenv("SMTPPORT")
    smtpAuthURL := os.Getenv("SMTPAUTHURL")

    msg := composeMessage(sender, recipient, body)
    smtpAuth := smtp.PlainAuth("", sender, pass, smtpAuthURL)

    err := smtp.SendMail(smtpPort, smtpAuth, sender, []string{recipient}, []byte(msg))

    if err != nil {
        log.Println("services: error sending email to " + recipient, err)
        return
    }

    log.Println("services: alert sent to " + recipient)
}

// BulkEmail - sends alerts to all home users
func BulkEmail(body string) {
    ds := utils.GetDataSource()
    c := ds.C("users")

    var users []models.User

    err := c.Find(bson.M{}).All(&users)

    if err != nil {
        log.Println("services: Alert recipients not found", err)
    }

    for _, u := range users {
        SendEmail(body, u.Username)
    }
}
