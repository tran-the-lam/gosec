package main

import (
	"fmt"
	"os"
)

const (
	username = "user"
)

func main() {
	password := os.Getenv("PASSWORD")
	fmt.Println("Doing something with: ", username, password)
}
