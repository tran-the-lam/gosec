
package main

import (
	"fmt"
	"syscall"
)

func RunCmd(command string) {
	_, err := syscall.ForkExec(command, []string{}, nil)
	if err != nil {
	    fmt.Printf("Error: %v\n", err)
	}
}

func main() {
	RunCmd("sleep")
}
