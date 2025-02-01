package main

import (
	"net/smtp"
)

const smtpServer = "smtp.gmail.com"
const smtpTlsPort = "587"
const from = "glamorgan9quaywest@gmail.com"

func sendEmail(emailBody string, contact *Contact, smtpSettings *SmtpSettings) {
	auth := smtp.PlainAuth("", smtpSettings.userId, smtpSettings.password, smtpServer)
	err := smtp.SendMail(smtpServer+":"+smtpTlsPort, auth, from, []string{contact.Email}, []byte(emailBody))
	check(err)
}
