package main

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
	FreezingC 	Celsius = 0
	BoilingC	Celsius = 100

	ckDiff 		float64 = 273.15
	cfDiff		float64 = 32
	cfFactor	float64 = 5/9
	fcFactor 	float64 = 9/5
)

// TODO make constants for conversions

func c2f(c Celsius) Fahrenheit {
	return Fahrenheit(c * Celsius(fcFactor) + Celsius(cfDiff))
}

func f2c(f Fahrenheit) Celsius {
	return Celsius((f-Fahrenheit(cfDiff)) * Fahrenheit(cfFactor))
}

func c2k(c Celsius) Kelvin {
	return Kelvin(c + Celsius(ckDiff))
}

func k2c(k Kelvin) Celsius {
	return Celsius(k - Kelvin(ckDiff))
}

func f2k(f Fahrenheit) Kelvin {
	return c2k(f2c(f))
}

func k2f(k Kelvin) Fahrenheit {
	return c2f(k2c(k))
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

	fmt.Println(f2c(c2f(FreezingC)))
	fmt.Println(k2c(c2k(BoilingC)))

	boilingK := c2k(BoilingC)
	fmt.Println(boilingK)
	fmt.Println(f2k(k2f(boilingK)))
}