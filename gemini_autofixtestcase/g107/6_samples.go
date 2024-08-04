
// A variable at function scope which is initialized to
// a constant string is secure (e.g. cannot be changed concurrently)
package main

import (
	"fmt"
	"net/http"
)
func main() {
	url1 := "test"
	var url2 string = "http://127.0.0.1"
	url2 = url1
	resp, err := http.Get(url2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp.Status)
}
