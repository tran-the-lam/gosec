
// starting a process with a variable as an argument
// even if not constant is not considered as dangerous
// because it has hardcoded value
package main

import (
	"log"
	"os/exec"
)

func main() {
	run := "sleep"
	cmd := exec.Command(run, "5")
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Waiting for command to finish...")
	err = cmd.Wait()
	log.Printf("Command finished with error: %v", err)
}
