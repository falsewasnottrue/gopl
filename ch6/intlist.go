package main

import (
	"fmt"
)

type IntList struct {
	Value int
	Next *IntList
}

// TODO head
// TODO tail
// TODO last
// TODO append
// TODO fmap
// TODO filter
// TODO bind

func main() {
	l := new (IntList)
	l.Value = 4711

	fmt.Printf("%v\n", l)
}