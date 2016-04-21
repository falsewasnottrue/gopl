package main

import "fmt"

type Weekday int

const (
	Monday Weekday = iota
	Tuesday
	Wednesday
	Thursday
	Friday
)

const (
	_ = 1 << (10 * iota)
	KiB
	MiB
	GiB
)

func main() {
	
	fmt.Printf("%v\n%T\n", Thursday, Thursday)

	fmt.Println(KiB)
	fmt.Println(MiB)
	fmt.Println(GiB)
}