
package main

import (
    "os"
    "path/filepath"
)

func openFile(filePath string) {
	_, err := os.OpenFile(filepath.Clean(filePath), os.O_RDONLY, 0600)
	if err != nil {
		panic(err)
	}
}

func main() {
    repoFile := "path_of_file"
	openFile(repoFile)
}
