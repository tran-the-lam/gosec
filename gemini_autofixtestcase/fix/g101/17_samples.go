package main

import (
	"fmt"
	"os"
)

var (
	cred string
)

func main() {
	cred = os.Getenv("cred")
	fmt.Println(cred)
}
