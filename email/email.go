package main

import (
	"encoding/csv"
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

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {
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

		for idx, element := range record {
			if idx == EMAIL_ADDDRESS && len(element) > 0 {
				fmt.Println(element)
			}
		}
	}

}
