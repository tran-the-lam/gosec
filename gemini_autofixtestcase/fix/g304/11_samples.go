package main

import (
	"io"
	"os"
	"path/filepath"
)

func createFile(file string) *os.File {
	f, err := os.Create(filepath.Clean(file))
	if err != nil {
		panic(err)
	}
	return f
}

func main() {
	s, err := os.Open("src")
	if err != nil {
		panic(err)
	}
	defer s.Close()

	d := createFile("dst")
	defer d.Close()

	_, err = io.Copy(d, s)
	if err != nil {
		panic(err)
	}
}
