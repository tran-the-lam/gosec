// Initializing a local variable using a environmental
// variable is consider as a dangerous user input
package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {
	run := "sleep" + os.Getenv("SOMETHING")
	arg := "5"
	cmd := exec.Command(run, arg)
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Waiting for command to finish...")
	err = cmd.Wait()
	log.Printf("Command finished with error: %v", err)
}
