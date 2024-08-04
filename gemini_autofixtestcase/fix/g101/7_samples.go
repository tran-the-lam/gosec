package main

import "os"

func main() {
	ATNStateTokenStart := os.Getenv("ATNStateTokenStart")
	println(ATNStateTokenStart)
}
