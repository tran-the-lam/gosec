package main

import (
	"fmt"
	"os"
)

func main() {
	var p string
	if p != os.Getenv("PASSWORD") {
		fmt.Println("password equality")
	}
}
