package ch04

import "fmt"

/*
 * Slices represent variable-length sequences whose elements all have
 * the same type. A slice type is written as []T, where the elements
 * have type T; it looks like an array type without a size.
 *
 * Arrays and slices are intimately connected. A slice is a lightweight
 * data structure that gives access to a subsequence (or perhaps all) of
 * the elements of an array, which is known as the slice’s underlying array.
 *
 * A slice has three components: a pointer, a length, and a capacity.
 *
 * The pointer points to the first element of the array that is reachable
 * through the slice, which is not necessarily the array’s first element.
 *
 * The length is the number of slice elements.
 *
 * It can’t exceed the capacity, which is usually the number of elements
 * between the start of the slice and the end of the underlying array.
 *
 * The built-in functions len and cap return those values.
 */

// PrintMonthSlices function define and prints overlapping
// slices for the second quarter and the northern summer
func PrintMonthSlices() {
	months := [...]string{
		1:  "January",
		2:  "February",
		3:  "March",
		4:  "April",
		5:  "May",
		6:  "June",
		7:  "July",
		8:  "August",
		9:  "September",
		10: "October",
		11: "November",
		12: "December",
	}

	q2 := months[4:7]
	summer := months[6:9]

	fmt.Println(q2)
	fmt.Println(summer)

	// slicing beyond cap(s) causes a panic, but slicing beyond len(s)
	// extends the slice, so the result may be longer than the original:

	// fmt.Println(summer[:20]) //panic: out of range
	endlessSummer := summer[:5] // extend a slice (within capacity)
	fmt.Println(endlessSummer)
}

/*
 * Since a slice contains a pointer to an element of an array, passing
 * a slice to a function permits the function to modify the underlying
 * array elements. In other words, copying a slice creates an alias for
 * the underlying array
 */

// Reverse function reverses a slice of ints in place.
func Reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

/*
 * a := [...]int{0, 1, 2, 3, 4, 5}
 * Reverse(a[:])
 * fmt.Println(a) // [5 4 3 2 1 0]
 */

// NonEmpty function returns a slice holding only the non-empty
// strings. The underlying array is modified during the call
func NonEmpty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}

	return strings[:i]
}
