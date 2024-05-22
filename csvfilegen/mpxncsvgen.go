package main

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	var numEntries int
	fmt.Print("Enter the number of entries you want to create: ")
	fmt.Scan(&numEntries)

	entries := generateEntries(numEntries)
	filename := generateFilenameWithTimestamp("output")
	writeCSV(filename, entries)
	fmt.Printf("CSV file '%s' created successfully.\n", filename)
}

func generateEntries(num int) []string {
	rand.Seed(time.Now().UnixNano())
	entries := make([]string, num)

	for i := 0; i < num; i++ {
		entries[i] = generate13DigitInt()
	}

	return entries
}

func generate13DigitInt() string {
	min := int64(1000000000000) // 13-digit minimum value (10^12)
	max := int64(9999999999999) // 13-digit maximum value (10^13 - 1)
	num := rand.Int63n(max-min+1) + min
	return strconv.FormatInt(num, 10)
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
