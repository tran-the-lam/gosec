
package main

import (
    "os"
    "path/filepath"
)

func openFile(dir string, filePath string) {
	fp := filepath.Join(dir, filePath)
	fp = filepath.Clean(fp)
	_, err := os.OpenFile(fp, os.O_RDONLY, 0600)
	if err != nil {
		panic(err)
	}
}

func main() {
    repoFile := "path_of_file"
	dir := "path_of_dir"
	openFile(dir, repoFile)
}
