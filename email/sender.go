package main

import "fmt"

func sendEmail(emailBody string, contact *Contact) {
	fmt.Println(contact.Email)
	fmt.Println(emailBody)
}