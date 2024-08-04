package main

import (
	"fmt"
	"os"
)

func main() {
	cred := os.Getenv("cred")
	fmt.Println(cred)
}
