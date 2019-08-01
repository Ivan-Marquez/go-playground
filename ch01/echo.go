// Package ch01 contains functions to print command-line arguments
package ch01

import (
	"fmt"
	"os"
	"strings"
)

// Echo1 prints its command-line arguments.
func Echo1() {
	var s, sep string

	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}

	fmt.Println(s)
}

// Echo2 prints its command-line arguments.
func Echo2() {
	s, sep := "", ""

	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}

	fmt.Println(s)
}

// Echo3 prints its command-line arguments.
func Echo3() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}
