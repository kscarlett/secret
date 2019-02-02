package main

import (
	"fmt"

	"github.com/kscarlett/secret/pkg/secret"
)

func main() {
	s := secret.NewInMemory("demo-password")
	err := s.Set("example-key", "23819d20-9b07-4b20-8a6e-a3c533fa4994")
	if err != nil {
		panic(err)
	}
	exampleKey, err := s.Get("example-key")
	if err != nil {
		panic(err)
	}
	fmt.Println("The key is: ", exampleKey)
}
