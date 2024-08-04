
// syscall.Exec function called with hardcoded arguments
// shouldn't be consider as a command injection
package main

import (
	"fmt"
	"syscall"
)

func main() {
	err := syscall.Exec("/bin/cat", []string{"/etc/passwd"}, nil)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
