
package main

import (
	"fmt"
	"syscall"
)

func RunCmd(command string) {
	_, _, err := syscall.StartProcess(command, []string{}, nil)
	if err != nil {
	    fmt.Printf("Error: %v\n", err)
	}
}

func main() {
	RunCmd("sleep")
}
