package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
	})
	err := (&http.Server{
		Addr:              ":1234",
		ReadHeaderTimeout: 3,
	}).ListenAndServe()
	if err != nil {
		panic(err)
	}
}
