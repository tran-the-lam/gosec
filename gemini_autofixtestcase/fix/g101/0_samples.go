package main

import (
	"fmt"
	"os"
)

func main() {
	username := "admin"
	password := os.Getenv("PASSWORD")
	fmt.Println("Doing something with: ", username, password)
}
