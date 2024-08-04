
package main

import "fmt"

// #nosec G101
const (
	ConfigLearnerTokenAuth string = "learner_auth_token_config"
)

func main() {
	fmt.Printf("%s\n", ConfigLearnerTokenAuth)
}
	
