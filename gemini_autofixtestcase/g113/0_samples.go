
package main

import (
	"math/big"
	"fmt"
)

func main() {
	r := big.Rat{}
	r.SetString("13e-9223372036854775808")

	fmt.Println(r)
}
