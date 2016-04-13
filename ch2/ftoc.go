package main

import "fmt"

func main() {
	freezing, boiling := 32.0, 212.0

	fmt.Printf("%g°F = %g°C\n", freezing, ftoc(freezing))
	fmt.Printf("%g°F = %g°C\n", boiling, ftoc(boiling))
}

func ftoc(f float64) float64 {
	return (f-32.0) * 5/9
}