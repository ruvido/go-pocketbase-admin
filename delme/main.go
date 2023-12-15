package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"
	"io"
)

func main() {
	file, err := os.Open("test.csv")
	if err != nil {
		log.Fatal("Error while reading the file", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	// Skip the header row
	_, err = reader.Read()
	if err != nil {
		log.Fatal(err)
	}

	for {
		// Read each record
		row, err := reader.Read()

		// Stop at EOF
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}

		// Assuming the date is in the second field (index 1)
		dateStr := row[1]
		// date, err := time.Parse("2006-01-02", dateStr)
		date, err := time.Parse("2006-01-02 15:04:05.999999999-07", dateStr)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(date)
	}
}
