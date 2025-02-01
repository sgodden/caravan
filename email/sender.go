package main

import "fmt"

const smtpServer = "smtp.gmail.com"
const smtpTlsPort = 587

func sendEmail(emailBody string, contact *Contact, smtpSettings *SmtpSettings) {
	fmt.Println(contact.Email)
	fmt.Println(emailBody)

	fmt.Println("Settings: ", smtpSettings)
}
