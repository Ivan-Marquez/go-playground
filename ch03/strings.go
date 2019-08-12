package ch03

import (
	"bytes"
	"fmt"
	"strings"
)

// basename function removes any prefix of s that looks like
// a file system path with components separated as slashes
// and it removes any suffix that looks like a file type.
func basename(s string) string {
	slash := strings.LastIndex(s, "/") // -1 if "/" not found
	s = s[slash+1:]

	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}

	return s
}

// fmt.Println(basename("a/b/c.go")) // "c"
// fmt.Println(basename("c.d.go")) // "c.d"
// fmt.Println(basename("abc")) // "abc"

// comma function inserts commas in a non-negative decimal
// integer string.
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

// intsToString function is like fmt.Sprintf(values) but adds
// commas.
func intsToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}

// fmt.Println(intsToString([]int{1, 2, 3})) // "[1, 2, 3]"
