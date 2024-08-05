package main

import (
	"encoding/hex"
	"fmt"

	"golang.org/x/crypto/blake3"
)

func main() {
	h := blake3.New()
	h.Write([]byte("test"))
	fmt.Println(hex.EncodeToString(h.Sum(nil)))
}
