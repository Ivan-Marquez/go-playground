// Package tempconv performs Celsius and Fahrenheit conversions.
package tempconv

import "fmt"

// Celsius temperature type
type Celsius float64

// Fahrenheit temperature type
type Fahrenheit float64

const (
	// AbsoluteZeroC represents 0°C
	AbsoluteZeroC Celsius = -273.15
	// FreezingC represents 0°C
	FreezingC Celsius = 0
	// BoilingC represents 0°C
	BoilingC Celsius = 100
)

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g°F", f)
}
