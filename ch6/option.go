package main

import (
	"fmt"
	"strconv"
)

type IntOption interface {
	Get() int
	GetOrElse(i int) int

	Filter(p func(val int) bool) IntOption

	// TODO FMap
	// TODO Bind
}

// --- Some
type Some struct {
	value int
}

func (ctx Some) String() string {
	return "Some(" + strconv.Itoa(ctx.value) + ")"
}

func (ctx Some) Get() int {
	return ctx.value
}

func (ctx Some) GetOrElse(i int) int {
	return ctx.value
}

func (ctx Some) Filter(p func(val int) bool) IntOption {
	if p(ctx.value) {
		return ctx
	} else {
		return None{}
	}
}

// --- None
type None struct {}

func (ctx None) String() string {
	return "None"
}

func (ctx None) Get() int {
	panic("Cannot access value of None")
}

func (ctx None) GetOrElse(i int) int {
	return i
}

func (ctx None) Filter(f func(i int) bool) IntOption {
	return ctx
}

// ---
func main() {
	var isEven = func(val int) bool {
		return val % 2 == 0
	}
	var isOdd = func(val int) bool {
		return val % 2 == 1
	}

	// ---
	var someone = Some{value: 1}
	
	fmt.Printf("%t %v\n", someone, someone)
	fmt.Printf("Some(1).get: %v\n", someone.Get())
	fmt.Println("Some(1).getOrElse(42): %v\n", someone.GetOrElse(42))

	fmt.Printf("Some(1).filter(isEven): %v\n", someone.Filter(isEven))
	fmt.Printf("Some(1).filter(isOdd): %v\n", someone.Filter(isOdd))

	// ---

	var noone = None{}

	fmt.Printf("%t %v\n", noone, noone)
	// fmt.Println(noone.Get()) -> panic
	fmt.Printf("None.getOrElse(42): %v\n", noone.GetOrElse(42))

}