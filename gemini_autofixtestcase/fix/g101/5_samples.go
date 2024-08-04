package main

import "os"

var password string

func init() {
	password = os.Getenv("PASSWORD")
}
