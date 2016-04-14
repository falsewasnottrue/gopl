package main

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
	FreezingC 	Celsius = 0
	BoilingC	Celsius = 100
)

func c2f(c Celsius) Fahrenheit {
	return Fahrenheit(c * 9/5 + 32)
}

func c2k(c Celsius) Kelvin {
	return Kelvin(c + 273.15)
}

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g°F", f)
}

func (k Kelvin) String() string {
	return fmt.Sprintf("%g°K", k)
}

func main() {
	fmt.Println(FreezingC)
	fmt.Println(c2f(BoilingC))
	fmt.Println(c2k(FreezingC))
}