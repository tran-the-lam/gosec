package main

import (
	"bytes"
	"compress/zlib"
	"io"
	"os"
)

func main() {
	buff := []byte{120, 156, 202, 72, 205, 201, 201, 215, 81, 40, 207,
		47, 202, 73, 225, 2, 4, 0, 0, 255, 255, 33, 231, 4, 147}
	b := bytes.NewReader(buff)

	r, err := zlib.NewReader(b)
	defer r.Close()
	if err != nil {
		panic(err)
	}
	buf := make([]byte, 8)
	const maxDecompressedSize = 1 << 20 // 1MB
	limitedReader := &io.LimitedReader{R: r, N: maxDecompressedSize}

	_, err = io.CopyBuffer(os.Stdout, limitedReader, buf)
	if err != nil {
		panic(err)
	}
}
