package main

import (
	"fmt"
	"unsafe"
)

func main() {
	chars := [...]byte{1, 2}
	charsPtr := &chars[0]
	slice := unsafe.Slice(charsPtr, len(chars))
	fmt.Printf("%v\n", slice)
	ptr := unsafe.SliceData(slice)
	fmt.Printf("ptr: %p\n", ptr)
}
