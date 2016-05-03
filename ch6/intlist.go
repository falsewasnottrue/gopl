package main

import (
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

func (list *IntList) HeadOption() IntOption {
	if (list == nil) {
		return None{}
	}
	return Some{value: list.Value}
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

func (list *IntList) FMap(f func(x int) int) *IntList {
	if (list == nil) {
		return nil
	}
	var head = new (IntList)
	head.Value = f(list.Value)

	return head.Append(list.Next.FMap(f))
}

func (list* IntList) Filter(p func(x int) bool) *IntList {
	if (list == nil) {
		return nil
	}
	if (p(list.Value) == false) {
		return list.Next.Filter(p)
	}

	// p(list.Value) == true
	var head = new (IntList)
	head.Value = list.Value
	return head.Append(list.Next.Filter(p))
}

func (list *IntList) Bind(f func(x int) *IntList) *IntList {
	if (list == nil) {
		return list
	}

	return f(list.Head()).Append(list.Next.Bind(f))
}