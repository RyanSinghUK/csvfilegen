package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"
)

func main() {
	entries := generateInvalidEntries()
	filename := generateFilenameWithTimestamp("invalid_entries")
	writeCSV(filename, entries)
	fmt.Printf("CSV file '%s' created successfully.\n", filename)
}

func generateInvalidEntries() []string {
	entries := []string{
		"",                            // Blank entry at the start
		"-1234567890123",              // Negative number
		"12345abcde67890",             // Alphanumeric entry
		"justastring",                 // String entry
		"1",                           // Minimum length integer
		"12345678901234567890",        // Maximum length integer
		"12345678901234",              // Entry greater than 13 digits
		time.Now().Format("20060102"), // Date format entry (YYYYMMDD)
		"null",                        // Null entry
		"",                            // Blank entry in the middle
		"9876543210987",               // Valid 13-digit integer for control
		"",                            // Blank entry at the end
	}

	return entries
}

func generateFilenameWithTimestamp(baseName string) string {
	timestamp := time.Now().Format("020106150405") // DDMMYYHHMMSS format
	return fmt.Sprintf("%s_%s.csv", baseName, timestamp)
}

func writeCSV(filename string, entries []string) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating CSV file:", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, entry := range entries {
		err := writer.Write([]string{entry})
		if err != nil {
			fmt.Println("Error writing to CSV file:", err)
			return
		}
	}
}
