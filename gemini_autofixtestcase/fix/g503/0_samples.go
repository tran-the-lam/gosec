package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"log"
)

func main() {
	key := make([]byte, 32) // AES-256 requires a 32-byte key
	if _, err := rand.Read(key); err != nil {
		log.Fatal(err)
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		log.Fatal(err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Fatal(err)
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		log.Fatal(err)
	}

	plaintext := []byte("I CAN HAZ SEKRIT MSG PLZ")
	ciphertext := gcm.Seal(nonce, nonce, plaintext, nil)

	fmt.Printf("Secret message is: %s\n", hex.EncodeToString(ciphertext))
}
