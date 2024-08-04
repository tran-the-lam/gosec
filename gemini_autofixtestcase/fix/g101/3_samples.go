package main

import (
	"fmt"
	"os"
)

const password = os.Getenv("PASSWORD")

func main() {
	username := "admin"
	fmt.Println("Doing something with: ", username, password)
}
