
// An exported variable declared a packaged scope is not secure
// because it can changed at any time
package main

import (
	"fmt"
	"net/http"
)

var Url string

func main() {
	resp, err := http.Get(Url)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp.Status)
}
