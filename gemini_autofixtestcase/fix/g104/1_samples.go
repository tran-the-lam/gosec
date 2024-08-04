package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func a() error {
	return fmt.Errorf("This is an error")
}

func b() {
	fmt.Println("b")
	if err := ioutil.WriteFile("foo.txt", []byte("bar"), os.ModeExclusive); err != nil {
		fmt.Println("error writing file")
	}
}

func c() string {
	return fmt.Sprintf("This isn't anything")
}

func main() {
	_ = a()
	if e := a(); e != nil {
		fmt.Println(e)
	}

	b()
	c()
}
