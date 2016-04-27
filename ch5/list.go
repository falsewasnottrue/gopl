package main

import "fmt"

func main() {
	var slice = []int{1,2,3,4,5}

	fmt.Printf("orig:   %v\n", slice)

	var addOne = func(x int) int {
		return x + 1
	}
	fmt.Printf("fmap:   %v\n", fmap(slice, addOne))

	var odd = func(x int) bool {
		return x % 2 == 1
	}
	fmt.Printf("filter: %v\n", filter(slice, odd))

	var even = func(x int) bool {
		return x % 2 == 0
	}
	fmt.Printf("filter: %v\n", filter(slice, even))

	// f := dups()
	// fmt.Println(bind(slice, f))
}

func dups() func(i int) []int {
	l := 0
	return func(value int) []int {
		l++
		ls := make([]int, l, l)
		for i := 0; i<l; i++ {
			ls[i] = value
		}
		return ls
	}
}

func fmap(ls []int, f func(i int) int) []int {
	var res = make([]int, len(ls), cap(ls))
	for i, val := range ls {
		res[i] = f(val)
	}
	return res
}

// TODO implement
func filter(ls []int, p func(int) bool) []int {
	res := make([]int, len(ls))
	pos := 0
	for _, val := range ls {
		if (p(val)) {
			res[pos] = val
			pos ++
		}
	}
	return res[:pos]
}

func bind(ls []int, f func(i int) []int) []int {
	return ls
}