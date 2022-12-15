// Package distanceconv performs Meters and Feet conversions.
package distanceconv

import "fmt"

type Meters float64
type Feets float64

func (m Meters) String() string { return fmt.Sprintf("%g m", m) }
func (f Feets) String() string  { return fmt.Sprintf("%g ft", f) }
