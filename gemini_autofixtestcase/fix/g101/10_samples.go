package main

import (
	"fmt"
	"os"
)

func main() {
	var password string
	if password != os.Getenv("PASSWORD") {
		fmt.Println("password equality")
	}
}
