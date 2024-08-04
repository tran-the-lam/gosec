package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	http.HandleFunc("/bad/", safeFileServer)
	http.HandleFunc("/", HelloServer)
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}

func safeFileServer(w http.ResponseWriter, r *http.Request) {
	// Sanitize and validate the URL path
	cleanPath := filepath.Clean(r.URL.Path)
	if !strings.HasPrefix(cleanPath, "/bad/") {
		http.Error(w, "Invalid path", http.StatusBadRequest)
		return
	}

	// Remove the "/bad/" prefix
	relativePath := strings.TrimPrefix(cleanPath, "/bad/")
	// Whitelist allowed characters (alphanumeric and some special characters)
	if !isValidPath(relativePath) {
		http.Error(w, "Invalid path", http.StatusBadRequest)
		return
	}

	// Serve the file
	fileServer := http.FileServer(http.Dir("/"))
	http.StripPrefix("/bad/", fileServer).ServeHTTP(w, r)
}

func isValidPath(path string) bool {
	for _, char := range path {
		if !(char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z' || char >= '0' && char <= '9' || char == '-' || char == '_' || char == '.') {
			return false
		}
	}
	return true
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}
