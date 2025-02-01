package main

import (
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

var contactsCsvFile string
var emailTemplate string

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
			err = tmpl.Execute(os.Stdout, &contact)
			check(err)
		} else {
			fmt.Println("Skipping: ", contact.Name)
		}
	}

}

func toContact(record []string) *Contact {
	bookAgain := true
	if record[1] == "NO" {
		bookAgain = false
	}

	return &Contact{record[0], bookAgain, record[2]}
}

func parseArgs() {
	flag.StringVar(&contactsCsvFile, "contacts", "", "CSV file holding contact details")
	flag.StringVar(&emailTemplate, "template", "", "Template to use for generated emails")
	flag.Parse()

	if contactsCsvFile == "" || emailTemplate == "" {
		flag.Usage()
		os.Exit(1)
	}
}
