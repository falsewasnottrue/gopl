package main

import "fmt"

func main() {
	var slice = make([]int, 3, 3)
	fmt.Println(slice)

	var addOne = func(x int) int {
		return x + 1
	}
	fmt.Println(fmap(slice, addOne))

	f := dups()
	fmt.Println(bind(slice, f))
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


func bind(ls []int, f func(i int) []int) []int {
	return ls
}