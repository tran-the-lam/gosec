package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func main() {
	h := sha256.New()
	h.Write([]byte("test"))
	fmt.Println(hex.EncodeToString(h.Sum(nil)))
}
