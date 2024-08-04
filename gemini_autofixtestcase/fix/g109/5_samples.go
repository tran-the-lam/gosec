
package main

import (
	"fmt"
	"strconv"
)
func main() {
	a, err := strconv.Atoi("a")
	b := int64(a) //#nosec G109
	fmt.Println(b, err)
}
