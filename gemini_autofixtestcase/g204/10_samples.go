
// Initializing a local variable using a environmental
// variable is consider as a dangerous user input
package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {
	var run = "sleep" + os.Getenv("SOMETHING")
	cmd := exec.Command(run, "5")
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Waiting for command to finish...")
	err = cmd.Wait()
	log.Printf("Command finished with error: %v", err)
}
