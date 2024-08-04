package main

import (
	"fmt"
	"strconv"
)

func main() {
	bigValue, err := strconv.ParseInt("32768", 10, 16)
	if err != nil {
		panic(err)
	}
	if bigValue < 0 {
		fmt.Println(bigValue)
	}
}
