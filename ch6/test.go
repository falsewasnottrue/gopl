package main

import (
	"fmt"
)

var isEven = func(n int) bool {
	return n % 2 == 0
}

var isOdd = func(n int) bool {
	return n % 2 == 1
}

var checkEven = func(val int) IntOption {
	if isEven(val) {
		return Some{value: val}
	} else {
		return None {}
	}
}

var expand = func(n int) *IntList {
	return build(n, n)
}

var add = func(x, y int) int {
	return x+y
}

var addOne = func(x int) int {
	return x+1
}

var mult = func(x, y int) int {
	return x*y
}

func main() {
	fmt.Println("IntList test suite")

	l := build(1,2,3)

	fmt.Printf("IntList l1: %v\n", l)
	fmt.Printf("l1.Head: %v\n", l.Head())
	fmt.Printf("l1.HeadOption: %v\n", l.HeadOption())
	fmt.Printf("l1.Tail: %v\n", l.Tail())
	fmt.Printf("l1.Tail.Head: %v\n", l.Tail().Head())

	l = l.Append(build(4,5,6))
	fmt.Printf("l1.Append(4,5,6): %v\n", l)

	sum := l.Foldl(0, add)
	fmt.Printf("l1.Foldl(0, add): %v\n", sum)
	prod := l.Foldl(1, mult)

	fmt.Printf("l1.Foldl(1, mult): %v\n", prod)
	fmt.Printf("l1.FMap(addOne): %v\n", l.FMap(addOne))
	fmt.Printf("l1: %v\n", l)
	fmt.Printf("l1.Filter(isEven): %v\n", l.Filter(isEven))
	fmt.Printf("l1.Filter(isOdd): %v\n", l.Filter(isOdd))
	fmt.Printf("l1.Bind(expand): %v\n", l.Bind(expand))

	// ---
	fmt.Printf("\nIntOption test suite\n")

	var maybeInt IntOption

	maybeInt = Some{value: 1}
	
	fmt.Printf("Some(1).get: %v\n", maybeInt.Get())
	fmt.Printf("Some(1).getOrElse(42): %v\n", maybeInt.GetOrElse(42))

	fmt.Printf("Some(1).filter(isEven): %v\n", maybeInt.Filter(isEven))
	fmt.Printf("Some(1).filter(isOdd): %v\n", maybeInt.Filter(isOdd))

	fmt.Printf("Some(1).map(+1): %v\n", maybeInt.FMap(addOne))
	fmt.Printf("Some(1).bind(checkEven) %v\n", maybeInt.Bind(checkEven))
	fmt.Printf("Some(2).bind(checkEven) %v\n", Some{value: 2}.Bind(checkEven))

	maybeInt = None{}
	// fmt.Println(noone.Get()) -> panic
	fmt.Printf("None.getOrElse(42): %v\n", maybeInt.GetOrElse(42))
	fmt.Printf("None.map(+1): %v\n", maybeInt.FMap(addOne))
	fmt.Printf("None.bind(checkEven): %v\n", maybeInt.Bind(checkEven))

	// ---
	fmt.Printf("\nIntSet test suite\n")

	var intSet IntSet = EmptySet()
	fmt.Printf("emptySet: %v\n", intSet)
	intSet.Add(1)
	fmt.Printf("{1}: %v\n", intSet)
	intSet.Add(1)
	fmt.Printf("{1}: %v\n", intSet)
	intSet.Add(2)
	fmt.Printf("{1,2}: %v\n", intSet)

	intSet.Remove(1)
	fmt.Printf("{2}: %v\n", intSet)
}