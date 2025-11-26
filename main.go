package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	// Check if note is provided
	if len(os.Args) < 2 {
		fmt.Println("Usage: autonote \"your note here\"")
		return
	}

	// Combine arguments into one note
	note := strings.Join(os.Args[1:], " ")

	// Get current timestamp
	timestamp := time.Now().Format("2025-01-02 15:04")

	// Format note (Markdown style)
	entry := fmt.Sprintf("- **%s**: %s\n", timestamp, note)

	// Open or create notes.md file
	file, err := os.OpenFile("notes.md", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error writing note:", err)
		return
	}
	defer file.Close()

	// Write note to file
	_, err = file.WriteString(entry)
	if err != nil {
		fmt.Println("Error saving note:", err)
		return
	}

	fmt.Println("✅ Note saved")

}
