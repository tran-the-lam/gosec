
package main

import (
	"log"
	"net/http"
)

func main() {
	err := http.ListenAndServeTLS(":8443", "cert.pem", "key.pem", nil)
	log.Fatal(err)
}
