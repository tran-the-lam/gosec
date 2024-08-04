package main

import (
	"archive/zip"
	"io"
	"os"
	"strconv"
)

func main() {
	r, err := zip.OpenReader("tmp.zip")
	if err != nil {
		panic(err)
	}
	defer r.Close()
	const maxDecompressedSize = 1 << 20 // 1MB

	for i, f := range r.File {
		out, err := os.OpenFile("output"+strconv.Itoa(i), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			panic(err)
		}

		rc, err := f.Open()
		if err != nil {
			panic(err)
		}

		limitedReader := &io.LimitedReader{R: rc, N: maxDecompressedSize}

		_, err = io.Copy(out, limitedReader)

		out.Close()
		rc.Close()

		if err != nil {
			panic(err)
		}
	}
}
