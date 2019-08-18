package ch04

import "fmt"

func printDefaultValues() {
	var q [3]int = [3]int{1, 2, 3}
	var r [3]int = [3]int{1, 2}

	fmt.Println(q)
	// Print the indices and elements.
	for i, v := range q {
		fmt.Printf("% d %d\n", i, v)
	} // Print the elements only.
	for _, v := range r {
		fmt.Printf("% d\n", v)
	}
}

type Currency int

// iota identifier is used in const declarations to simplify definitions of incrementing numbers.
// The value of iota is reset to 0 whenever the reserved word const appears in the source
// (i.e. each const block) and incremented by one afterwards.
const (
	USD Currency = iota
	EUR
	GBP
	RMB
)

func arrayWithEllipsis() {
	// In an array literal, if an ellipsis “...” appears in place of the length,
	// the array length is determined by the number of initializers.
	// It is also possible to specify a list of index and value pairs:
	symbol := [...]string{USD: "$", EUR: "€", GBP: "£", RMB: "¥"}

	fmt.Println(RMB, symbol[RMB]) // "3 ¥"
}

/*
 * When a function is called, a copy of each argument value is assigned to the
 * corresponding parameter variable, so the function receives a copy, not the
 * original. Passing large arrays in this way can be inefficient, and any changes
 * that the function makes to array elements affect only the copy, not the original.
 * In this regard, Go treats arrays like any other type, but this behavior is
 * different from languages that implicitly pass arrays by reference.
 *
 * We can explicitly pass a pointer to an array so that any modifications the
 * function makes to array elements will be visible to the caller. This
 * function zeroes the contents of a [32] byte array:
 */
func zeroes(ptr *[32]byte) {
	for i := range ptr {
		ptr[i] = 0
	}
}

/*
 * The array literal [32] byte{} yields an array of 32 bytes. Each element of the array
 * has the zero value for byte, which is zero. We can use that fact to write a different
 * version of zero:
 *
 * func zero( ptr *[ 32] byte) {
 *   *ptr = [32] byte{}
 * }
 */
