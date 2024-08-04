
package main

import "fmt"

func test() error {
	return nil
}

func main() {
	e := test()
	fmt.Println(e)
}
