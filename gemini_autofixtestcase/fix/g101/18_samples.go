package main

import (
	"fmt"
	"os"
)

func main() {
	apiKey := os.Getenv("ApiKey")
	fmt.Println(apiKey)
}
