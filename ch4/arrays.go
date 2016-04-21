package main

import (
	"fmt"
	"crypto/sha256"
)

func main() {
	text := "Hello world!"
	ch1 := sha256.Sum256([]byte(text))

	fmt.Printf("%T\n%v\n", ch1, ch1)
}