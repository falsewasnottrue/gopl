package main

import (
	"fmt"
	"strconv"
)

type IntList struct {
	Value int
	Next *IntList
}

func build(vals ...int) *IntList {
	var res *IntList
	for i := len(vals)-1; i>=0; i-- {
		l := new (IntList)
		l.Value = vals[i]
		l.Next = res

		res = l
	}

	return res
}

func (list *IntList) String() string {
	if (list == nil) {
		return "."
	}

	if (list.Next == nil) {
		return strconv.Itoa(list.Value) + "."
	}

	return strconv.Itoa(list.Value) + "," + list.Next.String()
}

func (list *IntList) Head() int {
	if (list == nil) {
		panic("no head on empty list")
	}
	return list.Value
}

func (list *IntList) Tail() *IntList {
	if (list == nil) {
		return nil
	}
	return list.Next
}

func (list *IntList) Append(other *IntList) *IntList {
	if (list == nil) {
		return other
	}
	l := list
	for ; l.Next != nil; l = l.Next {}
	l.Next = other

	return list
}

func (list *IntList) Foldl(init int, combine func(x,y int) int) int {
	var acc = init
	for l := list; l != nil; l = l.Next {
		acc = combine(acc, l.Value)
	}
	return acc
}

// TODO fmap
// TODO filter
// TODO bind

func main() {
	l := build(1,2,3)

	fmt.Printf("%v\n", l)
	fmt.Println(l.Head())
	fmt.Printf("%v\n", l.Tail())
	fmt.Println(l.Tail().Head())

	l = l.Append(build(4,5,6))
	fmt.Printf("%v\n", l)

	var add = func(x, y int) int {
		return x+y
	}
	sum := l.Foldl(0, add)
	fmt.Println(sum)

	var mult = func(x, y int) int {
		return x*y
	}
	prod := l.Foldl(1, mult)
	fmt.Println(prod)
}

