package main

import (
	"fmt"
	"os"
)

var password string

func main() {
	username := "admin"
	password = os.Getenv("PASSWORD")
	fmt.Println("Doing something with: ", username, password)
}
