
package main

import (
	"log"
	"net"
)

func main() {
		l, err := net.Listen("tcp", ":2000")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
}
