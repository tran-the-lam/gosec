package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

func sanitizeInput(input string) (string, error) {
	// Define a regular expression to match only safe characters (e.g., alphanumeric and limited special characters)
	re := regexp.MustCompile(`^[a-zA-Z0-9_\-]+$`)
	if !re.MatchString(input) {
		return "", fmt.Errorf("invalid input: %s", input)
	}
	return input, nil
}

func main() {
	dir, err := sanitizeInput(os.Getenv("server_root"))
	if err != nil {
		log.Fatalf("Invalid server_root: %v", err)
	}

	f3, err := sanitizeInput(os.Getenv("tainted_file3"))
	if err != nil {
		log.Fatalf("Invalid tainted_file3: %v", err)
	}

	// Define a whitelist of allowed directories
	allowedDirs := map[string]bool{
		"allowed_dir1": true,
		"allowed_dir2": true,
	}

	if !allowedDirs[dir] {
		log.Fatalf("Directory not allowed: %s", dir)
	}

	// Normalize and validate the file path
	filePath := filepath.Clean(filepath.Join("/var", dir, f3))
	if !filepath.HasPrefix(filePath, "/var/"+dir) {
		log.Fatalf("Invalid file path: %s", filePath)
	}

	body, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}
	log.Print(body)
}
