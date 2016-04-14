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

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c * Celsius(fcFactor) + Celsius(cfDiff))
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f-Fahrenheit(cfDiff)) * Fahrenheit(cfFactor))
}

func CToK(c Celsius) Kelvin {
	return Kelvin(c + Celsius(ckDiff))
}

func KToC(k Kelvin) Celsius {
	return Celsius(k - Kelvin(ckDiff))
}

func FToK(f Fahrenheit) Kelvin {
	return CToK(FToC(f))
}

func KToF(k Kelvin) Fahrenheit {
	return CToF(KToC(k))
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
	fmt.Println(CToF(BoilingC))
	fmt.Println(CToK(FreezingC))

	fmt.Println(FToC(CToF(FreezingC)))
	fmt.Println(KToC(CToK(BoilingC)))

	boilingK := CToK(BoilingC)
	fmt.Println(boilingK)
	fmt.Println(FToK(KToF(boilingK)))
}