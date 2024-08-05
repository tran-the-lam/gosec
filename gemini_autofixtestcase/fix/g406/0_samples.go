package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	data := []byte("test")
	hash := sha256.Sum256(data)
	fmt.Println(hash)
}
