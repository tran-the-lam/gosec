package main

import (
	"fmt"
	"os"
)

func main() {
	pw := os.Getenv("PASSWORD")
	fmt.Println(pw)
}
