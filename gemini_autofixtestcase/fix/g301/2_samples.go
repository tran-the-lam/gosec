
package main

import (
	"fmt"
	"os"
)

func main() {
	err := os.Mkdir("/tmp/mydir", 0600)
	if err != nil {
		fmt.Println("Error when creating a directory!")
		return
	}
}
