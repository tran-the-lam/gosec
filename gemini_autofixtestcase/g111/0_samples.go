package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
)

func main() {
	http.Handle("/bad/", http.StripPrefix("/bad/", http.FileServer(http.Dir("/"))))
	http.HandleFunc("/", HelloServer)
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	// Step 1: Clean the input path
	cleanPath := filepath.Clean(r.URL.Path[1:])

	// Step 2: Enforce allowed characters (e.g., alphanumeric, hyphens, underscores, slashes)
	validPathPattern := `^[a-zA-Z0-9_\-/]+$`
	matched, err := regexp.MatchString(validPathPattern, cleanPath)
	if err != nil || !matched {
		http.Error(w, "Invalid path", http.StatusBadRequest)
		return
	}

	// Step 3: Whitelist allowed paths (example: only allow paths under /safe/directory)
	allowedPaths := []string{"/safe/directory"}
	isAllowed := false
	for _, allowedPath := range allowedPaths {
		if filepath.HasPrefix(cleanPath, allowedPath) {
			isAllowed = true
			break
		}
	}
	if !isAllowed {
		http.Error(w, "Path is not allowed", http.StatusForbidden)
		return
	}

	// Step 4: Normalize the path and compare it to the safe directory
	absPath, err := filepath.Abs(cleanPath)
	if err != nil {
		http.Error(w, "Error processing path", http.StatusInternalServerError)
		return
	}

	safeDir := "/safe/directory"
	if !filepath.HasPrefix(absPath, safeDir) {
		http.Error(w, "Path is not allowed", http.StatusForbidden)
		return
	}

	fmt.Fprintf(w, "Hello, %s!", cleanPath)
}
