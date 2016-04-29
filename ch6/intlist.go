package main

import (
	"fmt"
	"strconv"
)

type IntList struct {
	Value int
	Next *IntList
}

// TODO build
func build(vals ...int) *IntList {
	return nil
}

// TODO toString
func (list *IntList) String() string {
	if (list == nil) {
		return "."
	}

	if (list.Next == nil) {
		return strconv.Itoa(list.Value) + "."
	}

	return strconv.Itoa(list.Value) + "," + list.Next.String()
}

// TODO head
// TODO tail
// TODO last
// TODO append
// TODO fold
// TODO fmap
// TODO filter
// TODO bind

func main() {
	i := 42
	fmt.Println(string(i))

	l := build(1,2,3)

	fmt.Printf("%v\n", l)
}