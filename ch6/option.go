package main

import (
	// "fmt"
	"strconv"
)

type IntOption interface {
	Get() int
	GetOrElse(i int) int

	Filter(p func(val int) bool) IntOption
	FMap(f func(val int) int) IntOption
	Bind(f func(val int) IntOption) IntOption
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

func (ctx Some) FMap(f func(val int) int) IntOption {
	return Some{value : f(ctx.value)}
}

func (ctx Some) Bind(f func(val int) IntOption) IntOption {
	return f(ctx.value)
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
	return ctx // None
}

func (ctx None) FMap(f func(val int) int) IntOption {
	return ctx // None
}

func (ctx None) Bind(f func(val int) IntOption) IntOption {
	return ctx // None
}