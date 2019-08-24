package main

import (
	"fmt"

	"github.com/ivan-marquez/go-playground/ch04"
)

func main() {
	ch04.PrintMonthSlices()

	a := [...]int{0, 1, 2, 3, 4, 5}
	ch04.Reverse(a[:])
	fmt.Println(a) // [5 4 3 2 1 0]

	data := []string{"one", "", "three"}
	fmt.Printf("%q\n", ch04.NonEmpty(data))
	fmt.Printf("%q\n", data)

	ch04.Maps()
}
