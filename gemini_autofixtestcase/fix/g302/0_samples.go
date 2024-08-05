package main

import (
	"fmt"
	"os"
)

func main() {
	err := os.Chmod("/tmp/somefile", 0600)
	if err != nil {
		fmt.Println("Error when changing file permissions!")
		return
	}
}
