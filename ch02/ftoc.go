package ch02

import "fmt"

// PrintFToC prints two Fahrenheit-to-Celsius conversions.
func PrintFToC() {
	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("%g°F or %g°C\n", freezingF, fToC(freezingF))
	fmt.Printf("%g°F or %g°C\n", boilingF, fToC(boilingF))
}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}
