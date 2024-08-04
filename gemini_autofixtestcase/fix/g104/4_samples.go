
package main

import (
	"bytes"
)

type a struct {
	buf *bytes.Buffer
}

func main() {
	a := &a{
		buf: new(bytes.Buffer),
	}
	a.buf.Write([]byte{0})
}
