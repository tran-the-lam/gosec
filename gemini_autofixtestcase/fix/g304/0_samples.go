package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func main() {
	f := os.Getenv("tainted_file")
	// f, err := sanitizeInput(os.Getenv("tainted_file"))
	// if err != nil {
	// 	log.Fatalf("Invalid tainted_file: %v", err)
	// }
	body, err := ioutil.ReadFile(filepath.Clean(f))
	if err != nil {
		log.Printf("Error: %v\n", err)
	}
	log.Print(body)

}
