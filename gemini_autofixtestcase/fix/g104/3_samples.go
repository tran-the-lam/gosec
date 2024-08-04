
// +build go1.10

package main

import "strings"

func main() {
	var buf strings.Builder
	_, err := buf.WriteString("test string")
	if err != nil {
		panic(err)
	}
}
package main

func dummy(){}
