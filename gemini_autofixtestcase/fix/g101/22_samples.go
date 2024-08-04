package main

import (
	"fmt"
)

// #nosec G101
const (
	ConfigLearnerTokenAuth string = "learner_auth_token_config" // #nosec G101
)

func main() {
	fmt.Printf("%s\n", ConfigLearnerTokenAuth)
}
