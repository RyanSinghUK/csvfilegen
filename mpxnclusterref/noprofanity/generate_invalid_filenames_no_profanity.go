package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// Define invalid characters for filenames
const (
	invalidSpecialChars = `<>:"/\|?*` + "`" + `~!@#$%^&+=[]{};'`
	whitespace          = " "
)

func main() {
	rand.Seed(time.Now().UnixNano())

	var numEntries int
	fmt.Print("Enter the number of invalid filename entries you want to generate: ")
	fmt.Scan(&numEntries)

	entries := generateInvalidFilenames(numEntries)
	for _, entry := range entries {
		fmt.Println(entry)
	}
}

func generateInvalidFilenames(num int) []string {
	entries := make([]string, num)
	for i := 0; i < num; i++ {
		entries[i] = generateInvalidFilename()
	}
	return entries
}

func generateInvalidFilename() string {
	choices := []func() string{
		generateFilenameWithInvalidChars,
		generateFilenameTooLong,
		generateFilenameTooShort,
		generateFilenameWithLeadingWhitespace,
		generateFilenameWithTrailingWhitespace,
		generateFilenameWithWhitespaceBetween,
	}
	return choices[rand.Intn(len(choices))]()
}

func generateFilenameWithInvalidChars() string {
	length := rand.Intn(36) + 1
	var sb strings.Builder
	for i := 0; i < length; i++ {
		sb.WriteByte(invalidSpecialChars[rand.Intn(len(invalidSpecialChars))])
	}
	return sb.String()
}

func generateFilenameTooLong() string {
	length := rand.Intn(10) + 37
	return generateRandomString(length)
}

func generateFilenameTooShort() string {
	return ""
}

func generateFilenameWithLeadingWhitespace() string {
	return whitespace + generateRandomString(rand.Intn(35)+1)
}

func generateFilenameWithTrailingWhitespace() string {
	return generateRandomString(rand.Intn(35)+1) + whitespace
}

func generateFilenameWithWhitespaceBetween() string {
	length := rand.Intn(34) + 2
	return generateRandomString(length/2) + whitespace + generateRandomString(length/2)
}

func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var sb strings.Builder
	for i := 0; i < length; i++ {
		sb.WriteByte(charset[rand.Intn(len(charset))])
	}
	return sb.String()
}
