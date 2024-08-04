
package main

import (
	"fmt"
	"unsafe"
)

func main() {
	chars := [...]byte{1, 2}
	charsPtr := &chars[0]
	str := unsafe.String(charsPtr, len(chars))
	fmt.Printf("%s\n", str)
	ptr := unsafe.StringData(str)
	fmt.Printf("ptr: %p\n", ptr)
}
