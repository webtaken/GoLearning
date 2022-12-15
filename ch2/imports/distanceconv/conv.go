package distanceconv

// MToF converts Meters to Feets.
func MToF(m Meters) Feets { return Feets(m * 3.281) }

// FToM converts Feets to Meters.
func FToM(f Feets) Meters { return Meters(f / 3.281) }
