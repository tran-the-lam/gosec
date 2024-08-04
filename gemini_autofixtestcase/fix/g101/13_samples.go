package main

import (
	"fmt"
	"os"
)

func main() {
	var p string
	if os.Getenv("PASSWORD") != p {
		fmt.Println("password equality")
	}
}
