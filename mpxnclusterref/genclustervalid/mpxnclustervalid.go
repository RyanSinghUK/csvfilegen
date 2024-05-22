package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// Define valid characters for filenames
const (
	letters      = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits       = "0123456789"
	specialChars = "-_."
	validChars   = letters + digits + specialChars
	maxLength    = 36
)

func main() {
	rand.Seed(time.Now().UnixNano())

	var numEntries int
	fmt.Print("Enter the number of filename entries you want to generate: ")
	fmt.Scan(&numEntries)

	entries := generateFilenames(numEntries)
	for _, entry := range entries {
		fmt.Println(entry)
	}
}

func generateFilenames(num int) []string {
	entries := make([]string, num)
	for i := 0; i < num; i++ {
		length := rand.Intn(maxLength) + 1
		entries[i] = generateFilename(length)

		// Randomly decide whether to append the date
		if rand.Intn(2) == 0 {
			entries[i] += "_" + time.Now().Format("02012006") // DDMMYYYY
		}
	}
	return entries
}

func generateFilename(length int) string {
	var sb strings.Builder
	for i := 0; i < length; i++ {
		sb.WriteByte(validChars[rand.Intn(len(validChars))])
	}
	return sb.String()
}
