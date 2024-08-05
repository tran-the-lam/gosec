
package main

import (
	"fmt"
	"os"
)

func main() {
	err := os.Chmod("/tmp/mydir", 0400)
	if err != nil {
		fmt.Println("Error")
		return
	}
}
