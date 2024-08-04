
package main

import (
	"log"
	"os/exec"
	"context"
)

func main() {
	err := exec.CommandContext(context.Background(), "git", "rev-parse", "--show-toplevel").Run()
 	if err != nil {
		log.Fatal(err)
	}
  	log.Printf("Command finished with error: %v", err)
}
