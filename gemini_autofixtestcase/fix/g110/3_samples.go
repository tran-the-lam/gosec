
package main

import (
	"io"
	"os"
)

func main() {
	s, err := os.Open("src")
	if err != nil {
		panic(err)
	}
	defer s.Close()

	d, err := os.Create("dst")
	if err != nil {
		panic(err)
	}
	defer d.Close()

	_, err = io.Copy(d, s)
	if  err != nil {
		panic(err)
	}
}
