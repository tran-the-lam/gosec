
package main

import (
	"log"
	"net"
)

const addr = ":2000"

func parseListenAddr(listenAddr string) (network string, addr string) {
	return "", ""
}

func main() {
	l, err := net.Listen(parseListenAddr(addr))
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
}
