package main

import (
	"strconv"
)

type IntSet map[int]bool

func (ctx IntSet) String() string {
	var res = "{"
	var i = 0
	for k, v := range ctx {
		if (v) {
			res += strconv.Itoa(k)
			i++
		}
		if (i > 0 && i < ctx.Size()) {
			res += ","
		}
	}

	res += "}"
	return res
}

func EmptySet() IntSet {
	return make(map[int]bool)
}

func (ctx IntSet) Add(val int) IntSet {
	ctx[val] = true
	return ctx
}

func (ctx IntSet) Remove(val int) IntSet {
	ctx[val] = false
	return ctx
}

func (ctx IntSet) Size() int {
	var res = 0
	for _, v := range ctx {
		if (v) {
			res++
		}
	}
	return res
}
// FMap
// Filter
// Bind
// ToList