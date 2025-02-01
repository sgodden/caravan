package main

import (
	"encoding/csv"
	"flag"
	"fmt"
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
	Name string
	Email string
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

	fileData, err := os.ReadFile("contacts.csv")
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

		name := record[0]
		fmt.Println(name)

		// for idx, element := range record {
		// 	if idx == EMAIL_ADDDRESS && len(element) > 0 {
		// 		fmt.Println(element)
		// 	}
		// }
	}

}

func parseArgs() {
	flag.StringVar(&contactsCsvFile, "contacts", "", "CSV file holding contact details")
	flag.StringVar(&emailTemplate, "template", "", "Template to use for generated emails")
	flag.Parse()

	if contactsCsvFile == "" || emailTemplate == "" {
		flag.Usage()
		os.Exit(1)
	}

	fmt.Println("contacts: ", contactsCsvFile)
	fmt.Println("template: ", emailTemplate)
}

func renderTemplate() {

}
