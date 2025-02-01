package main

import (
	"bytes"
	"encoding/csv"
	"flag"
	"fmt"
	"html/template"
	"io"
	"log"
	"os"
	"strings"
)

const (
	NAME           = iota
	BOOK_AGAIN     = iota
	EMAIL_ADDDRESS = iota
)

type Contact struct {
	Name      string
	BookAgain bool
	Email     string
}

type SmtpSettings struct {
	userId   string
	password string
}

var contactsCsvFile string
var emailTemplate string

var smtpSettings *SmtpSettings

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {
	parseArgs()

	tmpl, err := template.ParseFiles(emailTemplate)
	check(err)

	fileData, err := os.ReadFile(contactsCsvFile)
	check(err)

	stringData := string(fileData)

	csvReader := csv.NewReader(strings.NewReader(stringData))
	csvReader.FieldsPerRecord = 3

	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		check(err)

		// skip the header
		if record[0] == "Name" {
			continue
		}

		contact := toContact(record)

		if contact.BookAgain {
			sendEmail(renderEmailBody(tmpl, contact), contact, smtpSettings)
		} else {
			fmt.Println("Skipping: ", contact.Name)
		}
	}

}

func renderEmailBody(template *template.Template, contact *Contact) string {
	var bytes bytes.Buffer

	err := template.Execute(&bytes, contact)
	check(err)

	return bytes.String()
}

func toContact(record []string) *Contact {
	bookAgain := true
	if record[1] == "NO" {
		bookAgain = false
	}

	return &Contact{record[0], bookAgain, record[2]}
}

func parseArgs() {
	var smtpUserId string
	var smtpPassword string

	flag.StringVar(&contactsCsvFile, "contacts", "", "CSV file holding contact details")
	flag.StringVar(&emailTemplate, "template", "", "Template to use for generated emails")
	flag.StringVar(&smtpUserId, "smtp-user-id", "", "User ID for the SMTP server")
	flag.StringVar(&smtpPassword, "smtp-password", "", "Password for the SMTP server")
	flag.Parse()

	if contactsCsvFile == "" ||
		emailTemplate == "" ||
		smtpUserId == "" ||
		smtpPassword == "" {

		flag.Usage()
		os.Exit(1)
	}

	smtpSettings = &SmtpSettings{smtpUserId, smtpPassword}
}
