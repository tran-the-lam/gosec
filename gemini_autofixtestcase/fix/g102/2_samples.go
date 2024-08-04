
package main

import (
	"log"
	"net"
)

func parseListenAddr(listenAddr string) (network string, addr string) {
	return "", ""
}

func main() {
	addr := ":2000"
	l, err := net.Listen(parseListenAddr(addr))
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
}
