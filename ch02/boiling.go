package ch02

import "fmt"

const boilingF = 212.0

// Boiling prints the boiling point of water
func Boiling() {
	var f = boilingF
	var c = (f - 32) * 5 / 9

	fmt.Printf("boiling point = %g°F or %g°C\n", f, c)
}
