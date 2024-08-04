
package main

import (
	"log"
	"net"
	"net/http"
)

func main() {
	l, err := net.Listen("tcp", ":8443")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
	err = http.ServeTLS(l, nil, "cert.pem", "key.pem")
	log.Fatal(err)
}
