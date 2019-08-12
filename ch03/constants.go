package ch03

import "fmt"

// When a sequence of constants is declared as a group,
// the right-hand side expression may be omitted for all
// but the first of the group, implying that the previous
// expression and its type should be used again.
const (
	a = 1
	b
	c = 2
	d
)

// ConstantsSequence prints values of constants
// declared as a group.
func ConstantsSequence() {
	fmt.Println(a, b, c, d) // "1 1 2 2"
}
