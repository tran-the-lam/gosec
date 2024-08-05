package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please enter file to read: ")
	file, _ := reader.ReadString('\n')
	file = file[:len(file)-1]
	f, err := os.Open(filepath.Clean(filepath.Join("/tmp/service/", file)))
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	contents := make([]byte, 15)
	if _, err = f.Read(contents); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Println(string(contents))
}
