package main

import (
	"fmt"
	"os"
)

var (
	apiKey string
)

func main() {
	apiKey = os.Getenv("cred")
	fmt.Println(apiKey)
}
