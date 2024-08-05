package main

import (
	"io/ioutil"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	content := []byte("hello\ngo\n")
	err := ioutil.WriteFile("/tmp/dat1", content, 0600)
	check(err)
}
