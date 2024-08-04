package main

import "fmt"

func main() {
	bearer := os.GetEnv("bearer")
	fmt.Println(bearer)
}
