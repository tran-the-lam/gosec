package main

import (
	"fmt"
	"strconv"
)

func main() {
	bigValue, err := strconv.Atoi("2147483648")
	if err != nil {
		panic(err)
	}
	value := int64(bigValue)
	fmt.Println(value)
}
