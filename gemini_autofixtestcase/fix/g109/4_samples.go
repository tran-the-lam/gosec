
package main

import (
	"fmt"
	"strconv"
)

func main() {
	value := 10
	if value == 10 {
		value, _ := strconv.Atoi("2147483648")
		fmt.Println(value)
	}
	v := int64(value)
	fmt.Println(v)
}
