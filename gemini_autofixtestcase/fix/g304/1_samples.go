package main

import (
	"log"
	"os"
	"path/filepath"
)

func main() {
	// Predefined directory
	predefinedDir := "/var/safe_directory"

	f, err := sanitizeInput(os.Getenv("tainted_file"))
	if err != nil {
		log.Fatalf("Invalid tainted_file: %v", err)
	}

	// Construct the file path using the predefined directory
	filePath := filepath.Join(predefinedDir, f)

	// Normalize the file path to prevent directory traversal
	cleanFilePath := filepath.Clean(filePath)
	if !filepath.HasPrefix(cleanFilePath, predefinedDir) {
		log.Fatalf("Invalid file path: %s", cleanFilePath)
	}

	body, err := os.ReadFile(cleanFilePath)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}
	log.Print(body)

}
