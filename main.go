package main

import (
	"os"

	"github.com/ivan-marquez/go-playground/ch01"
)

func main() {
	ch01.FetchAll(os.Args[1:])
}
