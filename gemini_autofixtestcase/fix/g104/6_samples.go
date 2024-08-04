
package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

func createBuffer() *bytes.Buffer {
	return new(bytes.Buffer)
}

func main() {
	new(bytes.Buffer).WriteString("*bytes.Buffer")
	fmt.Fprintln(os.Stderr, "fmt")
	new(strings.Builder).WriteString("*strings.Builder")
	_, pw := io.Pipe()
	pw.CloseWithError(io.EOF)

	createBuffer().WriteString("*bytes.Buffer")
	b := createBuffer()
	b.WriteString("*bytes.Buffer")
}
