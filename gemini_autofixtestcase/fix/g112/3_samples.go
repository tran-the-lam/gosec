
package main

import (
	"fmt"
	"net/http"
	"sync"
)

type Server struct {
	hs  *http.Server
	mux *http.ServeMux
	mu  sync.Mutex
}

func New(listenAddr string) *Server {
	mux := http.NewServeMux()

	return &Server{
	hs: &http.Server{ // #nosec G112 - Not publicly exposed
		Addr:    listenAddr,
		Handler: mux,
	},
	mux: mux,
	mu:  sync.Mutex{},
	}
}

func main() {
	fmt.Print("test")
}
