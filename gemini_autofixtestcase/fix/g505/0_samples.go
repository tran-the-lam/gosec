package main

import (
	"crypto/sha256"
	"fmt"
	"os"
)

func main() {
	for _, arg := range os.Args {
		fmt.Printf("%x - %s\n", sha256.Sum256([]byte(arg)), arg)
	}
}
