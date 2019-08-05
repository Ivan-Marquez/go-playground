package ch02

import "fmt"

// Celsius temperature representation
type Celsius float64

// Farenheit temperature representation
type Farenheit float64

const (
	// AbsoluteZeroC represents 0°C
	AbsoluteZeroC Celsius = -273.15
	// FreezingC represents 0°C
	FreezingC Celsius = 0
	// BoilingC represents 0°C
	BoilingC Celsius = 100
)

// CToF function converts Celsius to Farenheit
func CToF(c Celsius) Farenheit {
	return Farenheit(c*9/5 + 32)
}

// FToC function converts Farenheit to Celsius
func FToC(f Farenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

/*
 * This package defines two types, Celsius and Fahrenheit, for the two
 * units of temperature. Even though both have the same underlying type,
 * float64, they are not the same type, so they cannot be compared or
 * combined in arithmetic expressions. Distinguishing the types makes it
 * possible to avoid errors like inadvertently combining temperatures in
 * the two different scales; an explicit type conversion like Celsius(t)
 * or Fahrenheit(t) is required to convert from a float64.
 *
 * Celsius(t) and Fahrenheit(t) are conversions, not function calls. They
 * don’t change the value or representation in any way, but they make the
 * change of meaning explicit. On the other hand, the functions CToF and
 * FToC convert between the two scales; they do return different values.
 *
 * For every type T, there is a corresponding conversion operation T(x)
 * that converts the value x to type T. A conversion from one type to
 * another is allowed if both have the same underlying type, or if both
 * are unnamed pointer types that point to variables of the same
 * underlying type; these conversions change the type but not the
 * representation of the value. If x is assignable to T, a conversion is
 * permitted but is usually redundant.
 */

/*
 * Named types also make it possible to define new behaviors for values of
 * the type. These behaviors are expressed as a set of functions associated
 * with the type, called the type’s methods.
 *
 * The declaration below, in which the Celsius parameter c appears before the
 * function name, associates with the Celsius type a method named String that
 * returns c’s numeric value followed by °C.
 *
 * Many types declare a String method of this form because it controls how
 * values of the type appear when printed as a string by the fmt package.
 */

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}
