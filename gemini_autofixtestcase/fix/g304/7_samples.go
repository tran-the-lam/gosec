
package main

import (
    "os"
    "path/filepath"
)

func main() {
    repoFile := "path_of_file"
    cleanRepoFile := filepath.Clean(repoFile)
    _, err := os.OpenFile(cleanRepoFile, os.O_RDONLY, 0600)
    if err != nil {
        panic(err)
    }
}
