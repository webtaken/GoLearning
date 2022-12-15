// Package tempconv performs Celsius and Fahrenheit temperature computations.
package main

import "fmt"

type Celsius float64
type Fahrenheit float64
type MathIntPointer *int
type EngineeringIntPointer *int

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func main() {
	l1 := 5
	var p1 MathIntPointer = &l1
	var p2 EngineeringIntPointer = &l1
	fmt.Println(p1, *p1)
	fmt.Println(p2, *p2)
	fmt.Println(p1 == MathIntPointer(p2))
	c1 := Celsius(10)
	fmt.Printf("Celcius val: %v\n", c1)
}

func CToF(c Celsius) Fahrenheit  { return Fahrenheit(c*9/5 + 32) }
func FToC(f Fahrenheit) Celsius  { return Celsius((f - 32) * 5 / 9) }
func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }
