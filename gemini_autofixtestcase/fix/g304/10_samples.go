
package main

import (
    "os"
    "path/filepath"
)

func main() {
    repoFile := "path_of_file"
	relFile, err := filepath.Rel("./", repoFile)
	if err != nil {
		panic(err)
	}
    _, err = os.OpenFile(relFile, os.O_RDONLY, 0600)
    if err != nil {
        panic(err)
    }
}
