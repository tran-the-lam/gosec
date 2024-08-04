package main

import (
	"fmt"
	"os"
)

var (
	pw string
)

func main() {
	pw = os.Getenv("PASSWORD")
	fmt.Println(pw)
}
