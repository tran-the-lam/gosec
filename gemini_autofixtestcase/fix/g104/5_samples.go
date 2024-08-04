package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func a() error {
	fmt.Println("a")
	err := ioutil.WriteFile("foo.txt", []byte("bar"), os.ModeExclusive)
	if err != nil {
		fmt.Println("error writing file")
		return err
	}
	return nil
}

func main() {
	err := a()
	if err != nil {
		fmt.Println(err)
	}
}
