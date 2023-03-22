package main

import (
	"flag"
	"fmt"

	"gopl.io/ch7/tempconv"
)

type Kelvin float64

func CToK(c tempconv.Celsius) Kelvin    { return Kelvin(c + 273.15) }
func FToK(f tempconv.Fahrenheit) Kelvin { return Kelvin((f-32.0)*5.0/9.0 + 273.15) }

func (k Kelvin) String() string { return fmt.Sprintf("%g°K", k) }

// *celsiusFlag satisfies the flag.Value interface.
type celsiusFlag struct{ tempconv.Celsius }
type kelvinFlag struct{ Kelvin }

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit) // no error check needed
	switch unit {
	case "C", "°C":
		f.Celsius = tempconv.Celsius(value)
		return nil
	case "F", "°F":
		f.Celsius = tempconv.FToC(tempconv.Fahrenheit(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

func (f *kelvinFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit) // no error check needed
	switch unit {
	case "K", "°K":
		f.Kelvin = Kelvin(value)
	case "C", "°C":
		f.Kelvin = CToK(tempconv.Celsius(value))
		return nil
	case "F", "°F":
		f.Kelvin = FToK(tempconv.Fahrenheit(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

// CelsiusFlag defines a Celsius flag with the specified name,
// default value, and usage, and returns the address of the flag variable.
// The flag argument must have a quantity and a unit, e.g., "100C".
func CelsiusFlag(name string, value tempconv.Celsius, usage string) *tempconv.Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}

// KelvinFlag defines a Kelvin flag with the specified name,
// default value, and usage, and returns the address of the flag variable.
// The flag argument must have a quantity and a unit, e.g., "100K".
func KelvinFlag(name string, value Kelvin, usage string) *Kelvin {
	f := kelvinFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Kelvin
}

var tempC = tempconv.CelsiusFlag("tempC", 20.0, "the temperature in celsius")
var tempK = KelvinFlag("tempK", 20.0, "the temperature in kelvin's")

func main() {
	flag.Parse()
	fmt.Println(*tempC)
	fmt.Println(*tempK)
}
