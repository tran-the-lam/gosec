package main

import (
	"fmt"
	"os"
)

func main() {
	var password string
	if os.Getenv("PASSWORD") == password {
		fmt.Println("password equality")
	}
}
