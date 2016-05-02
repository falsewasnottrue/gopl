package main

import (
	"fmt"
)

type IntOption interface {
	Get() int
	GetOrElse(i int) int

	// TODO FMap
	// TODO Filter
	// TODO Bind
}

type Some struct {
	value int
}

func (ctx Some) Get() int {
	return ctx.value
}

func (ctx Some) GetOrElse(i int) int {
	return ctx.value
}

// TODO implementation None
type None struct {}

func (ctx None) Get() int {
	panic("Cannot access value of None")
}

func (ctx None) GetOrElse(i int) int {
	return i
}

func main() {
	var someone = Some{value: 1}
	
	fmt.Printf("%t %v\n", someone, someone)
	fmt.Println(someone.Get())
	fmt.Println(someone.GetOrElse(42))

}