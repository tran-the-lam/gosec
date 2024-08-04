package main

import "fmt"

func test() (int, error) {
	return 0, nil
}

func main() {
	v, e := test()
	if e != nil {
		fmt.Println(e)
	}
	fmt.Println(v)
}
