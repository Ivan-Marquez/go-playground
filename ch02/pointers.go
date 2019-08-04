package ch02

import "fmt"

// ComparePointers does various comparisons with pointers
func ComparePointers() {
	var x, y int

	// Pointers are comparable; two pointers are equal if and only
	// if they point to the same variable or both are nil.
	fmt.Println("Comparing pointers:", &x == &x, &x == &y, &x == nil) // true false false
}

// CompareVariableAddresses checks if two function calls
// return the same memory address
func CompareVariableAddresses() {
	/* It is perfectly safe for a function to return the address of
	 * a local variable. For instance, in the code below, the local
	 * variable v created by this particular call to f will remain in
	 * existence even after the call has returned, and the pointer p
	 * will still refer to it:
	 */

	fmt.Println(f() == f()) // "false"
}

func f() *int {
	v := 1

	return &v
}

func incr(p *int) int {
	*p++ // increments what p points to; does not change p

	return *p
}

// IncrementByReference function increments the variable that its
// argument points to and returns the new value of the variable
// so it may be used in an expression (call by reference).
func IncrementByReference() {
	v := 1
	incr(&v) // side effect; v is now 2

	fmt.Println(incr(&v)) // "3" (now v is 3)
}

// NewInt function shows another way to create a variable
// using the built-in function new().
func NewInt() *int {
	return new(int)
}

/*
 * A variable created with new() is no different from an ordinary
 * local variable whose address is taken, except that there’s no
 * need to invent (and declare) a dummy name, and we can use new(T)
 * in an expression. Thus new() is only a syntactic convenience, not
 * a fundamental notion.
 * The function above is the same as the following one:
 * func NewInt() *int {
 *   var dummy int
 *   return &dummy
 * }
 */
