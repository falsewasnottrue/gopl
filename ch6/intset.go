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

func (ctx IntSet) FMap(f func(val int) int) IntSet {
	var res = EmptySet()
	for k, v := range ctx {
		if (v) {
			res.Add(f(k))
		}
	}
	return res
}

func (ctx IntSet) Filter(p func(val int) bool) IntSet {
	var res = EmptySet()
	for k, v := range ctx {
		if (v && p(k)) {
			res.Add(k)
		}
	}
	return res
}

func (ctx IntSet) Bind(f func(val int) IntSet) IntSet {
	var res = EmptySet()
	for k, v := range ctx {
		if (v) {
			for k2, v2 := range f(k) {
				if (v2) {
					res.Add(k2)
				}
			}
		}
	}
	return res
}
// ToList