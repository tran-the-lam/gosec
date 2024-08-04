package main

import (
	"fmt"
	"net/http"
	"net/url"
)

// Escape special characters in the URL components
func escapeURL(rawURL string) (string, error) {
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}

	escapedURL := &url.URL{
		Scheme:   parsedURL.Scheme,
		Host:     parsedURL.Host,
		Path:     url.PathEscape(parsedURL.Path),
		RawQuery: url.QueryEscape(parsedURL.RawQuery),
	}

	return escapedURL.String(), nil
}

// Validate the URL to ensure it adheres to expected formats
func validateURL(rawURL string) bool {
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return false
	}

	if parsedURL.Scheme != "http" && parsedURL.Scheme != "https" {
		return false
	}

	if parsedURL.Host == "" {
		return false
	}

	return true
}

func get(urlStr string) {
	if !validateURL(urlStr) {
		fmt.Println("Invalid URL")
		return
	}

	escapedURL, err := escapeURL(urlStr)
	if err != nil {
		fmt.Println("Error escaping URL:", err)
		return
	}

	resp, err := http.Get(escapedURL)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	fmt.Println(resp.Status)
}

func main() {
	const urlStr = "http://127.0.0.1"
	get(urlStr)
}
