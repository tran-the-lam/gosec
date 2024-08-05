package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func main() {
	f2 := os.Getenv("tainted_file2")
	f2, err := sanitizeInput(os.Getenv("tainted_file2"))
	if err != nil {
		log.Fatalf("Invalid tainted_file2: %v", err)
	}

	body, err := ioutil.ReadFile(filepath.Clean(filepath.Join("/tmp/", f2)))
	if err != nil {
		log.Printf("Error: %v\n", err)
	}
	log.Print(body)
}
