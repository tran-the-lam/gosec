package main

import "fmt"

var (
	bearer string
)

func main() {
	bearer = os.GetEnv("bearer")
	fmt.Println(bearer)
}
