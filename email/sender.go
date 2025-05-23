package main

import (
	"fmt"
	"net/smtp"
)

const smtpServer = "smtp.gmail.com"
const smtpTlsPort = "587"
const from = "glamorgan9quaywest@gmail.com"

func sendEmail(emailBody string, contact *Contact, smtpSettings *SmtpSettings, dryRun bool) {
	fmt.Println("Sending to: ", contact.Email)
	if !dryRun {
		auth := smtp.PlainAuth("", smtpSettings.userId, smtpSettings.password, smtpServer)
		err := smtp.SendMail(smtpServer+":"+smtpTlsPort, auth, from, []string{contact.Email}, []byte(emailBody))
		check(err)
	} else {
		fmt.Println("Dry run, skipping email send")
	}
}
