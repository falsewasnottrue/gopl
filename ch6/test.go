package main

import (
	"fmt"
)

func main() {
	fmt.Println("IntList test suite")

	l := build(1,2,3)

	fmt.Printf("%v\n", l)
	fmt.Println(l.Head())
	fmt.Printf("%v\n", l.Tail())
	fmt.Println(l.Tail().Head())

	l = l.Append(build(4,5,6))
	fmt.Printf("%v\n", l)

	// foldl
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

	// fmap
	var addOne = func(x int) int {
		return x+1
	}
	fmt.Printf("%v\n", l.FMap(addOne))
	fmt.Printf("%v\n", l)

	// filter
	var isEven = func(n int) bool {
		return n % 2 == 0
	}
	fmt.Printf("%v\n", l.Filter(isEven))

	var isOdd = func(n int) bool {
		return n % 2 == 1
	}
	fmt.Printf("%v\n", l.Filter(isOdd))

	// bind
	var expand = func(n int) *IntList {
		return build(n, n)
	}
	fmt.Printf("%v\n", l.Bind(expand))

	// ---

	fmt.Println("IntOption test suite")

	var checkEven = func(val int) IntOption {
		if isEven(val) {
			return Some{value: val}
		} else {
			return None {}
		}
	}

	// ---
	var someone = Some{value: 1}
	
	fmt.Printf("%t %v\n", someone, someone)
	fmt.Printf("Some(1).get: %v\n", someone.Get())
	fmt.Println("Some(1).getOrElse(42): %v\n", someone.GetOrElse(42))

	fmt.Printf("Some(1).filter(isEven): %v\n", someone.Filter(isEven))
	fmt.Printf("Some(1).filter(isOdd): %v\n", someone.Filter(isOdd))

	fmt.Printf("Some(1).map(+1): %v\n", someone.FMap(addOne))
	fmt.Printf("Some(1).bind(checkEven) %v\n", someone.Bind(checkEven))
	fmt.Printf("Some(2).bind(checkEven) %v\n", Some{value: 2}.Bind(checkEven))

	// ---

	var noone = None{}

	fmt.Printf("%t %v\n", noone, noone)
	// fmt.Println(noone.Get()) -> panic
	fmt.Printf("None.getOrElse(42): %v\n", noone.GetOrElse(42))
	fmt.Printf("None.map(+1): %v\n", noone.FMap(addOne))
	fmt.Printf("None.bind(checkEven): %v\n", noone.Bind(checkEven))
}