
package main

import (
	"fmt"
	"path/filepath"
)

type foo struct {
}

func (f *foo) doSomething(silly string) error {
	whoCares, err := filepath.Rel(THEWD, silly)
	if err != nil {
		return err
	}
	fmt.Printf("%s", whoCares)
	return nil
}

func main() {
	f := &foo{}

	if err := f.doSomething("irrelevant"); err != nil {
		panic(err)
	}
}

package main

var THEWD string
