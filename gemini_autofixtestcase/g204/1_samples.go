
// Calling any function which starts a new process with using
// command line arguments as it's arguments is considered dangerous
package main

import (
	"context"
	"log"
	"os"
	"os/exec"
)

func main() {
	err := exec.CommandContext(context.Background(), os.Args[0], "5").Run()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Command finished with error: %v", err)
}
