
// exec.Command from supplemental package sys/execabs
// using variable arguments
package main

import (
	"context"
	"log"
	"os"
	exec "golang.org/x/sys/execabs"
)

func main() {
	err := exec.CommandContext(context.Background(), os.Args[0], "5").Run()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Command finished with error: %v", err)
}
